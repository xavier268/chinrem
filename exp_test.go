package chinrem

import (
	"math/big"
	"math/rand"
	"testing"
)

func TestExpInternal(t *testing.T) {

	for a := int64(0); a < 30; a++ {
		for b := int64(0); b < 15; b++ {
			if a != 0 && b != 0 {
				for m := int64(2); m < 10; m++ {
					r := expi(a, b, m)
					aa, bb, mm, rr := big.NewInt(a), big.NewInt(b), big.NewInt(m), big.NewInt(r)
					aa.Exp(aa, bb, mm)
					if aa.Cmp(rr) != 0 {
						t.Fatalf("(expi)Got : %v^%v=%v[%v]\tWanted : %v^%v=%v[%v]\n", a, b, r, m, a, b, aa, m)
					}

				}
			}
		}
	}

}

func TestExp(t *testing.T) {
	e := NewCREngine(10)
	rd := rand.New(rand.NewSource(42))

	for i := 0; i < 1000; i++ {

		a := e.NewCRIRand(rd)
		n := big.NewInt(rd.Int63n(6546546))
		r := e.NewCRI().Exp(a, n)

		aa := a.ToBig()
		rr := r.ToBig()
		rtrue := big.NewInt(0).Exp(aa, n, a.Limit())

		if rtrue.Cmp(rr) != 0 {
			t.Fatalf("Got %v^%v=%v[%v], wanted %v^%v=%v[%v]\n", aa, n, rr, a.Limit(), aa, n, rtrue, a.Limit())
		}

	}
}
