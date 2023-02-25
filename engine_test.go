package chinrem

import (
	"fmt"
	"math/big"
	"testing"
)

func TestEngineVisual(t *testing.T) {

	e := NewCREngine(0)
	fmt.Println(e.String())
	e.verifyCoprimes(t)

	e = NewCREngine(5)
	fmt.Println(e.String())
	e.verifyCoprimes(t)

	e = NewCREngine(50)
	fmt.Println(e.String())
	e.verifyCoprimes(t)

}

func (e *CREngine) verifyCoprimes(t *testing.T) {

	primes := make([]*big.Int, e.size)
	for i, p := range e.primes {
		primes[i] = big.NewInt(p)
	}

	for i, cp := range e.coprimes {

		r := big.NewInt(0)

		for j, p := range primes {
			if i == j {
				if r.Mod(cp, p).Int64() != 1 {
					t.Fatal(r, cp, p)
				}
			} else {
				if r.Mod(cp, p).Int64() != 0 {
					t.Fatal(r, cp, p)
				}
			}
		}
	}
}
