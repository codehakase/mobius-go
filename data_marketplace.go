package mobius

import (
	"net/url"
)

// MarketPlace strutures to communicate with the Data Marketplace
type MarketPlace struct {
	mx *Mobiusimpl
}

type Feed struct {
	DataFeed struct {
		UID         string `json:"uid"`
		Name        string `json:"name"`
		Description string `json:"description"`
		ImageURL    string `json:"image_url"`
		Price       string `json:"price"`
		Descriptor  struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"descriptor"`
	} `json:"data_feed"`
	LastUpdated string `json:"last_updated,omitempty"`
}

// Get Returns DataFeed and last update timestamp, updated when new DataPoints are added.
func (m MarketPlace) Get(dataFeedUID string) (*Feed, error) {
	r := newRequest(generateApiUrl(m.mx, dataFeedEndpoint), m.mx)
	r.addHeader("Content-Type", "application/x-www-form-urlencoded")

	data := url.Values{}
	data.Set("data_feed_uid", dataFeedUID)

	r.QueryValues = data

	var response Feed
	err := getJSONResponse(r, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

// Create a new DataPoint for the DataFeed
func (m MarketPlace) Create(data interface{}) (*Feed, error) {
	r := newRequest(generateApiUrl(m.mx, createDataEndpoint), m.mx)
	r.addHeader("Content-Type", "application/json")

	var response Feed
	err := postResponseFromJSON(r, data, &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}

// Buy purchases a Data Feed and sends its data to an Ethereum Contract Address
func (m MarketPlace) Buy(dataFeedUID, address string) (*Feed, error) {
	r := newRequest(generateApiUrl(m.mx, buyEndpoint), m.mx)
	r.addHeader("Content-Type", "application/x-www-form-urlencoded")

	data := url.Values{}
	data.Set("app_uid", m.mx.AppUID)
	data.Add("data_feed_uid", dataFeedUID)
	data.Add("address", address)

	var response Feed
	err := postResponseFromJSON(r, data.Encode(), &response)
	if err == nil {
		return &response, nil
	}
	return nil, err
}
