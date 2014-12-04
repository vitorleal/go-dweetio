package main

import (
  "net/http"
  "io/ioutil"
  "encoding/json"
  "fmt"
)

const (
  BaseUrl = "https://dweet.io"
)

type Dweetio struct {
  Key string
}

type Dweets struct {
	This string
	By   string
	The  string
	With []struct {
		Thing   string
		Created string
		Content map[string]interface{}
	}
}

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


//Get last dweet for a Thing
func (api *Dweetio) GetLastDweetFor(thing string) (dweets *Dweets, err error) {
  uri := fmt.Sprintf("%s/get/latest/dweet/for/%s", BaseUrl, thing)
  resp, err := http.Get(uri)
  defer resp.Body.Close()

  //If error
  if err != nil {
    return nil, err
  }

  dweets = api.ReadData(resp)

	return dweets, nil
}

//Get all dweet for a Thing
func (api *Dweetio) GetDweetsFor(thing string) (dweets *Dweets, err error) {
  uri := fmt.Sprintf("%s/get/dweets/for/%s", BaseUrl, thing)
  resp, err := http.Get(uri)
  defer resp.Body.Close()

  //If error
  if err != nil {
    return nil, err
  }

  dweets = api.ReadData(resp)

	return dweets, nil
}


//Main func
func main() {
	api := Dweetio{}
	res, err := api.GetDweetsFor("vitorleal")

  if err != nil {
    fmt.Println(err)
	}

	fmt.Println(res)
}
