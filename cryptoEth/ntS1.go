package cryptoEth

import (
	contracts "api-ngmi/contracts"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (contract *EthContract) GetBalanceS1(userAddress common.Address) (*big.Int, error) {
	client, err := ethclient.Dial(contract.NodeProvider)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(contract.HexAddress)
	instance, err := contracts.NewS1citizen(address, client)
	if err != nil {
		log.Fatal("can't establish connection S1", err)
		return nil, err
	}

	result, err := instance.BalanceOf(&bind.CallOpts{}, userAddress)
	if err != nil {
		log.Fatal("can't get balance of", err)
		return nil, err
	}

	return result, nil
}

func HasNTS1(userAddress string) (bool, error) {
	hexUserAddress := common.HexToAddress(userAddress)

	balance, err := NTS1.GetBalanceS1(hexUserAddress)
	if err != nil {
		log.Fatal("balance can't be retrieved", err)
		return false, err
	}

	return balance.Cmp(big.NewInt(0)) > 0, nil
}
