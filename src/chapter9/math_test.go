package math

import (
	"testing"

	m "../chapter8/math"
)

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {
	for _, tp := range tests {
		v := m.Average(tp.values)
		if v != tp.average {
			t.Error(
				"For", tp.values,
				"expected", tp.average,
				"got", v,
			)
		}
	}
}
