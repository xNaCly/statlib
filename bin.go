package statlib

import "github.com/xnacly/statlib/staterror"

// Computes the amount of possible ways of choosing k elements from a set of n
// values. Errors if k > n with type *staterror.StatError.
//
// Edge cases:
//   - k == n -> 1
//   - k == 0 -> 1
//
// This implementation is highly efficient by minimizing the upper
// bound of the iteration by taking advantage of the symmetry of the
// coefficient + applying the multiplicative formula.
func BinomialCoefficient(n uint64, k uint64) (uint64, error) {
	if k > n {
		return 0, staterror.New("can't compute the binomial coefficient for k > n")
	}

	if k == n || k == 0 {
		return 1, nil
	}

	// Due to the symmetry of the binomial coefficient with regard to k and n −
	// k, calculation may be optimised by setting the upper limit of the
	// product above to the smaller of k and n − k, see
	// https://en.wikipedia.org/wiki/Binomial_coefficient#Computing_the_value_of_binomial_coefficients
	kn := n - k
	if kn < k {
		k = kn
	}

	// see https://en.wikipedia.org/wiki/Binomial_coefficient#Multiplicative_formula
	var r uint64 = 1
	for i := uint64(0); i < k; i++ {
		r = r * (n - i) / (i + 1)
	}

	return r, nil
}
