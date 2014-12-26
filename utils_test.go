package dweetio

import (
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
