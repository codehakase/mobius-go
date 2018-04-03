package mobius

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type InvalidResponseErr struct {
	Expect []int
	Got    int
	URL    string
	Data   []byte
}

type httpRequest struct {
	Client      *http.Client
	URL         string
	Data        map[string]string
	Headers     map[string]string
	mx          *Mobiusimpl
	QueryValues url.Values
}

type httpResponse struct {
	Code int
	Data []byte
}

type pair struct {
	key   string
	value string
}

// expected denotes an expected list of known-good HTTP codes returned from the Mobius API
var expected = []int{202, 200}

// String formats the error message
func (e *InvalidResponseErr) String() string {
	return fmt.Sprintf("\nInvalidResponseError: Details\n URL=%s \nExpectedOneOf=%#v \nGot=%d \nError: %s", e.URL, e.Expect, e.Got, string(e.Data))
}

// Error makes the struct implement the error interface
func (e *InvalidResponseErr) Error() string {
	return e.String()
}
func newRequest(url string, m *Mobiusimpl) *httpRequest {
	return &httpRequest{URL: url, Client: http.DefaultClient, mx: m}
}

func (r *httpRequest) addHeader(key, value string) {
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}
	r.Headers[key] = value
}

func (r *httpRequest) getRequest() (*httpResponse, error) {
	return r.makeRequest("GET", nil)
}

func (r *httpRequest) postRequest(payload interface{}) (*httpResponse, error) {
	return r.makeRequest("POST", payload)
}

func (r *httpResponse) parseFromJSON(v interface{}) error {
	log.Println("unmarshaling: ", string(r.Data))
	return json.Unmarshal(r.Data, v)
}

// getJSONResponse performs a GET request, returning a JSON parsed response
func getJSONResponse(r *httpRequest, v interface{}) error {
	r.addHeader("User-Agent", MobiusGoUserAgent)

	resp, err := r.getRequest()
	if err != nil {
		return err
	}
	if !resOK(resp.Code) {
		return withErr(r.URL, expected, resp)
	}
	log.Print("get data: ", string(resp.Data))
	return resp.parseFromJSON(v)
}

// postResponseFromJSON performs a GET request, returning a JSON parsed response
func postResponseFromJSON(r *httpRequest, payload interface{}, v interface{}) error {
	r.addHeader("User-Agent", MobiusGoUserAgent)
	resp, err := r.postRequest(payload)
	if err != nil {
		return err
	}
	if !resOK(resp.Code) {
		return withErr(r.URL, expected, resp)
	}

	return resp.parseFromJSON(v)
}

func (r *httpRequest) makeRequest(method string, payload interface{}) (*httpResponse, error) {
	url, err := url.Parse(r.URL)
	if err != nil {
		return nil, err
	}

	url.RawQuery = r.QueryValues.Encode()

	var buffer io.Reader
	if payload != nil {
		buffer = strings.NewReader(payload.(string))
	} else {
		buffer = nil
	}

	req, err := http.NewRequest(method, url.String(), buffer)
	if err != nil {
		return nil, err
	}
	req.Header.Set("x-api-key", r.mx.APIKey)
	// Add any other headers
	for header, value := range r.Headers {
		req.Header.Add(header, value)
	}
	response := httpResponse{}
	resp, err := r.Client.Do(req)
	if resp != nil {
		response.Code = resp.StatusCode
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response.Data = respBody
	return &response, nil
}

// resOK searches a list of expected response codes, for a match. If found, the resposne code is considered good
func resOK(code int) bool {
	for _, i := range expected {
		if code == i {
			return true
		}
	}
	return false
}

// withErr creates a new error with specified conditions
func withErr(url string, expect []int, resp *httpResponse) error {
	return &InvalidResponseErr{
		URL:    url,
		Expect: expect,
		Got:    resp.Code,
		Data:   resp.Data,
	}
}
