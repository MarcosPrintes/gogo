package math

import "testing"

type testpair struct {
	values  []float64
	average float64
}

var avs = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{5, 5}, 2},
	{[]float64{7, 42}, 112},
	{[]float64{1, 89}, 2},
}

func TestAverage(t *testing.T) {
	var v float64
	for _, pair := range avs {
		v = Average(pair.values)
		if v != pair.average {
			t.Error(
				"for: ", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}

}
