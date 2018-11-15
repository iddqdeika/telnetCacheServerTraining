package main

import (
	"net"

	"homework5/Service"
)

func main() {
	service := Service.NewService()
	listener, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		panic(err)
	}
	for true {
		con, err := listener.Accept()
		if err != nil {
			println(err)
			continue
		}
		go service.Accept(con)
	}
}
