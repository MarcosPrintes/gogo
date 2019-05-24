package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8087")
	checkConnection(conn, err)

	// conn2, err := net.Dial("udp", ":8087")
	// checkConnection(conn2, err)

	// conn3, err := net.Dial("tcp", ":8087")
	// checkConnection(conn3, err)
}

func checkConnection(con net.Conn, err error) {
	if err != nil {
		log.Fatal("connection error: ", err.Error())
		os.Exit(1)
	}

	fmt.Println("Connection type => ", con)
}
