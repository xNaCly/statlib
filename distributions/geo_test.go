package distributions

import (
	"fmt"
	"testing"
)

func TestGeo(t *testing.T) {
	type Input struct {
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

		{
			i: Input{
				p: 0.23,
				k: 1,
			},
			e: Expected{
				median:   3,
				mode:     1,
				variance: 0.77,
				mean:     4.34782608695653,
				prob:     0.23,
			},
		},
	}

	for _, e := range table {
		t.Run(fmt.Sprint(e), func(t *testing.T) {
			g := Geometric{
				P: e.i.p,
			}
			if !compareFloats(e.e.median, g.Median()) {
				t.Errorf("Median %.14f != %.14f\n", e.e.median, g.Median())
			}
			if !compareFloats(e.e.mode, g.Mode()) {
				t.Errorf("Mode %.14f != %.14f\n", e.e.mode, g.Mode())
			}
			if !compareFloats(e.e.variance, g.Variance()) {
				t.Errorf("Variance %.14f != %.14f\n", e.e.variance, g.Variance())
			}
			if !compareFloats(e.e.mean, g.Mean()) {
				t.Errorf("Mean %.14f != %.14f\n", e.e.mean, g.Mean())
			}
			val, err := g.Prob(e.i.k)
			if err != nil {
				panic(err)
			}
			if !compareFloats(e.e.prob, val) {
				t.Errorf("Prob %.14f != %.14f\n", e.e.prob, val)
			}
		})
	}
}
