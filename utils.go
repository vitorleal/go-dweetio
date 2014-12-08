package dweetio

import (
  "io/ioutil"
  "encoding/json"
)

//Read data from request body
func (api *Dweetio) ReadData(resp *http.Response) (dweets *Dweets) {
  //Read the response body
  body, err := ioutil.ReadAll(resp.Body)

  //If error
  if err != nil {
    fmt.Println(err)
  }

  //Unmarshal the request body
  if err := json.Unmarshal(body, &dweets); err != nil {
	  fmt.Println(err)
  }

	return dweets
}

