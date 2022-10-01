package cryptoEth

import (
	contracts "api-ngmi/contracts"
	"log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TierFor(userAddress string) (int, error) {
	client, err := ethclient.Dial(NgmiTreasury.NodeProvider)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(NgmiTreasury.HexAddress)

	instance, err := contracts.NewNgmiTreasury(address, client)
	if err != nil {
		log.Fatal("can't establish connection S1", err)
		return 0, err
	}

	hexUserAddress := common.HexToAddress(userAddress)

	tier, err := instance.GetUserTier(&bind.CallOpts{}, hexUserAddress)
	if err != nil {
		log.Fatal("balance can't be retrieved", err)
		return 0, err
	}

	return int(tier.Int64()), nil
}
