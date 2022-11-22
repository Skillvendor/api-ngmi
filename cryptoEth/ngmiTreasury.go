package cryptoEth

import (
	contracts "api-ngmi/contracts"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var silverTiers = []int{2, 4}
var goldTiers = []int{3, 5, 6}

func TreasuryTierToAppTier(tier int) int {
	fmt.Println("This is what I got from contract", tier)
	if utils.ContainsInt(goldTiers, tier) {
		return constants.Gold
	}

	if utils.ContainsInt(silverTiers, tier) {
		return constants.Silver
	}

	return constants.Free
}

func TreasuryTierFor(userAddress string) (int, error) {
	client, err := ethclient.Dial(NgmiTreasury.NodeProvider)
	if err != nil {
		fmt.Println(err)
	}

	address := common.HexToAddress(NgmiTreasury.HexAddress)

	instance, err := contracts.NewNgmiTreasury(address, client)
	if err != nil {
		fmt.Println("can't establish connection to NGMITreasury", err)
		return 0, err
	}

	hexUserAddress := common.HexToAddress(userAddress)

	subscribed, err := instance.IsUserSubscribed(&bind.CallOpts{}, hexUserAddress)

	if err != nil {
		fmt.Println("subscription can't be retrieved", err)
		return 0, err
	}

	if !subscribed {
		fmt.Println("it is not SUBSCRIBED")
		return 0, nil
	}

	tier, err := instance.GetUserTier(&bind.CallOpts{}, hexUserAddress)
	if err != nil {
		fmt.Println("balance can't be retrieved", err)
		return 0, err
	}

	return TreasuryTierToAppTier(int(tier.Int64())), nil
}
