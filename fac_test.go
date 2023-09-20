package statlib

import (
	"fmt"
	"testing"
)

func TestFac(t *testing.T) {
	table := []struct {
		in  uint64
		exp uint64
	}{
		{in: 0, exp: 1},
		{in: 1, exp: 1},
		{in: 2, exp: 2},
		{in: 3, exp: 6},
		{in: 4, exp: 24},
		{in: 5, exp: 120},
		{in: 6, exp: 720},
		{in: 7, exp: 5040},
		{in: 8, exp: 40320},
		{in: 9, exp: 362880},
		{in: 10, exp: 3628800},
		{in: 11, exp: 39916800},
		{in: 12, exp: 479001600},
		{in: 13, exp: 6227020800},
		{in: 14, exp: 87178291200},
		{in: 15, exp: 1307674368000},
		{in: 16, exp: 20922789888000},
		{in: 17, exp: 355687428096000},
		{in: 18, exp: 6402373705728000},
		{in: 19, exp: 121645100408832000},
	}
	for _, e := range table {
		t.Run(fmt.Sprint(e.in), func(t *testing.T) {
			r, err := Fac(e.in)
			if err != nil {
				t.Error(err)
			}
			if r != e.exp {
				t.Errorf("%d != %d\n", r, e.exp)
			}
		})
	}
}
