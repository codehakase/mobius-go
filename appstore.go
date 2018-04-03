package mobius

import (
	"fmt"
	"net/url"
)

type AppStore struct {
	mx *Mobiusimpl
}

type UseResponse struct {
	Success    bool   `json:"success"`
	NumCredits string `json:"num_credits"`
}

type CreditResponse struct {
	Success    bool   `json:"success"`
	NumCredits string `json:"num_credits"`
}

type BalanceResponse struct {
	NumCredits string `json:"num_credits"`
}

// Credit sends n number of credits to user with email.
// The credits are deducted from the app itself if it has enough balance.
// Returns true if successful and false if the app did not have enough credits.
//
// email - Email of the user you want to credit
func (a AppStore) Credit(email string, numCredits int) (*CreditResponse, error) {
	r := newRequest(generateApiUrl(a.mx, creditEndpoint), a.mx)
	r.addHeader("Content-Type", "application/x-www-form-urlencoded; param=value")

	data := url.Values{}
	data.Set("app_uid", a.mx.AppUID)
	data.Add("email", email)
	data.Add("num_credits", fmt.Sprintf("%s", numCredits))

	var response CreditResponse
	err := postResponseFromJSON(r, data, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

// Use n amount of credit from a user's balance
func (a AppStore) Use(email string, numCredits int) (*UseResponse, error) {
	r := newRequest(generateApiUrl(a.mx, useEndpoint), a.mx)
	r.addHeader("Content-Type", "application/x-www-form-urlencoded; param=value")

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
func (a AppStore) Balance(email string) (*BalanceResponse, error) {
	r := newRequest(generateApiUrl(a.mx, balanceEndpoint), a.mx)
	r.addHeader("Content-Type", "application/x-www-form-urlencoded; param=value")

	data := url.Values{}
	data.Set("app_uid", a.mx.AppUID)
	data.Add("email", email)

	r.QueryValues = data

	var response BalanceResponse
	err := getJSONResponse(r, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}
