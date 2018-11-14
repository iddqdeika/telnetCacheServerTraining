package main

import (
	"net"
)

func main(){
	service := GetNewService()
	listener, err := net.Listen("tcp","0.0.0.0:9999")
	if err!=nil{
		panic(err)
	}
	for true{
		con, err := listener.Accept()
		if err!=nil{
			println(err)
			continue
		}
		service.Accept(con)
	}
}

