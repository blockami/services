package handler

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"

	common "github.com/ethereum/go-ethereum/common"

	ethabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"

	interpret "github.com/blockami/servicesinterpret/proto"

	abi "github.com/blockami/servicesabi/proto"
)

type Interpret struct {
	abi_client abi.AbiService
}

// Return a new handler
func New(service *service.Service) *Interpret {
	abi_client := abi.NewAbiService("abi", service.Client())

	return &Interpret{
		abi_client: abi_client,
	}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Interpret) InterpretFunction(ctx context.Context, req *interpret.InterpretFunctionRequest, rsp *interpret.InterpretFunctionResponse) error {

	abi_response, err := e.abi_client.ContractAbi(context.Background(), &abi.AbiRequest{
		Network: req.Network,
		Address: req.Contract,
	})
	if err != nil {
		return nil
	}

	contract_abi, err := ethabi.JSON(strings.NewReader(abi_response.Abi))
	if err != nil {
		return errors.InternalServerError("abi.parse", err.Error())
	}

	//transfer to contract
	if len(req.Input) < 8 {
		return nil
	}

	decodedSig, err := hex.DecodeString(req.Input[0:8])
	if err != nil {
		return errors.InternalServerError("decode.abi_method", err.Error())
	}

	// recover Method from signature and ABI
	method, err := contract_abi.MethodById(decodedSig)
	if err != nil {
		log.Debug(err)
		return nil
	}

	// decode txInput Payload
	decodedData, err := hex.DecodeString(req.Input[8:])
	if err != nil {
		log.Debug(err)
		return nil
	}

	receivedMap := map[string]interface{}{}
	err = method.Inputs.UnpackIntoMap(receivedMap, decodedData)
	if err != nil {
		log.Debug(err)
		return nil
	}

	rsp.Method = method.Name
	rsp.MethodFull = method.Sig
	for k, v := range receivedMap {
		rsp.Parameters = append(rsp.Parameters,
			&interpret.Parameter{
				Name:  k,
				Value: fmt.Sprintf("%v", v),
			},
		)
	}

	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Interpret) InterpretLog(ctx context.Context, req *interpret.InterpretLogRequest, rsp *interpret.InterpretLogResponse) error {
	abi_response, err := e.abi_client.ContractAbi(context.Background(), &abi.AbiRequest{
		Network: req.Network,
		Address: req.Contract,
	})
	if err != nil {
		return nil
	}

	contract_abi, err := ethabi.JSON(strings.NewReader(abi_response.Abi))
	if err != nil {
		return nil
	}

	event, err := contract_abi.EventByID(common.HexToHash(req.Topics[0]))
	if err != nil {
		return nil
	}

	var indexedArguments []ethabi.Argument
	for _, v := range event.Inputs {
		if v.Indexed {
			indexedArguments = append(indexedArguments, v)
		}
	}

	var hashTopics []common.Hash
	for _, v := range req.Topics[1:] {
		hashTopics = append(hashTopics, common.HexToHash(v))
	}

	indexedParams := make(map[string]interface{})
	err = ethabi.ParseTopicsIntoMap(indexedParams, indexedArguments, hashTopics)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	dataHash, err := hex.DecodeString(req.Data)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	dataParams := map[string]interface{}{}
	err = contract_abi.UnpackIntoMap(dataParams, event.Name, dataHash)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	rsp.Name = event.Name
	// rsp.MethodFull = method.Sig
	for k, v := range indexedParams {
		rsp.Parameters = append(rsp.Parameters,
			&interpret.Parameter{
				Name:  k,
				Value: fmt.Sprintf("%v", v),
			},
		)
	}

	for k, v := range dataParams {
		rsp.Parameters = append(rsp.Parameters,
			&interpret.Parameter{
				Name:  k,
				Value: fmt.Sprintf("%v", v),
			},
		)
	}

	return nil
}
