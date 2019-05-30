package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	harberdash "github.com.br/MarcosPrintes/twirp/pb"
)

func main() {

	client := harberdash.NewHarberDashProtobufClient("http://localhost:8080", &http.Client{})

	hat, err := client.Makehat(context.Background(), &harberdash.Size{Inches: 32})

	if err != nil {
		log.Fatal("Error client makehat => ", err.Error())
		os.Exit(1)
	}

	fmt.Println("has a hat => ", hat)

}
