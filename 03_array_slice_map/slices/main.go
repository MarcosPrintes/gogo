package main

import "fmt"

func main() {
	/*
		Slices => segment of an array
	*/
	// create slice S
	a := make([]float64, 5, 10) // built-in make function:(array, size slice, array capacity)
	fmt.Println("slice a => ", a)

	/*
		create slice n [low:high] low : start  high: end
		n := [0:5]
		n := [0:] => [0:len(arr)]
		n := [:5] => [0:5]
		n := [:] => [0:len(arr)]
	*/

	arr := []float64{1, 2, 3, 4, 6}
	fmt.Println("Array arr => ", arr)
	sliceArr := arr[2:4]
	fmt.Println("slice from array arr [2:4]", sliceArr)

	// ====== Slices functions ==========
	// append()
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 6, 7)
	fmt.Println("array slice1 => ", slice1)
	fmt.Println("slice from slice12 => ", slice2)

	/**
	  copy(2, 1)
		The contents of 1 are copied into 2 , but since 2 has room for only two elements only the first two elements of 1 are copied.
	*/
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	fmt.Println("slice4 => ", slice4)
	copy(slice4, slice3) // append to slice4 2 first elements from slice 3
	fmt.Println("slice4 => ", slice4)
	fmt.Println("copy => ", slice3, slice4)
}
