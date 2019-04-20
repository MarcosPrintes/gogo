package main

/* ============== CHANNELS ================
-
*/

import (
	"fmt"
	"time"
)

/*
	- A channel type is represented with the keyword chan followed by the type of the things that are
	passed on the channel, operator	<- used to send and receive messages on the channel.

*/

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "pinger" // send msg pinger
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "ponger"
	}
}

func printer(c chan string) {
	for {
		msg := <-c // receiver msg from c channel
		fmt.Println("printer message: ", msg)
		time.Sleep(time.Second * 1)
	}
}

func pinger2(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pinger 2"
	}
}

func printer2(c <-chan string) {
	for {
		msg := <-c
		fmt.Println("printer message: ", msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	fmt.Println("channels")
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	go pinger2(c)
	go printer2(c)
	var input string
	fmt.Scanln(&input)

}
