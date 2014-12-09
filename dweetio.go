package dweetio

import (
	"fmt"
	"net/http"
)

const (
	BaseUrl = "https://dweet.io"
)

//Return the current implementation version
func Version() string {
	return "0.1.0"
}

type Dweetio struct {
	Key string
}

//Get last dweet for a Thing
func (api *Dweetio) GetLastDweetFor(thing string) (dweets interface{}, err error) {
	uri := fmt.Sprintf("%s/get/latest/dweet/for/%s", BaseUrl, thing)
	res, err := http.Get(uri)

	//If error
	api.ReturnError(err)

	dweets, err = api.ReadData(res)

	//If error
	api.ReturnError(err)

	return dweets, nil
}

//Get all dweet for a Thing
func (api *Dweetio) GetDweetsFor(thing string) (dweets interface{}, err error) {
	uri := fmt.Sprintf("%s/get/dweets/for/%s", BaseUrl, thing)
	res, err := http.Get(uri)

	//If error
	api.ReturnError(err)

	dweets, err = api.ReadData(res)

	//If error
	api.ReturnError(err)

	return dweets, nil
}
