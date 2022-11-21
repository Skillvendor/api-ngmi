package user

import (
	"api-ngmi/constants"
	"api-ngmi/cryptoEth"
	"api-ngmi/models"
	"api-ngmi/redis"
	"fmt"
)

func TierFor(address string) int {
	ntPayment := models.NtPayment{AccessWallet: address}

	found := ntPayment.Find()

	if found {
		isS1, _ := cryptoEth.HasNTS1(address)
		if isS1 {
			fmt.Println("Has s1")
			return constants.Gold
		}

		isS2, _ := cryptoEth.HasNTS2(address)
		if isS2 {
			// return utils.Max(constants.Silver, tier)
			fmt.Println("Has s2")
			return constants.Gold
		}

		fmt.Println("it is not S1/S2")
	}

	fmt.Println("Getting treasury tier")
	tier, _ := cryptoEth.TreasuryTierFor(address)
	fmt.Println("this is the treasury tier", tier)

	if tier > 0 {
		return tier
	}

	return 0
}

func AccessLevelFor(address string) int {
	tier, err := redis.AccessLevelFor(address)

	fmt.Println("Redis error?", err)
	if err == nil {
		return tier
	}

	tier = TierFor(address)
	redis.CacheAccessLevelFor(address, tier)

	fmt.Println("I am returning this tier", tier)
	return tier
}
