package main

//====== Structs and Interfaces =====

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("structs and interfaces")

	c1 := new(Circle) // this allocate memory form all fields

	// c2 := Circle{x:1, y:2, r:3}
	c2 := Circle{x: 1.1, y: 2.2, r: 3.4}

	fmt.Println("structs c1:", c1)
	fmt.Println("structs c2:", c2)
	fmt.Println("structs c2:", c2.area())

	p := Person{name: "Marcos"}
	p.talk()

	a1 := Android{person: p, model: "md-124"}
	fmt.Println("android 1:", a1)

	a2 := Android2{Person: p, model: "md-121"}
	fmt.Println("android 2:", a2)

	/*
	 - we can access talk() method of Person directly when used embedded types
	 - The is-a relationship works this way intuitively: People can talk, an android is a person, therefore an android can talk.
	*/
	a2.Person.talk()
	a2.talk()
}

/*
	- A struct is a type which contains named fields
*/
type Circle struct {
	x, y, r float64
}

func areaCicle(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

/* ===== Methods =========
- special type of function
*/
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

/*======= Embedded Types/* =========
- A struct's fields usually represent the has-a relationship (e.g Circle has a radius)
*/
type Person struct {
	name string
}

func (p *Person) talk() {
	fmt.Println("person talk: ", p.name)
}

// Go support relationshio has-a, is-a using embedded types (anonymous fields)
type Android struct {
	person Person
	model  string
}

type Android2 struct {
	Person
	model string
}

/* ===================== Interfaces =======================
- Like a struct an interface is created using the type keyword
- But instead create fields, we create methods, a method set, a list of methods that a type must have to implement a interface
*/

type Shape interface {
	area() float64
}

func totalArea(shape ...Shape) float64 {
	var area float64
	for _, v := range shape {
		area += v
	}
	return area
}
