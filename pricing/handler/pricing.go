package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/blockami/services/common/key"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/errors"

	pricing "github.com/blockami/services/pricing/proto"
)

type Pricing struct {
}

// Return a new handler
func New() *Pricing {
	return &Pricing{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Pricing) Call(ctx context.Context, req *pricing.PricingRequest, rsp *pricing.PricingResponse) error {

	cvNetwork, ok := key.NetworkMapCovalent[req.Network]
	if !ok {
		return errors.InternalServerError("covalent", "network not supported")
	}
	var cvContract string

	if len(req.Contract) != 42 {
		cvContract, ok = key.NetworkAddressMap[req.Contract]
		if !ok {
			return errors.InternalServerError("covalent", "contract not supported")
		}
	} else {
		cvContract = req.Contract
	}

	result, err := get_price_covalent(cvNetwork, cvContract, req.Currency, int(req.Timestamp))
	if err != nil {
		return err
	}

	amount, err := strconv.ParseFloat(req.Amount, 64)
	if err != nil {
		errors.InternalServerError("web3token.decimals", err.Error())
	}

	rsp.AmountDecimals = float32(amount) / float32(math.Pow10(int(result.ContactMetadata.Decimals)))
	rsp.UnitValue = result.Price
	rsp.TotalValue = rsp.AmountDecimals * rsp.UnitValue

	return nil
}

type CVData struct {
	Data []CVPrice `json:"data"`
}

type CVPrice struct {
	Prices []Price `json:"prices"`
}

type Price struct {
	ContactMetadata CVMetadata `json:"contract_metadata"`
	Date            string     `json:"date"`
	Price           float32    `json:"price"`
}

type CVMetadata struct {
	Decimals int    `json:"contract_decimals"`
	Name     string `json:"contract_name"`
	Symbol   string `json:"contract_ticker_symbol"`
}

func get_price_covalent(network string, contract string, currency string, timestamp int) (Price, error) {
	cv_key, err := config.Get("covalenthq.key")
	if err != nil {
		return Price{}, errors.InternalServerError("covalenthq.config", err.Error())
	}

	date := time.Unix(int64(timestamp), 0)

	timestamp_from := date.Format("2006-01-02")
	timestamp_to := date.Format("2006-01-02")
	url := fmt.Sprintf("https://api.covalenthq.com/v1/pricing/historical_by_addresses_v2/%s/%s/%s/?key=%s&from=%s&to=%s&prices-at-asc=false", network, currency, contract, cv_key.String(""), timestamp_from, timestamp_to)

	res, err := http.Get(url)
	if err != nil {
		return Price{}, errors.InternalServerError("covalenthq.query", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Price{}, errors.InternalServerError("covalenthq.read_body", err.Error())
	}

	if res.StatusCode != 200 {
		return Price{}, errors.InternalServerError("covalenthq.error_code", string(body))
	}

	var data CVData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println(string(body))
		return Price{}, errors.InternalServerError("covalenthq.marshal", string(body))
	}

	if len(data.Data[0].Prices) > 0 {
		return data.Data[0].Prices[0], nil
	}
	return Price{}, nil
}
