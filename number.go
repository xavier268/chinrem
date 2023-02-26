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
	fmt.Fprintf(sb, "%v (%v)", c.ToBig(), c.rm)
	return sb.String()
}

func (e *CREngine) NewCRI() *CRI {
	c := new(CRI)
	c.rm = make([]int64, e.size)
	c.e = e
	return c
}

// Creates a Normalized CRI
func (e *CREngine) NewCRIInt64(value int64) *CRI {

	c := e.NewCRI()
	for i := range c.rm {
		v := value % e.primes[i]
		if v < 0 {
			v = v + e.primes[i]
		}
		c.rm[i] = v
	}
	return c
}

// Creates a Normalized CRI
func (e *CREngine) NewCRIBig(value *big.Int) *CRI {
	var z big.Int
	c := e.NewCRI()
	for i := range c.rm {
		c.rm[i] = z.Mod(value, big.NewInt(e.primes[i])).Int64()
	}
	return c
}

// Creates a Normalized CRI
func (e *CREngine) NewCRISlice(value []int64) *CRI {
	if len(value) != e.size {
		panic("Provided slice should match CREngine size")
	}
	c := e.NewCRI()
	copy(c.rm, value)
	c.Normalize()
	return c
}

// SameEngine checks if both CRI are pointing to the same CREngine, or engine having same size.
func SameEngine(a, b *CRI) bool {
	return a != nil && b != nil && a.e.size == b.e.size
}

// Equal compares, assuming normalized form - see Normalize.
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

// Cmp compares x and y and returns:
//
//	-1 if c < a
//	 0 if c == a
//	+1 if c > a
//
// The order defined is a total ordering that should match natural order for small positive values.
// Normalization is assumed and not checked.
// Different engines size will appear as different numbers.
func (c *CRI) Cmp(a *CRI) int {

	if !SameEngine(a, c) { // sensible values if not same size, to avoid equality.
		if c.e.size-a.e.size > 0 {
			return +1
		} else {
			return -1
		}
	}

	for i := len(c.rm) - 1; i >= 0; i-- {
		switch {
		case c.rm[i] > a.rm[i]:
			return +1
		case c.rm[i] < a.rm[i]:
			return -1
		default: // loop if equal ...}
		}
	}
	return 0
}

// Normalize brings each modulo between 0 and p(i)-1.
// This is the canonical form aof a CRI.
func (c *CRI) Normalize() {
	for i, r := range c.rm {
		a := r % c.e.primes[i]
		if a < 0 {
			a += c.e.primes[i]
		}
		c.rm[i] = a
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

// Clone c into another CRI, using the provided new engine, en.
// c is unchanged. Cloning to a larger engine is costly, to a shorter engine is cheap.
func (c *CRI) CloneE(en *CREngine) *CRI {
	if en.size == c.e.size { // same sized engine, no change.
		return c.Clone()
	}

	if en.size < c.e.size { // truncating
		cc := en.NewCRI()
		copy(cc.rm, c.rm[:en.size])
		return cc
	}

	// extending. Convert to a big as an intermediate value.
	return en.NewCRIBig(c.ToBig())

}
