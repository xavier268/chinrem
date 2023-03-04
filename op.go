package chinrem

import (
	"fmt"
	"math/big"
)

// Test is c is zero, modulo Limit.
func (c *CRI) IsZero() bool {
	for _, r := range c.rm {
		if r != 0 {
			return false
		}
	}
	return true
}

// Test if c is 1 modulo Limit
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
		c.rm[i] = (a.rm[i] * b.rm[i]) % p
	}
	return c
}

// utility that returns g as the gcd of a and b, and u,v such that au + bv = g, using Euclid algorithm.
// By convention, gcd(0,0) = 0 and gcd(0,a) = a
func gcd(a, b int64) (g, u, v int64) {

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

var ErrDivideByZero = fmt.Errorf("divide by 0")
var ErrNotDivisible = fmt.Errorf("is not divisible")

// Quo computes quotient q of a/b modulo limit, such that a = bq modulo limit, and stores result in c, returns error if not divisible.
// In general, this is very different from the usual integer quotient a/b.
func (c *CRI) Quo(a, b *CRI) error {

	bIsZero := true

	for i, bi := range b.rm {
		ai := a.rm[i]
		pi := a.e.primes[i]

		if ai == 0 {
			if bi != 0 {
				bIsZero = false
				c.rm[i] = 0
				continue
			} else {
				// ai=bi=0 ... ambiguous, could be anything !
				c.rm[i] = 1
			}
		} else {
			// ai != 0
			if bi == 0 {
				return ErrNotDivisible
			} else {
				bIsZero = false
				if bi == ai {
					c.rm[i] = 1
				} else {
					_, r, _ := gcd(bi, pi)
					c.rm[i] = (r * ai) % pi
				}
			}
		}
	}
	if bIsZero {
		return ErrDivideByZero
	}
	return nil
}

// computes a ^b moldulo m.
// b should be 0 or positive.
func expi(a, b, m int64) (r int64) {
	if b < 0 || (a == 0 && b == 0) || (m <= 1) {
		panic(fmt.Sprintf("operation not defined  : %v^%v[%v]", a, b, m))
	}
	a = a % m

	if a == 0 || a == 1 || b == 1 {
		return a
	}

	r = 1
	for b > 0 {
		if b%2 == 1 {
			// b is odd, multiply result
			r = (r * a) % m
		}
		b = b >> 1
		a = (a * a) % m
	}
	return r
}

/*
// Deprecated
func expb(a int64, n *big.Int, buffer *big.Int, m int64) (r int64) {
	if n.Sign() < 0 || (m <= 1) {
		panic(fmt.Sprintf("operation not defined yet : %v^%v[%v]", a, n, m))
	}
	a = a % m

	if n.Sign() == 0 {
		return 1
	}

	if a == 0 || a == 1 || (n.IsInt64() && n.Int64() == 1) {
		return a
	}

	r = 1
	buffer.Set(n)
	for buffer.Sign() > 0 {

		if buffer.Bit(0) == 1 {
			// b is odd, multiply result
			r = (r * a) % m
		}
		buffer.Rsh(buffer, 1)
		a = (a * a) % m
	}
	return r
}

*/

// ExpI computes a^n modulo limit, where the exponent is a positive int64, stores the result in c and returns it.
func (c *CRI) ExpI(a *CRI, n int64) *CRI {

	for i, ai := range a.rm {
		c.rm[i] = expi(ai, n, a.e.primes[i])
	}
	return c
}

/*
// Deprecated
func (c *CRI) ExpB(a *CRI, n *big.Int) *CRI {

	if n.Sign() == 0 {
		copy(c.rm, a.rm)
		return c
	}

	nn := big.NewInt(0)
	nn.Mod(n, a.Phi())

	buffer := big.NewInt(0)
	for i, ai := range a.rm {
		c.rm[i] = expb(ai, n, buffer, a.e.primes[i])
	}

	return c
}
*/

// Exp computes a^n modulo limit, where the exponent is a positive int64, stores the result in c and returns it.
func (c *CRI) Exp(a *CRI, n *big.Int) *CRI {
	switch n.Sign() {
	case 0:
		copy(c.rm, a.rm)
		return c
	case -1:
		panic("negative exponents are not implemented")
	default:

		nn := big.NewInt(0).Mod(n, a.Phi())
		aa := a.Clone()

		for i := range c.rm {
			c.rm[i] = 1
		}
		for nn.Sign() > 0 {

			if nn.Bit(0) == 1 {
				// n is odd, multiply result
				c.Mul(c, aa)
			}
			nn.Rsh(nn, 1)
			aa.Mul(aa, aa)
		}
		return c
		/*
			// this alternative, although simpler, is less efficient (x2) ...
			aa := a.ToBig()
			c.SetBig(aa.Exp(aa, n, a.Limit()))
			return c
		*/
	}

}
