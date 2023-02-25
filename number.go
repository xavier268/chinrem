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

// extendPrime is used internally to grow the prime database as needed, to contain at least n primes.
func extendPrime(n int) {

	// add primes as needed
	for len(primes) < n {
		p := 2 + primes[len(primes)-1]
		pr := false
		for !pr {
			pr = true
			for _, k := range primes {
				if p%k == 0 {
					pr = false
					p += 2
					break
				}
				if k*k >= p {
					pr = true
					break
				}
			}
		}
		primes = append(primes, p)
	}
}

// CRI is the main type for Chinese Remainer Integers.
type CRI struct {
	rm []int64
}

// Size provides the number of primes used as a base.
func (c *CRI) Size() int {
	return len(c.rm)
}

// Limit is a big.Int represention of the products of all the primes used as a base by c.
// All numbers are represented modulo this large integer.
func (c *CRI) Limit() *big.Int {
	b := big.NewInt(1)
	for i := range c.rm {
		b.Mul(b, big.NewInt(primes[i]))
	}
	return b
}

func NewCRI(size int) *CRI {
	c := new(CRI)
	c.rm = make([]int64, size)
	return c
}

func NewCRIInt64(size int, value int64) *CRI {
	extendPrime(size)
	c := new(CRI)
	c.rm = make([]int64, size)
	for i := range c.rm {
		c.rm[i] = value % primes[i]
	}
	return c
}

func NewCRIBig(size int, value *big.Int) *CRI {
	extendPrime(size)
	var z big.Int
	c := new(CRI)
	c.rm = make([]int64, size)
	for i := range c.rm {
		c.rm[i] = z.Mod(value, big.NewInt(primes[i])).Int64()
	}
	return c
}

func NewCRISlice(value []int64) *CRI {
	extendPrime(len(value))
	c := new(CRI)
	c.rm = make([]int64, len(value))
	copy(c.rm, value)
	return c
}

// Equal compares, assuming canonical form - see Normalize.
func (c *CRI) Equal(d *CRI) bool {
	if d == nil {
		return false
	}
	if len(c.rm) != len(d.rm) {
		return false
	}
	for i := range d.rm {
		if d.rm[i] != c.rm[i] {
			return false
		}
	}
	return true
}

// Normalize brings each modulo between 0 and p(i)-1.
// This is the canonical form aof a CRI.
func (c *CRI) Normalize() {
	for i, r := range c.rm {
		c.rm[i] = r % primes[i]
	}
}

// Minus changes the sign of c
func (c *CRI) Minus() {
	for i, r := range c.rm {
		c.rm[i] = (-r) % primes[i]
	}
}

func (c *CRI) Clone() *CRI {
	b := NewCRI(c.Size())
	copy(b.rm, c.rm)
	return b
}

func (c *CRI) ToBig() *big.Int {

	panic("to do")

}
