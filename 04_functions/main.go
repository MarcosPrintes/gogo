package main

// ====== Function =========
import "fmt"

func main() {

	xs := []float64{12, 23, 34, 45, 56}

	fmt.Println("average => ", average(xs))
	fmt.Println("f2 => ", f2())

}

/*
- functions are built at the top
- Each time we call a function we push it onto the call stack and each time we return from a function we
pop the last function off of the stack.

*/

func average(xs []float64) float64 {
	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))
}

func f2() (r int) {
	r = 2
	return
}
