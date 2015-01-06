package dweetio

import (
	"fmt"
	"net/http"
)

//GET alert for a Thing
func (api *Dweetio) GetAlertFor(thing string) (dweets interface{}, err error) {
	if api.Key == "" {
		return nil, fmt.Errorf("You need a LOCKED thing to get alerts")
	}

	uri := api.GetUri("/get/alert/for/", thing)
	fmt.Println(uri)

	res, err := http.Get(uri)

	//If error
	api.ReturnError(err)

	dweets, err = api.ReadData(res)

	//If error
	api.ReturnError(err)

	return dweets, nil
}

//REMOVE alert for a Thing
func (api *Dweetio) RemoveAlertFor(thing string) (dweets interface{}, err error) {
	if api.Key == "" {
		return nil, fmt.Errorf("You need a LOCKED thing to remove alerts")
	}

	uri := api.GetUri("/remove/alert/for/", thing)
	fmt.Println(uri)

	res, err := http.Get(uri)

	//If error
	api.ReturnError(err)

	dweets, err = api.ReadData(res)

	//If error
	api.ReturnError(err)

	return dweets, nil
}

//CREATE alert for a Thing
func (api *Dweetio) SetAlertFor(thing string, recipients string, condition string) (dweet interface{}, err error) {
	if api.Key == "" {
		return nil, fmt.Errorf("You need a LOCKED thing to create alerts")
	}

	uri := api.GetAlertUri(recipients, thing, condition)
	fmt.Println(uri)

	res, err := http.Get(uri)

	//If error
	api.ReturnError(err)

	dweet, err = api.ReadData(res)

	return dweet, nil
}
