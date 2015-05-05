package dweetio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	BaseUri = "https://dweet.io"
)

//Read data from request body
func (api *Dweetio) ReadData(res *http.Response, dweets interface{}) (err error) {
	//Read the response body
	body, err := ioutil.ReadAll(res.Body)

	//If error
	api.ReturnError(err)

	defer res.Body.Close()

	//Unmarshal the request body
	err = json.Unmarshal(body, dweets)

	//If error
	api.ReturnError(err)

	return nil
}

//Return error
func (api *Dweetio) ReturnError(err error) (dweets interface{}, e error) {
	if err != nil {
		return nil, err
	}

	return
}

//Return the dweetio uri with the lock key if is set
func (api *Dweetio) GetUri(uri string, thing string) (dweetUri string) {
	dweetUri = fmt.Sprintf("%s%s%s", BaseUri, uri, thing)

	if api.Key != "" {
		params := url.Values{}
		params.Add("key", api.Key)
		dweetUri = fmt.Sprintf("%s?%s", dweetUri, params.Encode())
	}

	return dweetUri
}

//Return the dweetio alert uri with the lock key if is set
func (api *Dweetio) GetAlertUri(recipients []string, thing string, condition string) (dweetUri string) {
	conditionEncoded := url.QueryEscape(condition)
	dweetUri = fmt.Sprintf("%s/alert/%s/when/%s/%s", BaseUri, strings.Join(recipients, ","), thing, conditionEncoded)

	if api.Key != "" {
		params := url.Values{}
		params.Add("key", api.Key)
		dweetUri = fmt.Sprintf("%s?%s", dweetUri, params.Encode())
	}

	return dweetUri
}

