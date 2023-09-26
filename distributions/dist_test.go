package distributions

import "testing"

func TestCompareFloats(t *testing.T) {
	tests := []struct {
		a float64
		b float64
		r bool
	}{
		{0, 0, true},
		{0, 1, false},
		{0, 0.0701, false},
		{1, 0.0701, false},
		{0.0701, 0.0701, true},
		{0, 5.3, false},
		{5.4, 5.3, false},
		{0.155045, 0.155045, true},
	}
	for _, i := range tests {
		r := compareFloats(i.a, i.b)
		if r != i.r {
			t.Fatalf("%f == %f should result in %t, but was %t", i.a, i.b, i.r, r)
		}
	}
}
