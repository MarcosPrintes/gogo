package main

import "fmt"

func main() {
	/* =========================================== Maps ======================================
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
