package main

import "fmt"

/**===== Pointers =======
- Pointers reference a location in memory where a value is stored rather than the value itself.
	*type => pointer
	&variable => address from a variable
*/

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)

	// new(), also get a pointer, take an argument, store, return a pointer
	newXptr := new(int)
	one(newXptr)
	fmt.Println("newXptr: ", newXptr)
	y := 3
	square(&y)
	fmt.Println(y)
}

// *int pointer to value in address memory from x
func zero(xptr *int) {
	//xptr = 0 => generate a error, xptr is not a int, is location in memory which store a value
	*xptr = 0
}

func one(xptr *int) {
	*xptr = 1
}

func square(x *int) {
	*x = *x * *x
}
