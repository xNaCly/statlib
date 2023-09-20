package statlib

import (
	"github.com/xnacly/statlib/staterror"
)

// Returns the factorial of n, errors for n>=20 with error of type
// *staterror.StatError.
//
// All values are hardcoded in a switch case to omit the
// computation heavy process of calculating factorials.
func Fac(n uint64) (r uint64, e error) {
	if n >= 20 {
		e = staterror.New("can't compute factorials for n >= 20")
		return
	}

	switch n {
	case 0:
		r = 1
	case 1:
		r = 1
	case 2:
		r = 2
	case 3:
		r = 6
	case 4:
		r = 24
	case 5:
		r = 120
	case 6:
		r = 720
	case 7:
		r = 5_040
	case 8:
		r = 40_320
	case 9:
		r = 362_880
	case 10:
		r = 3_628_800
	case 11:
		r = 39_916_800
	case 12:
		r = 479_001_600
	case 13:
		r = 6_227_020_800
	case 14:
		r = 87_178_291_200
	case 15:
		r = 1_307_674_368_000
	case 16:
		r = 20_922_789_888_000
	case 17:
		r = 355_687_428_096_000
	case 18:
		r = 6_402_373_705_728_000
	case 19:
		r = 121_645_100_408_832_000
	}

	return r, nil
}
