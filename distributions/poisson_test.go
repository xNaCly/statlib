package distributions

import (
	"fmt"
	"testing"
)

func TestPoission(t *testing.T) {
	type Input struct {
		lambda float64
		x      float64
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
		{i: Input{lambda: 5.3, x: 2}, e: Expected{median: 5, mode: 5, variance: 5.3, mean: 5.3, prob: 0.070106936423}},
	}

	for _, e := range table {
		t.Run(fmt.Sprint(e), func(t *testing.T) {
			var p Distribution
			p, err := PoissonNew(e.i.lambda)
			if err != nil {
				t.Error(err)
			}
			if !compareFloats(e.e.median, p.Median()) {
				t.Errorf("Median %.14f != %.14f\n", e.e.median, p.Median())
			}
			if !compareFloats(e.e.mode, p.Mode()) {
				t.Errorf("Mode %.14f != %.14f\n", e.e.mode, p.Mode())
			}
			if !compareFloats(e.e.variance, p.Variance()) {
				t.Errorf("Variance %.14f != %.14f\n", e.e.variance, p.Variance())
			}
			if !compareFloats(e.e.mean, p.Mean()) {
				t.Errorf("Mean %.14f != %.14f\n", e.e.mean, p.Mean())
			}
			val, err := p.Prob(e.i.x)
			if err != nil {
				panic(err)
			}
			if !compareFloats(e.e.prob, val) {
				t.Errorf("Prob %.14f != %.14f\n", e.e.prob, val)
			}
		})
	}
}
