// The mobius-go package provides methods for interacting with the Mobius Network API.
// It automates the HTTP request/response cycle, encodings, and other details needed by the API.
// This SDK lets you do everything the API lets you, in a more Go-friendly way.
//
// For further information please see the Official Mobius API documentation at
// https://mobius.network/docs/
//
//  Original Author: Francis Sunday <twitter.com/codehakase>
//
// Examples
//

package mobius

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

var Debug = false

const (
	// MobiusGoUserAgent identifies the client to the server
	MobiusGoUserAgent         = "mobius-go/1.0"
	ApiBase                   = "https://mobius.network/api/v1"
	balanceEndpoint           = "app_store/balance"
	useEndpoint               = "app_store/use"
	creditEndpoint            = "app_store/credit"
	dataFeedEndpoint          = "data_marketplace/data_feed"
	createDataEndpoint        = "data_marketplace/data_feed"
	buyEndpoint               = "data_marketplace/buy"
	registerTokenEndpoint     = "tokens/register"
	createAddressEndpoint     = "tokens/create_address"
	registerAddressEndpoint   = "tokens/register_address"
	getAddressBalanceEndpoint = "tokens/balance"
	createTransferEndpoint    = "tokens/transfer/managed"
	getTransferInfoEndpoint   = "tokens/transfer/info"
)

// Mobius defines the supported subset of the Mobius API.
type Mobius struct {
	AppStore *AppStore
	Token    *Token
	Ctx      *Mobiusimpl
}

// MobiusImpl packages data needed to interact with Mobius API
type Mobiusimpl struct {
	Client  *http.Client
	APIBase string
	APIKey  string
	AppUID  string
	Req     *httpRequest
}

// New creates a new client instance.
func New(apikey, appuid string) *Mobius {
	mip := Mobiusimpl{
		Client:  http.DefaultClient,
		APIBase: ApiBase,
		APIKey:  apikey,
		AppUID:  appuid,
	}

	m := Mobius{
		Ctx:      &mip,
		AppStore: &AppStore{mx: &mip},
	}
	return &m
}

// NewFromEnv returns a new Mobius client using the environment variables
func NewFromEnv() (*Mobius, error) {
	appUID := os.Getenv("MOBIUS_APPUID")
	apiKey := os.Getenv("MOBIUS_APIKEY")
	if appUID == "" {
		return nil, errors.New("MOBIUS_APPUID environment variable not set")
	}
	if apiKey == "" {
		return nil, errors.New("MOBIUS_APIKEY environment variable not set")
	}

	mip := Mobiusimpl{
		Client:  http.DefaultClient,
		APIBase: ApiBase,
		APIKey:  appUID,
		AppUID:  appUID,
	}

	m := Mobius{
		Ctx:      &mip,
		AppStore: &AppStore{mx: &mip},
	}
	return &m, nil
}

// generateApiUrl renders a URL for an API endpoint using the api base, and the endpoint
func generateApiUrl(m *Mobiusimpl, endpoint string) string {
	return fmt.Sprintf("%s/%s", m.APIBase, endpoint)
}

// func testt() {
// 	mob := New("xxx", "xx")
// 	mob.Ctx.Req.addHeader("hello", "hi")
// }
