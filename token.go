package mobius

import (
	"fmt"
	"net/url"
)

type Token struct {
	mx      *Mobiusimpl
	TokenID string `json:"token_uid"`
}

type TokenData struct {
	UID       string `json:"uid,omitempty"`
	TokenType string `json:"token_type"`
	Name      string `json:"name"`
	Symbol    string `json:"symbol"`
	Issuer    string `json:"issuer"`
}

type CreateAddressResp struct {
	UID     string `json:"uid"`
	Address string `json:"address"`
}

type RegisterAddressResp struct {
	UID string `json:"uid"`
}

type GetBalanceResp struct {
	Address string    `json:"address"`
	Balance string    `json:"balance"`
	Token   TokenData `json:"token"`
}

type CreateTransferResp struct {
	TokenAddressUID string `json:"token_address_transfer_uid"`
}

type TokenTransferInfo struct {
	UID    string `json:"uid"`
	Status string `json:"status"`
	TxHash string `json:"tx_hash"`
}

// Register a Token
func (t Token) Register(tokenType, name, symbol, address string) (*TokenData, error) {
	if tokenType == "" {
		tokenType = "ERC20"
	}

	r := newRequest(generateApiUrl(t.mx, registerTokenEndpoint), t.mx)
	r.addHeader("Content-Type", "application/x-www-form-urlencoded; param=value")

	data := url.Values{}
	data.Set("token_type", tokenType)
	data.Add("name", name)
	data.Add("symbol", symbol)
	data.Add("address", address)

	var response TokenData
	err := postResponseFromJSON(r, data.Encode(), &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

// CreateAddress creates an address for the token specified
func (t Token) CreateAddress(tokenUID string) (*CreateAddressResp, error) {
	r := newRequest(generateApiUrl(t.mx, createAddressEndpoint), t.mx)
	r.addHeader("Content-Type", "application/x-www-form-urlencoded; param=value")

	data := url.Values{}
	data.Set("token_uid", tokenUID)

	var response CreateAddressResp
	err := postResponseFromJSON(r, data.Encode(), &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

// RegisterAddress registers an address for a token
func (t Token) RegisterAddress(tokenUID, address string) (*RegisterAddressResp, error) {
	r := newRequest(generateApiUrl(t.mx, registerAddressEndpoint), t.mx)
	r.addHeader("Content-Type", "application/x-www-form-urlencoded; param=value")

	data := url.Values{}
	data.Set("token_uid", tokenUID)
	data.Set("address", address)

	var response RegisterAddressResp
	err := postResponseFromJSON(r, data.Encode(), &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

// GetAddressBalance queries the number of tokens specified for its balance
func (t Token) GetAddressBalance(tokenUID, address string) (*GetBalanceResp, error) {
	r := newRequest(generateApiUrl(t.mx, getAddressBalanceEndpoint), t.mx)
	r.addHeader("Content-Type", "application/x-www-form-urlencoded; param=value")

	data := url.Values{}
	data.Set("token_uid", tokenUID)
	data.Set("address", address)

	r.QueryValues = data

	var response GetBalanceResp
	err := getJSONResponse(r, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

// CreateTransfer tranfers tokens form a Mobius managed address to a specified address
func (t Token) CreateTransfer(tokenAddressUID, addressTo string, numTokens int) (*CreateTransferResp, error) {
	r := newRequest(generateApiUrl(t.mx, createTransferEndpoint), t.mx)
	r.addHeader("Content-Type", "application/x-www-form-urlencoded; param=value")

	data := url.Values{}
	data.Set("token_address_uid", tokenAddressUID)
	data.Set("address_to", addressTo)
	data.Set("num_tokens", fmt.Sprintf("%s", numTokens))

	var response CreateTransferResp
	err := postResponseFromJSON(r, data.Encode(), &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

// GetTransferInfo queries the number of tokens specified for its balance
func (t Token) GetTransferInfo(tokenTransferID string) (*TokenTransferInfo, error) {
	r := newRequest(generateApiUrl(t.mx, getTransferInfoEndpoint), t.mx)
	r.addHeader("Content-Type", "application/x-www-form-urlencoded; param=value")

	data := url.Values{}
	data.Set("token_address_transfer_uid", tokenTransferID)

	r.QueryValues = data

	var response TokenTransferInfo
	err := getJSONResponse(r, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

// GetTokenUID returns the uid of a token
func (td TokenData) GetTokenUID() string {
	return td.UID
}
