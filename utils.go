package dweetio

import (
  "io/ioutil"
  "encoding/json"
  "net/http"
)

//Read data from request body
func (api *Dweetio) ReadData(resp *http.Response) (dweets interface{}, err error) {
  //Read the response body
  body, err := ioutil.ReadAll(resp.Body)

  //If error
  api.ReturnError(err)

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

