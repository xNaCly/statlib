package statlib

import (
	"fmt"
	"testing"
)

func TestBinomialCoefficient(t *testing.T) {
	type input struct {
		n uint64
		k uint64
	}
	table := []struct {
		in  input
		exp uint64
	}{
		{in: input{n: 1, k: 1}, exp: 1},
		{in: input{n: 1, k: 0}, exp: 1},
		{in: input{n: 6, k: 3}, exp: 20},
		{in: input{n: 49, k: 6}, exp: 13_983_816},
		{in: input{n: 6, k: 5}, exp: 6},
		{in: input{n: 10, k: 5}, exp: 252},
	}
	for _, e := range table {
		t.Run(fmt.Sprint(e.in), func(t *testing.T) {
			r, err := BinomialCoefficient(e.in.n, e.in.k)
			if err != nil {
				t.Error(err)
			}
			if r != e.exp {
				t.Errorf("%d != %d\n", r, e.exp)
			}
		})
	}
}
