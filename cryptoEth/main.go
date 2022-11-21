package cryptoEth

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var NTS1 EthContract
var NTS2 EthContract
var BSC Chain
var NgmiTreasury EthContract

var NTS1Addr = "0xb668beB1Fa440F6cF2Da0399f8C28caB993Bdd65"
var NTS2Addr = "0x9b091d2E0Bb88acE4fe8f0faB87b93D8bA932EC4"
var NgmiTreasuryAddrTestnet = "0x61Ceb923071C1379dabFbbb308906daD2Ac7471e"
var NgmiTreasuryAddrProd = "0xeecefec8baA7fFd7A66043B4D80B7e1Ea2f949f9"

var LucianAddress = "0x5f991BdCCCff8a5C6C00A5A708dc4f649Eb5887C"
var LucianAddress2 = "0x090E2883B09Efa50F553DF70b6bf716F1AE74e7B"

// var NTS1

type EthContract struct {
	NodeProvider string
	HexAddress   string
}

type Chain struct {
	NodeProvider string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
	}

	QuickNode := os.Getenv("QUICKNODE_NODE")
	InfuraNode := os.Getenv("INFURA_NODE")
	environment := os.Getenv("ENV")

	NTS1 = EthContract{
		NodeProvider: InfuraNode,
		HexAddress:   NTS1Addr,
	}

	NTS2 = EthContract{
		NodeProvider: InfuraNode,
		HexAddress:   NTS2Addr,
	}

	var treasuryAddress string

	if environment == "prod" {
		treasuryAddress = NgmiTreasuryAddrProd
	} else {
		treasuryAddress = NgmiTreasuryAddrTestnet
	}

	NgmiTreasury = EthContract{
		NodeProvider: QuickNode,
		HexAddress:   treasuryAddress,
	}

	BSC = Chain{
		NodeProvider: QuickNode,
	}
}
