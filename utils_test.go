package dweetio

import (
	"fmt"
	"testing"
)

func TestGetUri(t *testing.T) {
	api := Dweetio{}
	uri := api.GetUri("/get/latest/dweet/for/", "godweetio")
	uri2 := api.GetUri("/get/dweets/for/", "godweetio")

	if uri != "https://dweet.io/get/latest/dweet/for/godweetio" {
		t.Error("Error in the GetUri")
	}

	if uri2 != "https://dweet.io/get/dweets/for/godweetio" {
		t.Error("Error in the GetUri")
	}
}

func TestGetAlertUri(t *testing.T) {
	api := Dweetio{}
	uri := api.GetAlertUri("vitorleal1@gmail.com", "godweetio", "if(dweet.temp <= 32) return \"frozen\";")

	fmt.Println(uri)

	if uri != "https://dweet.io/alert/vitorleal1@gmail.com/when/godweetio/if%28dweet.temp+%3C%3D+32%29+return+%22frozen%22%3B" {
		t.Error("Error in the GetAlertUri")
	}
}
