package distributions

import "math"

// Geometric distribution here is the probability distribution of the number Y = X −
// 1 of failures before the first success, supported on the set {0,1,2,...}
//
// See https://en.wikipedia.org/wiki/Geometric_distribution
type Geometric struct {
	P float64 // probability of success
}

func (g *Geometric) Prob(k float64) (float64, error) {
	return g.P * math.Pow(1-g.P, k-1), nil
}

func (g *Geometric) Median() float64 {
	return math.Ceil(-1 / math.Log2(1-g.P))
}

func (g *Geometric) Mode() float64 {
	return 1
}

func (g *Geometric) Mean() float64 {
	return 1 / g.P
}

func (g *Geometric) Variance() float64 {
	return (1 - g.P) / g.P * g.P
}
