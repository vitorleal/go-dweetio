package dweetio

import (
	"net/http"
	"strings"
)

//Return the current implementation version
func Version() string {
	return "0.1.0"
}

//Dweetio struct
type Dweetio struct {
	Key string
}

//GET last dweet for a Thing
func (api *Dweetio) GetLastDweetFor(thing string) (dweets interface{}, err error) {
	uri := api.GetUri("/get/latest/dweet/for/", thing)
	res, err := http.Get(uri)

	//If error
	api.ReturnError(err)

	dweets, err = api.ReadData(res)

	//If error
	api.ReturnError(err)

	return dweets, nil
}

//GET all dweet for a Thing
func (api *Dweetio) GetDweetsFor(thing string) (dweets interface{}, err error) {
	uri := api.GetUri("/get/dweets/for/", thing)

	res, err := http.Get(uri)

	//If error
	api.ReturnError(err)

	dweets, err = api.ReadData(res)

	//If error
	api.ReturnError(err)

	return dweets, nil
}

//POST data to a Thing
func (api *Dweetio) DweetFor(thing string, data string) (dweet interface{}, err error) {
	uri := api.GetUri("/dweet/for/", thing)
	json := strings.NewReader(data)

	res, err := http.Post(uri, "application/json", json)

	//If error
	api.ReturnError(err)

	dweet, err = api.ReadData(res)

	return dweet, nil
}
