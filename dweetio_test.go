package dweetio

import (
	"fmt"
	"testing"
)

func TestDweetFor(t *testing.T) {
	api := Dweetio{}
	res, err := api.DweetFor("godweetio", "{ \"lorem\": \"ipsum\" }") // Note: doesn't hit network

	if err != nil {
		t.Error("Error in the DweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}

func TestDweetForWithKey(t *testing.T) {
	api := Dweetio{Key: "myLockKey"}
	res, err := api.DweetFor("godweetio", "{ \"lorem\": \"ipsum\" }") // Note: doesn't hit network

	if err != nil {
		t.Error("Error in the DweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}

func TestGetDweets(t *testing.T) {
	api := Dweetio{}
	res, err := api.GetDweetsFor("godweetio") // Note: doesn't hit network

	if err != nil {
		t.Error("Error in the GetDweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}

func TestGetLastDweetFor(t *testing.T) {
	api := Dweetio{}
	res, err := api.GetLastDweetFor("godweetio") // Note: doesn't hit network

	if err != nil {
		t.Error("Error in the GetLastDweetFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
