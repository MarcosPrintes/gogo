package main

import (
	"encoding/gob"
	"fmt"
	"net"
)

func server() {
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println("error server", err)
		return
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println("error server for", err)
			continue
		}

		go HandleServerConnection(c)
	}
}

func HandleServerConnection(c net.Conn) {
	var msg string
	err := gob.NewDecoder(c).Decode(&msg)
	if err != nil {
		fmt.Println("error handleConnection", err)
	} else {
		fmt.Println("Received", msg)
	}

	c.Close()
}

func client() {
	c, err := net.Dial("tcp", "127.0.0.1:9999")

	if err != nil {
		fmt.Println("error client", err)
		return
	}

	msg := "Any message!"
	fmt.Println("Sending: ", msg)
	err = gob.NewEncoder(c).Encode(msg)
	if err != nil {
		fmt.Println("error client: ", err)
	}
	c.Close()
}

func main() {
	go server()
	go client()

	var input string
	fmt.Scanln(&input)
}
