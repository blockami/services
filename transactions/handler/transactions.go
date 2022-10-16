package handler

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/micro/micro/v3/service/config"
	"github.com/micro/micro/v3/service/errors"
	log "github.com/micro/micro/v3/service/logger"
	"github.com/nanmu42/etherscan-api"

	transactions "github.com/blockami/transactions/proto"
)

// TODO initialize clients on demand depending on the network with overloaded struct and saving clients on an array
type Transactions struct {
}

func New() *Transactions {
	return &Transactions{}
}

func GetTransactionMessage(tx *types.Transaction) types.Message {
	msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), nil)
	if err != nil {
		log.Fatal(err)
	}
	return msg
}

func (e *Transactions) SingleTransaction(ctx context.Context, req *transactions.SingleTransactionRequest, rsp *transactions.SingleTransactionResponse) error {
	web3URL, err := config.Get("web3." + req.Network)
	if err != nil {
		log.Error(err)
		return errors.InternalServerError("web3.config", err.Error())
	}

	client, err := ethclient.Dial(web3URL.String(""))
	if err != nil {
		return errors.InternalServerError("web3.url", err.Error())
	}

	tx, _, err := client.TransactionByHash(context.Background(), common.HexToHash(req.TxHash))
	if err != nil {
		return errors.InternalServerError("web3.tx_by_hash", err.Error())
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(req.TxHash))
	if err != nil {
		return errors.InternalServerError("web3.receipt_by_hash", err.Error())
	}

	block, err := client.BlockByNumber(context.Background(), receipt.BlockNumber)
	if err != nil {
		return errors.InternalServerError("web3.block", err.Error())
	}

	rsp.TxHash = req.TxHash
	rsp.Block = receipt.BlockNumber.Uint64()
	rsp.From = GetTransactionMessage(tx).From().Hex()
	rsp.To = tx.To().Hex()
	rsp.Value = tx.Value().String()
	rsp.Input = hex.EncodeToString(tx.Data())
	rsp.Status = receipt.Status == 1
	rsp.Timestamp = block.Time()

	return nil
}

func (e *Transactions) AllTransactions(ctx context.Context, req *transactions.AllTransactionsRequest, rsp *transactions.AllTransactionsResponse) error {
	etherscanAPI, err := config.Get("etherscan." + req.Network + ".key")
	if err != nil {
		log.Error(err)
		return errors.InternalServerError("web3.config", err.Error())
	}

	client := etherscan.New(etherscan.Mainnet, etherscanAPI.String(""))

	var min, max = 0, 99999999

	if req.From != "" {
		from, err := strconv.Atoi(req.From)
		if err != nil {
			log.Error(err)
			return errors.InternalServerError("string convert", err.Error())
		}
		height, err := get_block_covalent(req.Network, from)
		if err != nil {
			log.Error("Cannot get block from timestamp using default")
		}
		min = height.Block
		fmt.Println(height.Block)
	}

	if req.To != "" {
		to, err := strconv.Atoi(req.To)
		if err != nil {
			log.Error(err)
			return errors.InternalServerError("string convert", err.Error())
		}
		height, err := get_block_covalent(req.Network, to)
		if err != nil {
			log.Error("Cannot get block from timestamp using default")
		}
		max = height.Block
	}

	normal_txs, err := client.NormalTxByAddress(req.Address, &min, &max, 0, 0, true)
	if err != nil {
		normal_txs = []etherscan.NormalTx{}
	}
	internal_txs, err := client.InternalTxByAddress(req.Address, &min, &max, 0, 0, true)
	if err != nil {
		internal_txs = []etherscan.InternalTx{}
	}

	//Used to avoid duplicates
	txArray := make(map[string]bool)

	for _, tx := range normal_txs {
		txArray[tx.Hash] = true
		rsp.Transactions = append(rsp.Transactions,
			&transactions.SingleTransactionResponse{
				TxHash:    tx.Hash,
				Block:     uint64(tx.BlockNumber),
				From:      tx.From,
				To:        tx.To,
				Value:     tx.Value.Int().String(),
				Input:     tx.Input,
				Status:    tx.TxReceiptStatus == "1",
				Timestamp: uint64(tx.TimeStamp.Time().Unix()),
			},
		)
	}

	for _, tx := range internal_txs {
		if !txArray[tx.Hash] {
			txArray[tx.Hash] = true
			rsp.Transactions = append(rsp.Transactions,
				&transactions.SingleTransactionResponse{
					TxHash: tx.Hash,
					Block:  uint64(tx.BlockNumber),
					From:   tx.From,
					To:     tx.To,
					Value:  tx.Value.Int().String(),
					Input:  tx.Input,
					Status: true,
				},
			)
		}
	}

	return nil
}

func (e *Transactions) TransactionLogs(ctx context.Context, req *transactions.TransactionLogsRequest, rsp *transactions.TransactionLogsResponse) error {

	web3URL, err := config.Get("web3." + req.Network)
	if err != nil {
		log.Error(err)
		return errors.InternalServerError("web3.config", err.Error())
	}

	client, err := ethclient.Dial(web3URL.String(""))
	if err != nil {
		return errors.InternalServerError("web3.url", err.Error())
	}

	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(req.TxHash))
	if err != nil {
		return errors.InternalServerError("web3.receipt_by_hash", err.Error())
	}

	for _, log := range receipt.Logs {

		new_log := &transactions.Log{
			TxHash:   req.TxHash,
			Contract: log.Address.String(),
			Index:    int32(log.Index),
			Data:     hex.EncodeToString(log.Data),
		}

		for _, topic := range log.Topics {
			new_log.Topics = append(new_log.Topics,
				topic.String(),
			)
		}

		rsp.Logs = append(rsp.Logs,
			new_log,
		)
	}

	return nil
}
