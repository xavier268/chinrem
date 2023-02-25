package chinrem

import (
	"math/big"
)

// The prime slice will be extented as needed.
var primes []int64

func init() {
	primes = make([]int64, 2, 100)
	primes[0], primes[1] = 2, 3
	extendPrime(100) // precompute 100 first primes.
}

// extendPrime is used internally to grow the prime base as needed, to contain at least n primes.
func extendPrime(n int) {
	// add primes as needed
	for len(primes) < n {
		p := 2 + primes[len(primes)-1]
		for !ip(p) {
			p += 2
		}
		primes = append(primes, p)
	}
}

// INTERNAL utility, do not use externally.
// Only checks primality against known primes.
// Use only for initialization.
func ip(p int64) bool {
	for _, k := range primes {
		if p%k == 0 {
			return false
		}
		if k*k >= p {
			return true
		}
	}
	return true
}

// CRI is the main type for Chinese Remainer Integers.
type CRI struct {
	rm []int64
}

// Size provides the number of primes used as a base.
func (c *CRI) Size() int {
	return len(c.rm)
}

// Max is a big.Int represention of the products of all the primes used as a base by c.
// Only number strictly below this value can be represented.
func (c *CRI) Max() *big.Int {
	b := big.NewInt(1)
	for i := range c.rm {
		b.Mul(b, big.NewInt(primes[i]))
		//fmt.Println("DEBUG - ", b)

	}
	return b
}

func NewCRI(size int) *CRI {
	c := new(CRI)
	c.rm = make([]int64, size)
	return c
}
