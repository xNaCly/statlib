package distributions

import (
	"math"

	"github.com/xnacly/statlib"
)

// Discrete probability distribution that describes the probability of k
// successes (random draws for which the object drawn has a specified feature)
// in `SmallN` draws, without replacement, from a finite population of size `N`
// that contains exactly `K` objects with that feature, wherein each draw is
// either a success or a failure.
//
// See https://en.wikipedia.org/wiki/Hypergeometric_distribution
type Hypergeometric struct {
	N      float64 // amount of events
	K      float64 // amount successes
	SmallN float64 // draws
}

func (h *Hypergeometric) Prob(k float64) (float64, error) {
	MnCrk, err := statlib.BinomialCoefficient(uint64(h.K), uint64(k))
	if err != nil {
		return 0, err
	}

	N_MnCrn_k, err := statlib.BinomialCoefficient(uint64(h.N-h.K), uint64(h.SmallN-k))
	if err != nil {
		return 0, err
	}

	NnCrSmallN, err := statlib.BinomialCoefficient(uint64(h.N), uint64(h.SmallN))
	if err != nil {
		return 0, err
	}

	return float64(MnCrk*N_MnCrn_k) / float64(NnCrSmallN), nil
}

// doesnt exist
func (h *Hypergeometric) Median() float64 {
	return 0
}

func (h *Hypergeometric) Mode() float64 {
	return math.Floor(((h.SmallN + 1) * (h.K + 1)) / (h.N + 2))
}

func (h *Hypergeometric) Mean() float64 {
	return h.SmallN * (h.K / h.N)
}

func (h *Hypergeometric) Variance() float64 {
	return h.SmallN * (h.K / h.N) * ((h.N - h.K) / h.N) * ((h.N - h.SmallN) / (h.N - 1))
}
