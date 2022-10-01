package cryptoEth

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var InfuraNode = "https://rinkeby.infura.io"
var NTS1Addr = "0xb668beB1Fa440F6cF2Da0399f8C28caB993Bdd65"

// var NTS1

type EthContract struct {
	NodeProvider string
	HexAddress   string
}

func (contract *EthContract) Connect() error {
	client, err := ethclient.Dial(InfuraNode)

	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(contract.HexAddress)
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	return err == nil
	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version)
}

func init() {

	NTS1, err := EthContract{
		NodeProvider: InfuraNode,
		HexAddress:   NTS1Addr,
	}
}
