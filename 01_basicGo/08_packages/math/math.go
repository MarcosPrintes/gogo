package math

/*
 - In Go if something starts with a capital letter that means other packages (and programs) are able to see it
 * If we had named the function "average" instead of "Average" our main program would not have been able to see it.

*/

// function average
func Average(xs []float64) float64 {
	x := 0.0
	for i := 0; i < len(xs); i++ {
		x += xs[i]
	}

	return x / float64(len(xs))
}
