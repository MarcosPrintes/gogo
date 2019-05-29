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

	// resourcePool := 3
	// sem := make(chan bool, resourcePool) // bufer

	// for i := 0; i < 10; i++ {
	// 	// pause for the scheduler
	// 	time.Sleep(100 * time.Millisecond)
	// 	// acquire a single resource
	// 	sem <- true
	// 	go printIt(sem, "i Test ", i)
	// }

	// // acquire the resources back; then we know it's done
	// for j := 0; j < resourcePool; j++ {
	// 	fmt.Println("j => ", j)
	// 	sem <- true
	// }

	ch := make(chan int, 2)

	go compute(ch)

	time.Sleep(5 * 1e9)

	// ch3 := <-ch
	<-ch

	// fmt.Println("ch3", ch3)
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
