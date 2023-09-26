package distributions

import (
	"fmt"
	"testing"
)

func TestBinomial(t *testing.T) {
	type Input struct {
		n uint64
		p float64
		k float64
	}
	type Expected struct {
		median   float64
		mode     float64
		mean     float64
		variance float64
		prob     float64
	}
	table := []struct {
		i Input
		e Expected
	}{

		{i: Input{n: 10, p: (1 / 6.0), k: 3}, e: Expected{median: 1, mode: 1, variance: 1.388889, mean: 1.666667, prob: 0.1550453596}},
	}

	for _, e := range table {
		t.Run(fmt.Sprint(e), func(t *testing.T) {
			var b Distribution
			b, err := BinomialNew(e.i.n, e.i.p)
			if err != nil {
				t.Error(err)
			}
			if !compareFloats(e.e.median, b.Median()) {
				t.Errorf("Median %.14f != %.14f\n", e.e.median, b.Median())
			}
			if !compareFloats(e.e.mode, b.Mode()) {
				t.Errorf("Mode %.14f != %.14f\n", e.e.mode, b.Mode())
			}
			if !compareFloats(e.e.variance, b.Variance()) {
				t.Errorf("Variance %.14f != %.14f\n", e.e.variance, b.Variance())
			}
			if !compareFloats(e.e.mean, b.Mean()) {
				t.Errorf("Mean %.14f != %.14f\n", e.e.mean, b.Mean())
			}
			val, err := b.Prob(e.i.k)
			if err != nil {
				panic(err)
			}
			if !compareFloats(e.e.prob, val) {
				t.Errorf("Prob %.14f != %.14f\n", e.e.prob, val)
			}
		})
	}
}
