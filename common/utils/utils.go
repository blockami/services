package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/micro/micro/v3/service/config"
)

func get_ethclient(network string) (*ethclient.Client, error) {
	web3URL, err := config.Get("web3." + network)
	if err != nil {
		return nil, err
	}

	ethclient, err := ethclient.Dial(web3URL.String(""))
	if err != nil {
		return nil, err
	}

	return ethclient, nil
}
