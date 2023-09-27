package distributions

import (
	"fmt"
	"math"

	"github.com/xnacly/statlib"
)

// Discrete probability distribution that expresses the probability of a given
// number of events occurring in a fixed interval of time or space if these
// events occur with a known constant mean rate and independently of the time
// since the last event
//
// See https://en.wikipedia.org/wiki/Poisson_distribution
type Poisson struct {
	Lambda float64 // rate
}

func ExamplePoisson() {
	pois := Poisson{
		Lambda: 5.3,
	}

	val, err := pois.Prob(2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Chance of getting exactly 2 calls in one minute = %f%%\n", val*100)
}

func (p *Poisson) Prob(k float64) (float64, error) {
	lambdaPowerK := math.Pow(p.Lambda, k)
	kFac, err := statlib.Fac(uint64(k))
	if err != nil {
		return 0, err
	}
	eulerPowerLambda := math.Pow(math.E, -1*p.Lambda)
	return (lambdaPowerK / float64(kFac)) * eulerPowerLambda, nil
}

func (p *Poisson) Median() float64 {
	return math.Floor(p.Lambda + (1 / 3) - (1 / 50 * p.Lambda))
}

func (p *Poisson) Mode() float64 {
	return math.Floor(p.Lambda)
}

func (p *Poisson) Mean() float64 {
	return p.Lambda
}

func (p *Poisson) Variance() float64 {
	return p.Lambda
}
