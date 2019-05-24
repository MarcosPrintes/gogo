package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("Starting the server ...")
	list, err := net.Listen("tcp", "localhost:8087")
	if err != nil {
		log.Fatal("listen error =>", err.Error())
	}

	for {
		conn, err := list.Accept()
		if err != nil {
			log.Fatal("accept connection error:", err.Error())
			return
		}
		go doStuff(conn)
	}
}

func doStuff(conn net.Conn) {
	fmt.Println("do stuff RemoteAddr", conn.RemoteAddr().String())
	for {
		buff := make([]byte, 512)
		_, err := conn.Read(buff)
		if err != nil {
			log.Fatal("doStuff error:", err.Error())
			return
		}

		fmt.Println("server received data", string(buff))

	}
}
