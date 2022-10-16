package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/blockami/common/key"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/store"

	token "github.com/blockami/token/proto"
)

type Token struct{}

// Return a new handler
func New() *Token {
	return &Token{}
}

type CVData struct {
	Data []CVToken `json:"data"`
}

type CVToken struct {
	Decimals int32  `json:"contract_decimals"`
	Name     string `json:"contract_name"`
	Ticker   string `json:"contract_ticker_symbol"`
	LogoUrl  string `json:"logo_url"`
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Token) GetTokenInfo(ctx context.Context, req *token.TokenInfoRequest, rsp *token.TokenInfoResponse) error {
	if val, ok := key.BaseTokensMap[req.Address]; ok {
		baseToken := CVToken{}
		err := json.Unmarshal([]byte(val), &baseToken)
		if err != nil {
			return errors.InternalServerError("store.jsonparse", err.Error())
		}

		rsp.Decimals = baseToken.Decimals
		rsp.LogoUrl = baseToken.LogoUrl
		rsp.Name = baseToken.Name
		rsp.Ticker = baseToken.Ticker

		return nil
	}

	records, _ := store.Read(req.Network + req.Address)
	if len(records) == 1 {
		var token CVToken
		err := json.Unmarshal(records[0].Value, &token)
		if err != nil {
			return errors.InternalServerError("store.jsonparse", err.Error())
		}

		rsp.Decimals = token.Decimals
		rsp.LogoUrl = token.LogoUrl
		rsp.Name = token.Name
		rsp.Ticker = token.Ticker

		return nil
	}

	cvNetwork, ok := key.NetworkMapCovalent[req.Network]
	if !ok {
		return errors.InternalServerError("covalent", "network not supported")
	}

	cv_key, err := config.Get("covalenthq.key")
	if err != nil {
		return errors.InternalServerError("covalenthq.config", err.Error())
	}

	url := fmt.Sprintf("https://api.covalenthq.com/v1/pricing/historical_by_addresses_v2/%s/%s/%s/?key=%s&prices-at-asc=false", cvNetwork, "USD", req.Address, cv_key.String(""))
	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		return errors.InternalServerError("covalenthq.query", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.InternalServerError("covalenthq.read_body", err.Error())
	}

	if res.StatusCode != 200 {
		return errors.InternalServerError("covalenthq.error_code", string(body))
	}

	var data CVData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return errors.InternalServerError("covalenthq.jsonparse", err.Error())
	}

	if len(data.Data) > 0 {
		rsp.Decimals = data.Data[0].Decimals
		rsp.LogoUrl = data.Data[0].LogoUrl
		rsp.Name = data.Data[0].Name
		rsp.Ticker = data.Data[0].Ticker

		marshalled, err := json.Marshal(data.Data[0])
		if err != nil {
			log.Error(err)
		}
		err = store.Write(&store.Record{
			Key:   req.Network + req.Address,
			Value: marshalled,
		})
		if err != nil {
			log.Error(err)
		}
	}

	return nil
}
