package distributions

import (
	"math"
)

type Distribution interface {
	Median() float64                 // value separating the higher half from the lower half of a data sample
	Mode() float64                   // value that appears most often in a set of data values
	Mean() float64                   // generalization of the weighted average. Informally, the expected value is the arithmetic mean of a large number of independently selected outcomes of a random variable
	Variance() float64               // variance is the expectation of the squared deviation from the mean of a random variable
	Prob(k float64) (float64, error) // computes probability
}

const epsilon = 1e-6

// returns true if equal within an epsilon of dist.epsilon
func compareFloats(a float64, b float64) bool {
	if a == b {
		return true
	}
	p := math.Abs(a) - math.Abs(b)
	return p < epsilon && p > 0
}
