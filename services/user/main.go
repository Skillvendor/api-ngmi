package user

import (
	"api-ngmi/cryptoEth"
	"api-ngmi/models"
	"errors"
)

func AccessLevelFor(user *models.User) (int, error) {
	tier, err := cryptoEth.TierFor(user.Address)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return 0, errors.New("can't get tier")
	}

	return tier, nil
}
