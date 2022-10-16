package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/micro/micro/v3/service"
	"github.com/open-policy-agent/opa/sdk"

	information "github.com/blockami/services/information/proto"
	interpret "github.com/blockami/services/interpret/proto"
	pricing "github.com/blockami/services/pricing/proto"
	transactions "github.com/blockami/services/transactions/proto"
)

type Information struct {
	transactions_client transactions.TransactionsService
	interpret_client    interpret.InterpretService
	pricing_client      pricing.PricingService
	opa_sdk             *sdk.OPA
}

// Return a new handler
func New(service *service.Service) *Information {
	transactions_client := transactions.NewTransactionsService("transactions", service.Client())
	interpret_client := interpret.NewInterpretService("interpret", service.Client())
	pricing_client := pricing.NewPricingService("pricing", service.Client())

	config := []byte(fmt.Sprintf(`{
		"services": {
			"blockami-repo": {
				"url": "https://blockami.github.io/opa-bundle/"
			}
		},
		"bundles": {
			"blockami-repo": {
				"resource": "blockami-bundle.tar.gz",
				"persist": true
			}
		}
	}`))

	opa_sdk, err := sdk.New(context.Background(), sdk.Options{
		Config: bytes.NewReader(config),
	})
	if err != nil {
		fmt.Println(err)
	}

	return &Information{
		interpret_client:    interpret_client,
		transactions_client: transactions_client,
		pricing_client:      pricing_client,
		opa_sdk:             opa_sdk,
	}
}

func EventsValue(events *[]*information.Event, network string, currency string, timestamp int, pricing_client pricing.PricingService) (float32, error) {
	highestPrio := 0
	var highestPrioEvent *information.Event

	for _, event := range *events {
		if event.Priority > int32(highestPrio) && event.Value != nil {
			highestPrioEvent = event
		}
	}

	if highestPrioEvent != nil {

		pricing_info, err := pricing_client.Call(context.Background(), &pricing.PricingRequest{
			Currency:  currency,
			Network:   network,
			Contract:  highestPrioEvent.Value.Token,
			Amount:    highestPrioEvent.Value.Amount,
			Timestamp: int32(timestamp),
		})
		if err != nil {
			return 0.0, err
		}

		return pricing_info.TotalValue, nil
	}

	return 0.0, nil
}

func (e *Information) TransactionInformation(ctx context.Context, req *information.TransactionInformationRequest, rsp *information.TransactionInformationResponse) error {

	if req.Network == "" {
		req.Network = "ETH"
	}

	if req.Currency == "" {
		req.Currency = "USD"
	}

	tx_response, err := e.transactions_client.SingleTransaction(context.Background(), &transactions.SingleTransactionRequest{
		Network: req.Network,
		TxHash:  req.TxHash,
	})
	if err != nil {
		return err
	}

	interpret_response, err := e.interpret_client.InterpretFunction(context.Background(), &interpret.InterpretFunctionRequest{
		Network:  req.Network,
		Contract: tx_response.To,
		Input:    tx_response.Input,
	})
	if err != nil {
		return err
	}

	canonicalAddress := common.HexToAddress(tx_response.From)
	if req.Address != "" {
		canonicalAddress = common.HexToAddress(req.Address)
	}

	rsp.Network = req.Network
	rsp.Address = canonicalAddress.String()
	rsp.Transaction = &information.Transaction{
		RawData: &information.TransactionData{
			TxHash:    tx_response.TxHash,
			Block:     tx_response.Block,
			From:      tx_response.From,
			To:        tx_response.To,
			Value:     tx_response.Value,
			Input:     tx_response.Input,
			Status:    tx_response.Status,
			Timestamp: tx_response.Timestamp,
		},
		Information: &information.TransactionInformation{
			Method:     interpret_response.Method,
			MethodFull: interpret_response.MethodFull,
		},
	}

	for _, v := range interpret_response.Parameters {
		rsp.Transaction.Information.Parameters = append(rsp.Transaction.Information.Parameters, &information.Parameter{
			Name:  v.Name,
			Value: v.Value,
		})
	}

	logs_response, err := e.transactions_client.TransactionLogs(context.Background(), &transactions.TransactionLogsRequest{
		Network: req.Network,
		TxHash:  req.TxHash,
	})
	if err != nil {
		fmt.Println("Error calling transactions: ", err)
		return err
	}

	for _, log := range logs_response.Logs {
		log_answer := &information.Log{
			RawData: &information.LogData{
				TxHash:   log.TxHash,
				Contract: log.Contract,
				Index:    log.Index,
				Data:     log.Data,
				Topics:   log.Topics,
			},
		}

		interpret_log_response, err := e.interpret_client.InterpretLog(context.Background(), &interpret.InterpretLogRequest{
			Network:  req.Network,
			Contract: log.Contract,
			Topics:   log.Topics,
			Data:     log.Data,
		})

		if err != nil {
			fmt.Println("Error calling interpret: ", err)
			rsp.Logs = append(rsp.Logs, log_answer)
			continue
		}

		log_answer.Information = &information.LogInformation{
			Name: interpret_log_response.Name,
		}

		for _, v := range interpret_log_response.Parameters {
			log_answer.Information.Parameters = append(log_answer.Information.Parameters, &information.Parameter{
				Name:  v.Name,
				Value: v.Value,
			})
		}

		rsp.Logs = append(rsp.Logs, log_answer)
	}

	info_map, err := StructToMap(rsp)
	if err != nil {
		fmt.Println("Error converting: ", err)
	}
	result, err := e.opa_sdk.Decision(ctx, sdk.DecisionOptions{Path: "/blockami", Input: info_map})
	if err != nil {
		fmt.Println("Error OPA: ", err)
	}

	for _, element := range result.Result.(map[string]interface{})["events"].([]interface{}) {
		test, err := json.Marshal(element)
		if err != nil {
			fmt.Println("Error marshal: ", err)
		}

		var event information.Event
		err = json.Unmarshal(test, &event)
		if err != nil {
			fmt.Println("Error marshal: ", err)
		}
		rsp.Events = append(rsp.Events, &event)
	}

	value, err := EventsValue(&rsp.Events, req.Network, req.Currency, int(rsp.Transaction.RawData.Timestamp), e.pricing_client)
	if err != nil {
		return err
	}

	rsp.Value = &information.Value{
		Amount:   value,
		Currency: req.Currency,
	}
	return nil
}

// Converts a struct to a map while maintaining the json alias as keys
func StructToMap(obj interface{}) (newMap map[string]interface{}, err error) {
	data, err := json.Marshal(obj) // Convert to a json string

	if err != nil {
		return
	}

	err = json.Unmarshal(data, &newMap) // Convert to a map
	return
}
