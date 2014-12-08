package dweetio

import (
  "net/http"
  "fmt"
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

