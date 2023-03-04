package chinrem

import (
	"fmt"
	"math/big"
	"math/rand"
	"strings"
)

// CRI is the main type to represent a large number, modulo the CREngine Limit.
type CRI struct {
	rm []int64
	e  *CREngine
}

// Limit is the large modulo with which all operations are conducted.
// Any CRI number is always represented modulo Limit.
func (c *CRI) Limit() *big.Int {
	return c.e.Limit()
}

// Phi is the product of all 'primes[i]-1].
// It is used to simplify exponentiation, since for any non nul number a, a^phi=1, modulo Limit.
func (c *CRI) Phi() *big.Int {
	return c.e.Phi()
}

func (c *CRI) String() string {
	sb := new(strings.Builder)
	fmt.Fprintf(sb, "%v (%v)", c.ToBig(), c.rm)
	return sb.String()
}

// Creates a new CRI representing 0.
func (e *CREngine) NewCRI() *CRI {
	c := new(CRI)
	c.rm = make([]int64, e.size)
	c.e = e
	return c
}

// Creates a CRI form an int64
func (e *CREngine) NewCRIInt64(value int64) *CRI {
	return e.NewCRI().SetInt64(value)
}

// Set c to the specified value, returning c
// c is Normalized.
func (c *CRI) SetInt64(value int64) *CRI {
	for i := range c.rm {
		pi := c.e.primes[i]
		v := value % pi
		if v < 0 {
			v = v + pi
		}
		c.rm[i] = v
	}
	return c
}

// Creates a  CRI from a big.Int
func (e *CREngine) NewCRIBig(value *big.Int) *CRI {
	return e.NewCRI().SetBig(value)
}

// Set c to the big value provided, returning c.
// c is normalized.
func (c *CRI) SetBig(value *big.Int) *CRI {
	var z big.Int
	for i := range c.rm {
		c.rm[i] = z.Mod(value, big.NewInt(c.e.primes[i])).Int64()
	}
	return c
}

// Set c to a, returning c
// No normalization is performed on c, if a was not already normalized.
func (c *CRI) Set(a *CRI) *CRI {
	copy(c.rm, a.rm)
	return c
}

// Creates a CRI from provided slice.
// Panic if length do not match.
func (e *CREngine) NewCRISlice(value []int64) *CRI {
	return e.NewCRI().SetSlice(value)
}

// Set c to the slice value. Panic if length do not match.
// C is normalized.
func (c *CRI) SetSlice(value []int64) *CRI {
	if len(value) != c.e.size {
		panic("Provided slice should match CREngine size")
	}
	copy(c.rm, value)
	c.Normalize()
	return c
}

// SameEngine checks if both CRI have same size of engine.
func SameEngine(a, b *CRI) bool {
	return a != nil && b != nil && a.e.size == b.e.size
}

// Equal compares.
// Normalization is assumed.
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
// The order defined is a total ordering that should match natural order for most small positive values.
// Normalization is assumed, but not enforced.
// Different engines size will generate a different ordering.
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
// This is the canonical form of a CRI.
func (c *CRI) Normalize() {
	for i, r := range c.rm {
		a := r % c.e.primes[i]
		if a < 0 {
			a += c.e.primes[i]
		}
		c.rm[i] = a
	}
}

// Create a new CRI by cloning an existing one.
// Not normalization occurs, if c was not already normalized.
func (c *CRI) Clone() *CRI {
	b := c.e.NewCRI()
	copy(b.rm, c.rm)
	return b
}

// Get the big.Int representation of c.
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

// Generates a random, normalized, CRI
func (e *CREngine) NewCRIRand(rd *rand.Rand) *CRI {
	return e.NewCRI().SetRandom(rd)
}

// Set a random, normalized value to the CRI
func (c *CRI) SetRandom(rd *rand.Rand) *CRI {
	for i := range c.rm {
		c.rm[i] = rd.Int63n(c.e.primes[i])
	}
	return c
}

// Clone c into another CRI, using the provided new engine, en.
// If c is smaller than both Limits, then the big.Int representation of c stays the same.
// Cloning to a larger engine is costly, cloning to a shorter engine is cheap.
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
