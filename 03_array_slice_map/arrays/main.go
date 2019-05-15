package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println("x => ", x)

	// short syntax
	z := [5]float64{3, 6, 5, 6, 7}
	fmt.Println("z (shorter sintax) => ", z)
	//definining value in each position
	var y [5]float64
	y[0] = 98
	y[1] = 93
	y[2] = 77
	y[3] = 82
	y[4] = 83

	fmt.Println("y => ", y)

	// len() => array length
	var total float64 = 0
	for i := 0; i < len(y); i++ {
		total += y[i]
	}
	fmt.Println("Total = ", total)
	/*
		len(y) is a int, total is a float => convert len(y) to float -> float64(len(y))
		In general to convert between types you use the type name like a function.
	*/

	fmt.Println("Total / 5 = ", total/float64(len(y)))

	// or
	var total2 float64 = 0
	// for i, value := range y { => error i declared but never used, won't allow you to create variables that you never use.
	for _, value := range y { // i => _
		total2 += value
	}

	fmt.Println("total2 => ", total2)
	fmt.Println("total2/len(y) => ", total2/float64(len(y)))
}
