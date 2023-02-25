package chinrem

import (
	"fmt"
	"math/big"
	"math/rand"
	"strings"
)

// CRI is the main type for Chinese Remainer Integers.
type CRI struct {
	rm []int64
	e  *CREngine
}

func (c *CRI) String() string {
	sb := new(strings.Builder)
	fmt.Fprintf(sb, "%#v\n", c)
	return sb.String()
}

func (e *CREngine) NewCRI() *CRI {
	c := new(CRI)
	c.rm = make([]int64, e.size)
	c.e = e
	return c
}

func (e *CREngine) NewCRIInt64(value int64) *CRI {

	c := e.NewCRI()
	for i := range c.rm {
		c.rm[i] = value % e.primes[i]
	}
	return c
}

func (e *CREngine) NewCRIBig(value *big.Int) *CRI {
	var z big.Int
	c := e.NewCRI()
	for i := range c.rm {
		c.rm[i] = z.Mod(value, big.NewInt(e.primes[i])).Int64()
	}
	return c
}

func (e *CREngine) NewCRISlice(value []int64) *CRI {
	if len(value) != e.size {
		panic("Provided slice should match CREngine size")
	}
	c := e.NewCRI()
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
		c.rm[i] = r % c.e.primes[i]
	}
}

// Minus changes the sign of c
func (c *CRI) Minus() {
	for i, r := range c.rm {
		c.rm[i] = (-r) % c.e.primes[i]
	}
}

func (c *CRI) Clone() *CRI {
	b := c.e.NewCRI()
	copy(b.rm, c.rm)
	return b
}

func (c *CRI) ToBig() *big.Int {

	b := big.NewInt(0)
	bb := big.NewInt(0)

	for i, cp := range c.e.coprimes {
		bb.Mul(big.NewInt(c.rm[i]), cp)
		b.Add(b, bb)
		b.Mod(b, c.e.limit)
	}

	return b

}

func (e *CREngine) NewCRIRand(rd *rand.Rand) *CRI {
	c := e.NewCRI()
	for i := range c.rm {
		c.rm[i] = rd.Int63n(c.e.primes[i])
	}
	return c
}
