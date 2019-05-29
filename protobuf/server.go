package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"

	"github.com.br/MarcosPrintes/protobuf/PbTest"
	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("start server")
	c := make(chan *PbTest.TestMessage)
	go func() {
		message := <-c
		ReadReceivedData(message)
	}()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fata error :%s", err.Error())
		os.Exit(1)
	}

	for {
		if conn, err := listener.Accept(); err == nil {
			go handleProtoClient(conn, c)
		} else {
			continue
		}
	}
}

// ReadReceivedData bla bla
func ReadReceivedData(data *PbTest.TestMessage) {
	msgItems := data.GetMessageItems()
	fmt.Println("receiving data")
	for _, items := range msgItems {
		fmt.Println(items)
	}
}

func handleProtoClient(conn net.Conn, c chan *PbTest.TestMessage) {
	fmt.Println("Connected!")
	defer conn.Close()
	var buf bytes.Buffer
	_, err := io.Copy(&buf, conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	pdata := new(PbTest.TestMessage)

	err = proto.Unmarshal(buf.Bytes(), pdata)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
	c <- pdata
}
