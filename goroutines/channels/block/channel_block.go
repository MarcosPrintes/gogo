package main

import (
	"fmt"
)

/*
// ========= block =========
func main() {
	ch1 := make(chan int)

	ch1 <- 1

	go f1(ch1)

}

func f1(ch chan int) {
	fmt.Println("f1 =>", <-ch)
}


*/

// func main() {
// 	ch1 := make(chan int)

// 	go send(ch1)
// 	go receive(ch1)

// 	time.Sleep(1e9)
// }

func send(ch chan int) {
	fmt.Println("start send")
	ch <- 1
	ch <- 2
	ch <- 3
}

func receive(ch chan int) {
	for {
		fmt.Println("1 - received =>", <-ch)
	}
}
