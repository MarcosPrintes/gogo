package main

// ====== Function =========
import "fmt"

func main() {

	xs := []float64{12, 23, 34, 45, 56}

	fmt.Println("average => ", average(xs))

	fmt.Println("f2 => ", f2())

	x, y := f3()
	fmt.Println("return multiple values f3 => ", x, y)

	fmt.Println("variadic function => ", add(1, 2, 3, 6, 8))

	// pass a slice in variadic function
	a := []int{1, 2, 3, 4}
	fmt.Println(add(a...))

	/* ==== Clojures =======

	- can access parent scope

	*/
	y2 := "valor y"
	add2 := func(x, y int) int {
		fmt.Println("y2 => ", y2)
		return x + y
	}

	fmt.Println("clojure 1 => ", add2(1, 2))

	nextEven := makeEvenGenertor()

	fmt.Println("make even => ", nextEven())
	fmt.Println("make even => ", nextEven())
	fmt.Println("make even => ", nextEven())

	fmt.Println("fatorial ", fatorial(6))

	/*
		defer, panic, recover
		- defer move the one() to end
		- defer is often used when resources need to be freed in some way.
		- deferred functions are run even if a run-time panic occurs.
	*/
	defer one()
	two()
	three()

	/**
	 - panic recover
		panic("PANIC")
		str := recover() -> call recover never happen, panic() stop execution of the function
		fmt.Println("recover => ", str)
	*/

	defer func() {
		str := recover()
		fmt.Println("recover => ", str) // recover a parameter from panic
	}()
	panic("PANIC")
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

/* ==== Return multiple values =====
- Multiple values are often used to return an error value along with the result ( x, err := f()(type1, type2) )
- a boolean x, ok := f()(type1, bool)
*/
func f3() (int, int) {
	return 5, 7
}

/* ==== Variadic functions ===== */
func add(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

/*
	- Each time that makEvenGenerator is called it adds 2 to the local i variable which – unlike normal local variables – persists between calls.
*/
func makeEvenGenertor() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

/* ============ Recursion =========*/
func fatorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * fatorial(x-1)
}

/*========= Defer, Panic & Recover ======*/
func one() {
	fmt.Println("func one")
}
func two() {
	fmt.Println("func two")
}
func three() {
	fmt.Println("func three")
}
