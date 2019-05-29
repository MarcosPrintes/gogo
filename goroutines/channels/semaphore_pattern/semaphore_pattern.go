package main

import (
	"fmt"
	"time"
)

type Empty interface{}

type semaphore chan Empty

func main() {

	// sem := make(chan semaphore, 10)

	ch := make(chan int, 2)

	// var smph semaphore

	// go sendValue(ch)

	// go received(ch)

	go s(ch, 2, 13)

	time.Sleep(2 * 1e9)

	r := <-ch

	fmt.Println("result => ", r)

}

func (s semaphore) P(n int) {
	empty := new(Empty)
	for i := 0; i < n; i++ {
		s <- empty
	}
}

func (s semaphore) V(n int) {
	for i := 0; i < n; i++ {
		<-s
	}
}

func s(ch chan int, a, b int) {
	ch <- a + b
}

// mutex
func (s semaphore) Lock() {
	s.P(1)
}

func (s semaphore) Unlock() {
	s.V(2)
}

// signal waits
func (s semaphore) wait(n int) {
	s.P(n)
}

func (s semaphore) Signal() {
	s.V(1)
}

func sendValue(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func received(ch chan int) {
	for {
		fmt.Println("received => ", <-ch)
	}
}

func sm2i(a, b int) int {
	return a + b
}
