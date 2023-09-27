package distributions

import (
	"math"

	"github.com/xnacly/statlib"
)

type Binomial struct {
	n float64 // number of trials
	p float64 // success probability for each trial
	q float64 // 1 - p
}

func (b *Binomial) Prob(k float64) (float64, error) {
	bin, err := statlib.BinomialCoefficient(uint64(b.n), uint64(k))
	if err != nil {
		return 0, err
	}
	return float64(bin) * math.Pow(b.p, k) * math.Pow(b.q, b.n-k), nil
}

func (b *Binomial) Median() float64 {
	return math.Floor(b.n * b.p)
}

func (b *Binomial) Mode() float64 {
	return math.Floor((b.n + 1) * b.p)
}

func (b *Binomial) Mean() float64 {
	return b.n * b.p
}

func (b *Binomial) Variance() float64 {
	return b.n * b.p * b.q
}
