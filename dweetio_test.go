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
		t.Error("Error in the GetDweetsFor with an non existent thing")
		fmt.Println(err.Error())
	}
	if err2 != nil {
		t.Error("Error in the GetDweetsFor with existent thing")
		fmt.Println(err2.Error())
	}

	fmt.Println(res)
	fmt.Println(res2)
}
