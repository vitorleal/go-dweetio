package dweetio

import (
	"fmt"
	"net"
)

func (api *Dweetio) ListenForDweetsFrom(thing string) {
	uri := api.GetUri("/listen/for/dweets/from/", thing)

	stream, err := net.Listen("tcp", uri)

	api.ReturnError(err)

	for {
		conn, err := stream.Accept()

		if err != nil {
			fmt.Println()
		}

		fmt.Println(conn)
	}
}
