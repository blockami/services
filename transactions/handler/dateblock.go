package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/blockami/common/key"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/errors"
)

type CVData struct {
	Data CVItems `json:"data"`
}

type CVItems struct {
	Items []Height `json:"items"`
}

type Height struct {
	Date  string `json:"signed_at"`
	Block int    `json:"height"`
}

func get_block_covalent(network string, timestamp int) (Height, error) {
	cvNetwork, ok := key.NetworkMapCovalent[network]
	if !ok {
		return Height{}, errors.InternalServerError("covalent", "network not supported")
	}

	cv_key, err := config.Get("covalenthq.key")
	if err != nil {
		return Height{}, errors.InternalServerError("covalenthq.config", err.Error())
	}

	date := time.Unix(int64(timestamp), 0)
	next_day_date := time.Unix(int64(timestamp+(24*60*60)), 0)

	timestamp_from := date.Format("2006-01-02")
	timestamp_to := next_day_date.Format("2006-01-02")
	url := fmt.Sprintf("https://api.covalenthq.com/v1/%s/block_v2/%s/%s/?key=%s&page-size=1", cvNetwork, timestamp_from, timestamp_to, cv_key.String(""))

	res, err := http.Get(url)
	if err != nil {
		return Height{}, errors.InternalServerError("covalenthq.query", err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Height{}, errors.InternalServerError("covalenthq.read_body", err.Error())
	}

	if res.StatusCode != 200 {
		return Height{}, errors.InternalServerError("covalenthq.error_code", string(body))
	}

	var data CVData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return Height{}, errors.InternalServerError("covalenthq.marshal", string(body))
	}

	if len(data.Data.Items) > 0 {
		return data.Data.Items[0], nil
	}
	return Height{}, nil
}
