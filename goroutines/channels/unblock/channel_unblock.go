package main

import (
	"fmt"
)

// func main() {

// 	ch := make(chan int, 100)

// 	go sendUnblock(ch)
// 	// go receiveUnblock(ch)

// 	time.Sleep(1e9)
// }

func sendUnblock(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func receiveUnblock(ch chan int) {
	for {
		fmt.Println("received => ", <-ch)
	}
}
