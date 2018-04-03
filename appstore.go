package mobius

import (
	"fmt"
	"net/url"
)

// AppStore defines the structure to communicate with a DApp store
type AppStore struct {
	mx *Mobiusimpl
}

type UseResponse struct {
	Success    bool   `json:"success"`
	NumCredits string `json:"num_credits"`
}

type balanceResponse struct {
	NumCredits string `json:"num_credits"`
}

// Use n amount of credit from a user's balance
func (a AppStore) Use(email string, numCredits int) (*UseResponse, error) {
	r := newRequest(generateApiUrl(a.mx, useEndpoint), a.mx)
	r.addHeader("Content-Type", "application/x-www-form-urlencoded")

	data := url.Values{}
	data.Set("app_uid", a.mx.AppUID)
	data.Add("email", email)
	data.Add("num_credits", fmt.Sprintf("%s", numCredits))

	var response UseResponse
	err := postResponseFromJSON(r, data, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

// Balance returns the balance of the user with email passed
func (a AppStore) Balance(email string) (*balanceResponse, error) {
	r := newRequest(generateApiUrl(a.mx, balanceEndpoint), a.mx)
	r.addHeader("Content-Type", "application/x-www-form-urlencoded")

	data := url.Values{}
	data.Set("app_uid", a.mx.AppUID)
	data.Add("email", email)

	r.QueryValues = data

	var response balanceResponse
	err := getJSONResponse(r, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}
