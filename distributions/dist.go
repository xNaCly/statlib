package distributions

import "math"

type Distribution interface {
	Median() float64                 // value separating the higher half from the lower half of a data sample
	Mode() float64                   // value that appears most often in a set of data values
	Mean() float64                   // generalization of the weighted average. Informally, the expected value is the arithmetic mean of a large number of independently selected outcomes of a random variable
	Variance() float64               // variance is the expectation of the squared deviation from the mean of a random variable
	Prob(k float64) (float64, error) // computes probability
}

const room = 0.0001

// subtracts a and b, returns true if result is smaller than room
func utilCompareFloats(a float64, b float64) bool {
	return (math.Abs(a) - math.Abs(b)) < room
}
