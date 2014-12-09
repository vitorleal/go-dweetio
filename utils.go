package dweetio

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//Read data from request body
func (api *Dweetio) ReadData(res *http.Response) (dweets interface{}, err error) {
	//Read the response body
	body, err := ioutil.ReadAll(res.Body)

	//If error
	api.ReturnError(err)

	defer res.Body.Close()

	//Unmarshal the request body
	err = json.Unmarshal(body, &dweets)

	//If error
	api.ReturnError(err)

	return dweets, nil
}

//Return error
func (api *Dweetio) ReturnError(err error) (dweets interface{}, e error) {
	if err != nil {
		return nil, err
	}

	return
}
