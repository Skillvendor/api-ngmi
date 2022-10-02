package user

import (
	"api-ngmi/constants"
	"api-ngmi/cryptoEth"
	"api-ngmi/lib/utils"
	"api-ngmi/redis"
)

func TierFor(address string) int {
	isS1, _ := cryptoEth.HasNTS1(address)
	if isS1 {
		return constants.Gold
	}

	tier, _ := cryptoEth.TreasuryTierFor(address)
	if tier > 0 {
		return tier
	}

	isS2, _ := cryptoEth.HasNTS2(address)
	if isS2 {
		return utils.Max(constants.Silver, tier)
	}

	if tier > 0 {
		return tier
	}

	return 0
}

func AccessLevelFor(address string) int {
	tier, err := redis.AccessLevelFor(address)
	if err != nil {
		return tier
	}

	tier = TierFor(address)
	redis.CacheAccessLevelFor(address, tier)

	return tier
}
