package distributions

import (
	"fmt"
	"testing"
)

func TestExample(t *testing.T) {
	bin, err := BinomialNew(10, 0.5)
	if err != nil {
		panic(err)
	}

	val, err := bin.Prob(3)
	if err != nil {
		panic(err)
	}

	fmt.Printf("chance of getting 3 heads when throwing 10 times = %f%%\n", val*100)
}

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
		{i: Input{n: 10, p: (1 / 6.0), k: 3}, e: Expected{median: 1, mode: 1, variance: 1.38, mean: 1.6, prob: 0.155}},
	}

	for _, e := range table {
		t.Run(fmt.Sprint(e), func(t *testing.T) {
			var b Distribution
			b, err := BinomialNew(e.i.n, e.i.p)
			if err != nil {
				t.Error(err)
			}
			if !utilCompareFloats(e.e.median, b.Median()) {
				t.Errorf("Median %f != %f\n", e.e.median, b.Median())
			}
			if !utilCompareFloats(e.e.mode, b.Mode()) {
				t.Errorf("Mode %f != %f\n", e.e.mode, b.Mode())
			}
			if !utilCompareFloats(e.e.variance, b.Variance()) {
				t.Errorf("Variance %f != %f\n", e.e.variance, b.Variance())
			}
			if !utilCompareFloats(e.e.mean, b.Mean()) {
				t.Errorf("Mean %f != %f\n", e.e.mean, b.Mean())
			}
			val, err := b.Prob(e.i.k)
			if err != nil {
				panic(err)
			}
			if !utilCompareFloats(e.e.prob, val) {
				t.Errorf("Mean %f != %f\n", e.e.mean, b.Mean())
			}
		})
	}
}
