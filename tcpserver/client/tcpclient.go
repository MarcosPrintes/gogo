package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {

	conn, err := net.Dial("tcp", ":8087")
	if err != nil {
		log.Fatal("client connection error: ", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)

	clientName, _ := inputReader.ReadString('\n')
	fmt.Println("client name: ", clientName)

	trimmedName := strings.Trim(clientName, "\r\n")
	fmt.Println("TrimmedName: ", trimmedName)

	for {
		fmt.Println("x")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n\r")

		fmt.Println("input: ", input)
		fmt.Println("trimmedInput: ", trimmedInput)

		if trimmedInput == "Q" {
			log.Fatal("exit")
			return
		}

		//send to server
		_, err := conn.Write([]byte(trimmedName + " says " + trimmedInput))
		if err != nil {
			log.Fatal("error write", err.Error())
		}

	}
}
