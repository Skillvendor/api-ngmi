package cryptoEth

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	Pending int = 2
	Success int = 1
	Fail    int = 0
)

func TransactionStatusFor(transaction_id string) int {
	client, err := ethclient.Dial(BSC.NodeProvider)
	if err != nil {
		log.Fatal(err)
	}

	txHash := common.HexToHash(transaction_id)

	tx, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return Pending
	}

	if int(tx.Status) == Success {
		return Success
	}

	return Fail
}
