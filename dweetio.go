package dweetio

import (
	"net/http"
	"strings"
	"time"
)

//Return the current implementation version
func Version() string {
	return "0.1.0"
}

//DweetIO struct
type Dweetio struct {
	Key string
}

//Dynamic JSON struct
type JSON struct {
	Thing   string
	Created time.Time
	Content interface{}
}

//Base Dweet struct
type Base struct {
	This string
	By   string
	The  string
}

//Miltiple Dweets struct
type Dweets struct {
	Base
	With []JSON
}

//Single Dweet struct
type Dweet struct {
	Base
	With JSON
}

//GET last dweet for a Thing
func (api *Dweetio) GetLastDweetFor(thing string) (dweets *Dweets, err error) {

	uri := api.GetUri("/get/latest/dweet/for/", thing)
	res, err := http.Get(uri)

	err = api.ReadData(res, &dweets)

	//If error
	api.ReturnError(err)

	return dweets, nil
}

//GET all dweet for a Thing
func (api *Dweetio) GetDweetsFor(thing string) (dweets *Dweets, err error) {
	uri := api.GetUri("/get/dweets/for/", thing)

	res, err := http.Get(uri)

	err = api.ReadData(res, &dweets)

	//If error
	api.ReturnError(err)

	return dweets, nil
}

//POST data to a Thing
func (api *Dweetio) DweetFor(thing string, data string) (dweet *Dweet, err error) {
	uri := api.GetUri("/dweet/for/", thing)
	json := strings.NewReader(data)

	res, err := http.Post(uri, "application/json", json)

	//If error
	api.ReturnError(err)

	err = api.ReadData(res, &dweet)

	return dweet, nil
}

