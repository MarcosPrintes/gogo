package main

// ====== Arrays, Slices, Maps =========
import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println("x => ", x)

	// syntax shorter
	z := [5]float64{3, 6, 5, 6, 7}
	fmt.Println("z (shorter sintax) => ", z)

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

	/*
		Slices => segment of an array
	*/
	// create slice S
	s := make([]float64, 5, 10) // built-in make function:(array, size slice, array capacity)
	fmt.Println("s => ", s)

	/*
		create slice n [low:high] low : start  high: end
		n := [0:5]
		n := [0:] => [0:len(arr)]
		n := [:5] => [0:5]
		n := [:] => [0:len(arr)]
	*/
	arr := []float64{1, 2, 3, 4, 5}
	n := arr[2:4]

	fmt.Println(n)

	// ====== Slices functions
	// append()
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 6, 7)
	fmt.Println("append => ", slice2)

	/**
	 copy(2, 1)
		The contents of 1 are copied into 2 , but since 2 has room for only two elements only the first two elements of 1 are copied.
	*/
	slice3 := []int{1, 2, 3}
	slice4 := make([]int, 2)
	copy(slice4, slice3)
	fmt.Println("copy => ", slice3, slice4)

	/* === Maps =====
		- A map is an unordered collection of key-value pairs
	  - maps are used to look up a value by its associated key
		- can be accessed using brackets
		- like arrays and slices, maps have to be initialized to be used
		- the length of a map can change as we add new values
	*/
	// var mp map[string]int  x is a map of strings keys to ints values
	mp := make(map[string]int)
	mp["key1"] = 10
	mp["key2"] = 11
	mp["key3"] = 13
	mp["key4"] = 14
	mp["key5"] = 15
	fmt.Println(mp)
	delete(mp, "key1") // remove item from key1
	fmt.Println(mp)
	fmt.Println(mp["key7"]) // maps return zero when try access a key that dont exist

	// name, ok := mp["key7"] // Accessing an element of a map can return two values	instead of just one, first: result the lookup second: wether or not successful(boolean)
	// fmt.Println(name, ok) => 0 false

	if name, ok := mp["key2"]; ok {
		fmt.Println("result => ", name, ok)
	}

	//==== shorter way to create maps
	elements := map[string]int{
		"el1": 5,
		"el2": 9,
		"el3": 3,
		"el4": 6,
		"el5": 1,
	}
	fmt.Println("elements (map shorter ) => ", elements)

	//===== eita maps => map of strings keys to maps of strings keys to strings values. ===
	elements2 := map[string]map[string]string{
		"element1": map[string]string{
			"name":  "element 1 name",
			"state": "state element 1",
		},
		"element2": map[string]string{
			"name":  "element 2 name",
			"state": "state element2",
		},
	}

	fmt.Println("elements2 => ", elements2)

	if name, ok := elements2["element1"]; ok {
		fmt.Println("check elements2[element1] => ", name, ok)
	}
}
