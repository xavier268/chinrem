package chinrem

import "fmt"

func (c *CRI) IsZero() bool {
	for _, r := range c.rm {
		if r != 0 {
			return false
		}
	}
	return true
}

func (c *CRI) IsOne() bool {
	for _, r := range c.rm {
		if r != 1 {
			return false
		}
	}
	return true
}

// Minus changes the sign of a, store the result in c, returning c
func (c *CRI) Minus(a *CRI) *CRI {
	for i, r := range c.rm {
		c.rm[i] = (-r) % a.e.primes[i]
	}
	return c
}

// Add a+b, storing result in c, returning c.
func (c *CRI) Add(a, b *CRI) *CRI {
	for i, p := range c.e.primes {
		c.rm[i] = (a.rm[i] + c.rm[i]) % p
	}
	return c
}

// Mul a*b, storing result in c, returning c.
func (c *CRI) Mul(a, b *CRI) *CRI {
	for i, p := range c.e.primes {
		c.rm[i] = (a.rm[i] * c.rm[i]) % p
	}
	return c
}

// Exp computes  a^b modulo limit, storing result in c, returning c.
func (c *CRI) Exp(a, b *CRI) *CRI {
	panic("todo")
}

// utility that returns g as the gcd of a and b, and u,v such that au + bv = g.
// by convention, gcd(0,0) = 0
// and gcd(0,a) = a
func gcd(a, b int64) (g, u, v int64) {

	// TODO - handle 0 or panic ?

	g, u, v = a, 1, 0
	var gg, uu, vv int64 = b, 0, 1

	for gg != 0 {
		q := g / gg
		g, u, v, gg, uu, vv = gg, uu, vv, g-q*gg, u-q*uu, v-q*vv
	}

	if g < 0 {
		g, u, v = -g, -u, -v
	}

	return g, u, v
}

var ErrNotInversible = fmt.Errorf("not inversible")

// Compute the inverse of a modulo Limit, store result in c.
// If no inverse can be found, return ErrNotInversible.
func (c *CRI) Inv(a *CRI) error {

	for i, r := range a.rm {
		p := a.e.primes[i]
		g, u, _ := gcd(r, p)
		if g != 1 {
			return ErrNotInversible
		}
		u = u % p
		if u < 0 {
			u = u + p
		}
		c.rm[i] = u
	}
	return nil
}
