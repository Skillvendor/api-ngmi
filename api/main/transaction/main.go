package transaction

import (
	"api-ngmi/cryptoEth"
	"api-ngmi/models"
	"api-ngmi/types"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func Create(w http.ResponseWriter, r *http.Request) error {
	var user models.User = r.Context().Value("user").(models.User)
	newTransaction := models.Transaction{}

	err := json.NewDecoder(r.Body).Decode(&newTransaction)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("can't decode transaction"),
		}
	}

	newTransaction.Address = user.Address
	newTransaction.State = models.Pending

	saved := newTransaction.Save()

	if !saved {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("can't save transaction"),
		}
	}

	err = json.NewEncoder(w).Encode(newTransaction)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("can't encode transaction"),
		}
	}

	return nil
}

func Show(w http.ResponseWriter, r *http.Request) error {
	var user models.User = r.Context().Value("user").(models.User)

	transaction := models.Transaction{Address: user.Address}

	found := transaction.FindPending()

	if !found {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("can't find a pending transaction"),
		}
	}

	err := json.NewEncoder(w).Encode(transaction)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error encoding transaction"),
		}
	}

	return nil
}

func Update(w http.ResponseWriter, r *http.Request) error {
	var user models.User = r.Context().Value("user").(models.User)
	transaction := models.Transaction{}

	id, err := strconv.Atoi(r.URL.Query().Get(":id"))

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("error getting id"),
		}
	}

	transaction.Id = id

	found := transaction.Find()

	if !found {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("can't find transaction"),
		}
	}

	if transaction.Address != user.Address {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("not your transaction"),
		}
	}

	status := cryptoEth.TransactionStatusFor(transaction.TransactionId)
	if status == cryptoEth.Pending {
		return &types.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        errors.New("nothing to update"),
		}
	}

	transaction.State = status

	transaction.Update()

	err = json.NewEncoder(w).Encode(transaction)

	if err != nil {
		return &types.RequestError{
			StatusCode: http.StatusInternalServerError,
			Err:        errors.New("can't encode transaction"),
		}
	}

	return nil
}
