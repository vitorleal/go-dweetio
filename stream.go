package dweetio

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (api *Dweetio) ListenForDweetsFrom(thing string) (dweets interface{}) {
	uri := api.GetUri("/listen/for/dweets/from/", thing)
	res, err := http.Get(uri)

	//If error
	api.ReturnError(err)

	for {
		json.NewDecoder(res.Body).Decode(&dweets)
		fmt.Println(dweets)

		return dweets
	}
}
