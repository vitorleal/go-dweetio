//Package dweetio is a client for the http://dweet.io api.
//
//API suport
//
//• Get last dweet for a thing
//• Get last dweet for a LOCKED thing
//• Get all dweets for a thing
//• Get all dweets for a LOCKED thing
//• Post new data for a thing
//• Post new data for a LOCKED thing
//• Set alert for a LOCKED thing
//• Get alert for a LOCKED thing
//• Remove alert for a LOCKED thing
//
//Example
//
//	import (
//		"fmt"
//		"github.com/vitorleal/go-dweetio"
//	)
//
//	func main() {
//		api := Dweetio{}
//		res, err := api.GetLastDweetFor("godweetio")
//
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//
//		fmt.Println(res)
//	}
//

package dweetio

