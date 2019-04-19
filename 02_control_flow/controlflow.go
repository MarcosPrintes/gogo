package main

import "fmt"

/* ===== control structures =========

Other programming languages have a lot of different
types of loops (while, do, until, foreach, ...) but Go only
has one that can be used in a variety of different ways

*/

func main() {

	// loop for
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("===================")
	for j := 1; j < 10; j++ {
		//if else
		if j%2 == 0 {
			fmt.Println("par", j)
		} else {
			fmt.Println("ímpar", j)
		}
	}

	s(2)
}

func s(val int) {
	//===== Switch ======
	switch val {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	default:
		fmt.Println("vish")
	}
}
