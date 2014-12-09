package dweetio

import(
  "testing"
  "fmt"
)

func TestSet(t *testing.T) {
	api := Dweetio{}
	res, err := api.GetDweetsFor("vitorlea")

	if err != nil {
	  fmt.Println(err.Error())
  }

	fmt.Println(res)
}
