package dweetio

import (
	"fmt"
	"testing"
)

//Get alert
func TestGetAlertFor(t *testing.T) {
	api := Dweetio{}
	_, err := api.GetAlertFor("godweetio") // Note: doesn't hit network

	if err == nil {
		t.Error("Error in the GetAlertFor")
	}

	fmt.Println(err)
}

func TestGetAlertForLocked(t *testing.T) {
	api := Dweetio{Key: "myLockKey"}
	res, err := api.GetAlertFor("godweetio") // Note: doesn't hit network

	if err != nil {
		t.Error("Error in the GetAlertFor")
		fmt.Println(err)
	}

	fmt.Println(res)
}

//Remove Alert
func TestRemoveAlertFor(t *testing.T) {
	api := Dweetio{}
	_, err := api.RemoveAlertFor("godweetio") // Note: doesn't hit network

	if err == nil {
		t.Error("Error in the DweetsFor")
	}

	fmt.Println(err)
}

func TestRemoveAlertForLocked(t *testing.T) {
	api := Dweetio{Key: "myLockKey"}
	res, err := api.SetAlertFor("godweetio", "vitorleal1@gmail.com", "if(dweet.alertValue > 10) return 'TEST: Greater than 10'; if(dweet.alertValue < 10) return 'TEST: Less than 10';") // Note: doesn't hit network

	if err != nil {
		t.Error("Error in the DweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}

//Set Alert
func TestSetAlertFor(t *testing.T) {
	api := Dweetio{}
	_, err := api.SetAlertFor("godweetio", "vitorleal1@gmail.com", "if(dweet.alertValue > 10) return 'TEST: Greater than 10'; if(dweet.alertValue < 10) return 'TEST: Less than 10';") // Note: doesn't hit network

	if err == nil {
		t.Error("Error in the DweetsFor")
	}

	fmt.Println(err)
}

func TestSetAlertForLocked(t *testing.T) {
	api := Dweetio{Key: "myLockKey"}
	res, err := api.RemoveAlertFor("godweetio") // Note: doesn't hit network

	if err != nil {
		t.Error("Error in the DweetsFor")
		fmt.Println(err.Error())
	}

	fmt.Println(res)
}
