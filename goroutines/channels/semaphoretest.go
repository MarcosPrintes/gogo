// Idiomatic Semaphore Example in Go
// Lucas Wagner

// Golang has no built-in facility to implement semaphores, so a common design
// pattern is to use buffered channels.

package main

import (
	"fmt"
	"time"
)

func printIt(sem chan bool, msg string, iter int) {

	// release the resource when we're done
	defer func() { <-sem }()

	fmt.Println(msg, " ", iter)

}

func main() {

	ch := make(chan int, 2)

	go compute(ch)

	time.Sleep(5 * 1e9)

	ch3 := <-ch
	<-ch

	fmt.Println("ch3", ch3)
}

func compute(ch chan int) {
	for i := 0; ; i++ {
		ch <- i // when it complete, signal on the channel
		fmt.Println("ch ", <-ch)
	}
}

func g2(c chan int) {
	for {
		fmt.Println("c => ", <-c)
	}
}
