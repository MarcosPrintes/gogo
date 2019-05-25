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

	fmt.Println("Result area c2 =>", areaCircle(&c2))

	fmt.Println("structs circle 1:", c1)
	fmt.Println("structs circle 2:", c2)
	fmt.Println("structs circle 2 area:", c2.area())

	p := Person{name: "Marcos"}
	p.talk()

	a1 := Android{person: p, model: "md-124"}
	fmt.Println("android 1:", a1)

	a2 := Android2{Person: p, model: "md-121"}
	fmt.Println("android 2:", a2)

	rt1 := Rectangle{x1: 1.1, y1: 2.2, x2: 3.3, y2: 4.4}
	fmt.Println("Rectangle 1:", rt1)
	fmt.Println("Rectangle 1 areRectangle:", areaRectangle(rt1.x1, rt1.y1, rt1.x2, rt1.y2))
	fmt.Println("Rectangle 1 area:", rt1.area())

	/*
	 - we can access talk() method of Person directly when used embedded types
	 - The is-a relationship works this way intuitively: People can talk, an android is a person, therefore an android can talk.
	*/
	a2.Person.talk()
	a2.talk()

	/* ----- interface prints */
	fmt.Println("totalArea =>", totalArea(&c2, &rt1))
	fmt.Println("totalPerimeter =>", totalPerimeter(&c2, &rt1))

}

/*
============================== structs ===================================
	- A struct is a type which contains named fields
*/
type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

/* -------- functions -------*/

func areaCircle(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func areaRectangle(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x2, y2)
	w := distance(x1, y1, x2, y2)
	return l * w
}

/* ======================== Methods ===================
- special type of function

 method, replace areCircle function, (c *Circle) => receiver. creating the function in this way it allows us to call the function using the . operator
	=> circle.area()

*/
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y2)
	return l * w
}

func (c *Circle) perimeter() float64 {
	l := c.x * 4
	return l
}

func (r *Rectangle) perimeter() float64 {
	l := r.x1 * 4
	return l
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
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

/* is-a */
type Android2 struct {
	Person
	model string
}

/* ===================== Interfaces =======================
- Like a struct an interface is created using the type keyword
- But instead create fields, we create set of methods, a list of methods that a type must have to implement a interface
*/

type MyInterface interface {
	f1()
	f2()
	f3()
}

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func totalPerimeter(shapes ...Shape) float64 {
	var perimeter float64
	for _, s := range shapes {
		perimeter = s.perimeter()
	}
	return perimeter
}

/*
	Interfaces can also be used as fields
*/

type MultiShape struct {
	shapes []Shape
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}
