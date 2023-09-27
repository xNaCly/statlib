package distributions

import (
	"math"

	"github.com/xnacly/statlib"
)

type Poisson struct {
	lambda float64 // rate
}

func (p *Poisson) Prob(k float64) (float64, error) {
	lambdaPowerK := math.Pow(p.lambda, k)
	kFac, err := statlib.Fac(uint64(k))
	if err != nil {
		return 0, err
	}
	eulerPowerLambda := math.Pow(math.E, -1*p.lambda)
	return (lambdaPowerK / float64(kFac)) * eulerPowerLambda, nil
}

func (p *Poisson) Median() float64 {
	return math.Floor(p.lambda + (1 / 3) - (1 / 50 * p.lambda))
}

func (p *Poisson) Mode() float64 {
	return math.Floor(p.lambda)
}

func (p *Poisson) Mean() float64 {
	return p.lambda
}

func (p *Poisson) Variance() float64 {
	return p.lambda
}
