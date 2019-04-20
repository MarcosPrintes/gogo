package main

import (
	"fmt"

	"github.com.br/MarcosPrintes/08_packages/math"
)

func main() {
	fmt.Println("packages")
	xs := []float64{1.3, 2.2, 3.2, 4.3, 5.4, 6.5, 7.7}
	fmt.Println(math.Average(xs))
}
