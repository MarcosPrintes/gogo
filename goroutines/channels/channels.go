package main

import (
	"fmt"
)

func sendData(ch chan string) {
	fmt.Println("start sendData()")
	ch <- "Washington \n"
	ch <- "Tripoli \n"
	ch <- "London \n"
	ch <- "Beigin \n"
	ch <- "Tokio \n"
	fmt.Println("end sendData()")
}

func getData(ch chan string) {
	fmt.Println("start getData()")
	var input string
	for {
		input = <-ch
		fmt.Printf("%s", input)
	}
	fmt.Println("end getData()")
}

// func main() {
// 	ch := make(chan string)
// 	go sendData(ch)
// 	go getData(ch)
// 	time.Sleep(10 * 1e9)

// }
