package dweetio

import(
  "testing"
  "fmt"
)

func TestSet(t *testing.T) {
	api := Dweetio{}
	res, err := api.GetDweetsFor("vitorleal")

	if err != nil {
	  fmt.Println(err)
  }

	fmt.Println(res)
}
