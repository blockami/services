package handler

import (
	"bytes"
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/store"
	"github.com/nanmu42/etherscan-api"

	abi "github.com/blockami/services/abi/proto"
)

type Abi struct{}

// Return a new handler
func New() *Abi {
	return &Abi{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Abi) ContractAbi(ctx context.Context, req *abi.AbiRequest, rsp *abi.AbiResponse) error {
	records, _ := store.Read(req.Network + req.Address)
	if len(records) == 1 {
		rsp.Abi = string(records[0].Value)
		return nil
	}

	web3URL, err := config.Get("web3." + req.Network)
	if err != nil {
		log.Error(err)
		return errors.InternalServerError("web3.config", err.Error())
	}

	ethclient, err := ethclient.Dial(web3URL.String(""))
	if err != nil {
		return errors.InternalServerError("web3.url", err.Error())
	}

	blocknumber, err := ethclient.BlockNumber(ctx)
	if err != nil {
		return errors.InternalServerError("blocknumber", err.Error())
	}
	if req.Block > 0 {
		blocknumber = uint64(req.Block)
	}

	address := req.Address
	proxyContract, err := get_proxy_contract(ctx, req.Address, blocknumber, ethclient)
	if err != nil {
		return errors.InternalServerError("proxy", err.Error())
	} else if proxyContract != "" {
		address = proxyContract
	}

	raw_abi, err := get_abi(req.Network, address)

	if err != nil {
		return err
	}

	rsp.Abi = raw_abi

	return nil
}

func get_abi(network string, address string) (string, error) {
	etherscanAPI, err := config.Get("etherscan." + network + ".key")
	if err != nil {
		log.Error(err)
		return "", errors.InternalServerError("web3.config", err.Error())
	}

	client := etherscan.New(etherscan.Mainnet, etherscanAPI.String(""))

	contract_abi, err := client.ContractABI(address)
	if err != nil {
		if strings.Contains(err.Error(), "Contract source code not verified") {
			return "", errors.InternalServerError("etherscan.contract_not_verified", "Contract source code not verified "+address)
		} else {
			log.Error(err)
			return "", errors.InternalServerError("etherscan.abi", err.Error())
		}
	}

	err = store.Write(&store.Record{
		Key:   network + address,
		Value: []byte(contract_abi),
	})
	if err != nil {
		log.Error(err)
	}

	return contract_abi, nil
}

func get_proxy_contract(ctx context.Context, address string, blockNumber uint64, ethclient *ethclient.Client) (string, error) {
	proxy_slots := [2]string{
		"0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc",
		"0x7050c9e0f4ca769c69bd3a8ef740bc37934f8e2c036e5a723fd8ee048ed3f8c3",
	}

	emptyByteVar := make([]byte, 32)

	for _, slot := range proxy_slots {
		proxy, err := ethclient.StorageAt(ctx, common.HexToAddress(address), common.HexToHash(slot), new(big.Int).SetUint64(blockNumber))
		if err != nil {
			return "", err
		}

		if !bytes.Equal(proxy, emptyByteVar) {
			return common.BytesToAddress(proxy).Hex(), nil
		}
	}

	return "", nil
}
