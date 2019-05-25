package main

/* ============== GOROUTINES ================
-  A goroutine is a function that is capable of running concurrently with other functions
*/

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 5; i++ {
		fmt.Println("goroutine", n, ": ", i)
		amt := time.Duration(rand.Intn(1000))
		time.Sleep(time.Millisecond * amt)
	}
}

func main() {
	fmt.Println("concurrency")
	/*------------ goroutines ------------*/
	for i := 0; i < 5; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)

}
