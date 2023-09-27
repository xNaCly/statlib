package distributions

import (
	"fmt"
	"testing"
)

func TestHyper(t *testing.T) {
	type Input struct {
		N      float64
		K      float64
		SmallN float64
		k      float64
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
				N:      20,
				K:      4,
				SmallN: 5,
				k:      0,
			},
			e: Expected{
				median:   0,
				mode:     1,
				variance: 0.63157894736843,
				mean:     1,
				prob:     0.28173374613004,
			},
		},
	}

	for _, e := range table {
		t.Run(fmt.Sprint(e), func(t *testing.T) {
			b := Hypergeometric{
				N:      e.i.N,
				K:      e.i.K,
				SmallN: e.i.SmallN,
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
