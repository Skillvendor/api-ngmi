package user

import (
	"api-ngmi/constants"
	"api-ngmi/cryptoEth"
	"api-ngmi/models"
	"api-ngmi/redis"
	"fmt"
	"time"
)

func TierFor(address string) int {
	ntPayment := models.NtPayment{AccessWallet: address}

	found := ntPayment.Find()

	fmt.Println("This is what I found", found)
	if found {
		isS1, _ := cryptoEth.HasNTS1(ntPayment.CitizenWallet)

		fmt.Println("Check if S1", isS1)
		if isS1 {
			fmt.Println("Has s1")
			return constants.Gold
		}

		isS2, _ := cryptoEth.HasNTS2(ntPayment.CitizenWallet)
		if isS2 {
			// return utils.Max(constants.Silver, tier)
			fmt.Println("Has s2")
			return constants.Gold
		}
	}

	communityPayment := models.CommunityPayment{AccessWallet: address}

	foundCommunity := communityPayment.Find()

	if foundCommunity {
		today := time.Now()

		expired := today.After(communityPayment.ExpiresAt)

		if !expired {
			if communityPayment.AccessLevel == "silver" {
				return constants.Silver
			}

			return constants.Gold
		}
	}

	tier, _ := cryptoEth.TreasuryTierFor(address)

	if tier > 0 {
		return tier
	}

	return 0
}

func AccessLevelFor(address string) int {
	tier, err := redis.AccessLevelFor(address)

	if err == nil {
		return tier
	}

	tier = TierFor(address)
	redis.CacheAccessLevelFor(address, tier)

	return tier
}
