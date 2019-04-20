package main

import "fmt"

/**

===== Integer ====
int	    =>
	 - int8
	 - int16
	 - int32
	 - int64
uint    =>
	- uint8
	- uint16
	- uint32
	- uint64

	uintptr => large enough to hold the bit pattern of any pointer.

==== Floating point ===
- values which can be represented: “not a number” ( NaN , for things like 0/0 ) and
  positive and negative infinity. ( +∞ and −∞ )
- float 32
- float64

- Two additional types for representing complex numbers (numbers with imaginary parts)
complex64
complex128

**/

/* Go is lexically scoped using blocks */

var a = "valor de a" // global scope, can be accessed by any function

func main() {

	//print
	fmt.Println(len("print golang"))

	// sum integer , floating point
	fmt.Println("Um coelho", 1.1+1.2)

	//strings, acces a position (return a byte)
	fmt.Println("migor"[0])

	// boolean, a special 1 bit integer ,true or false (on or off)
	// operators && and, || or, ! not]
	fmt.Println(true && true)

	//variables
	// name := value -> var name = value
	var y string = "valor y"
	x := "valor x"
	fmt.Println(x)
	fmt.Println(y)

	f() // has access to "a"

	// constants
	const c = "constant val"

	//multiple variables
	var (
		p = 5
		q = 6
		r = 7
	)

	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)

	fmt.Println("enter:")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}

func f() {
	fmt.Println(a)
}
