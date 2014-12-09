package dweetio

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	api := Dweetio{}
	res, err := api.GetDweetsFor("dontExist")
	res2, err2 := api.GetDweetsFor("vitorleal")

	if err != nil {
		fmt.Println(err.Error())
	}
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	fmt.Println(res)
	fmt.Println(res2)
}
