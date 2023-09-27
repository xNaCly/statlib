package distributions

import (
	"math"

	"github.com/xnacly/statlib"
)

// Discrete probability distribution of the number of successes in a sequence
// of n independent experiments, each asking a yes–no question, and each with
// its own Boolean-valued outcome: success (with probability p) or failure (1-p).
//
// requires: 0 < P < 1
//
// See https://en.wikipedia.org/wiki/Binomial_distribution
type Binomial struct {
	N float64 // number of trials
	P float64 // success probability
	Q float64 // failure probability
}

func (b *Binomial) Prob(k float64) (float64, error) {
	bin, err := statlib.BinomialCoefficient(uint64(b.N), uint64(k))
	if err != nil {
		return 0, err
	}
	return float64(bin) * math.Pow(b.P, k) * math.Pow(b.Q, b.N-k), nil
}

func (b *Binomial) Median() float64 {
	return math.Floor(b.N * b.P)
}

func (b *Binomial) Mode() float64 {
	return math.Floor((b.N + 1) * b.P)
}

func (b *Binomial) Mean() float64 {
	return b.N * b.P
}

func (b *Binomial) Variance() float64 {
	return b.N * b.P * b.Q
}
