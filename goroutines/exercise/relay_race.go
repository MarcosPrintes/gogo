package main

import (
	"fmt"
	"time"
)

func main() {

	//create a unbufferd channel
	baton := make(chan int)

	go Runner(baton)

	time.Sleep(2000 * time.Millisecond)
	baton <- 1

	// Give the runners time to race
	time.Sleep(10000 * time.Millisecond)
}

// Runner bla bla bla
func Runner(baton chan int) {
	var newRunner int

	runner := <-baton // wait receive baton

	fmt.Println("start run")
	fmt.Printf("Runner %d Running With Baton\n", runner) //start

	if runner != 4 {
		newRunner += runner + 1
		fmt.Printf("Runner %d to the Line\n", newRunner)
		go Runner(baton)
	}

	time.Sleep(100 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("Runner %d finish, Race Over\n", runner)
		return
	}

	// Exchange the baton for the next runner
	fmt.Printf("Runner %d Exchange With Runner %d\n", runner, newRunner)
	baton <- newRunner

}
