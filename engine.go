package chinrem

import (
	"fmt"
	"math/big"
	"strings"
)

// CREngine is the configuation and processing object.
// It can be safely accessed concurrently, because it is never modified once created.
type CREngine struct {
	// The core parameters to compute efficiently modulo limit.
	size   int     // number of primes to consider
	primes []int64 // The prime base, int 64 format

	// The following is less efficient, but is only used for input/output and format conversion.
	limit    *big.Int   // product of all primes
	phi      *big.Int   // product of all (prime - 1)
	coprimes []*big.Int // coprimes [i] is the product of primes[j] for j != i, multiplied by its own inverse modulo prime[i], then modulo limit
}

func NewCREngine(size int) *CREngine {
	e := new(CREngine)
	e.size = size
	e.initPrimes()
	e.initLimit()
	e.initCoprimes()
	return e
}

// initPrimes compute the primes according to the size set in the engine.
// size should be >= 3, or it will be modified and set to 3.
// coprimes are untouched.
func (e *CREngine) initPrimes() {

	if e.size <= 3 {
		e.size = 3
	}

	e.primes = make([]int64, e.size)
	e.primes[0], e.primes[1], e.primes[2] = 2, 3, 5

	for i := 3; i < e.size; i++ {
		p := 2 + e.primes[i-1]
		pr := false
		for !pr {
			pr = true
			for j := 0; j < i; j++ {
				k := e.primes[j]
				if k == 0 {
					fmt.Println("DEBUG", k, i, j, p, e.primes)
				}
				if p%k == 0 {
					pr = false
					p += 2 // not prime, proceed
					break
				}
				if k*k >= p {
					pr = true // prime, proceed
					break
				}
			}
		}
		// now, p is prime
		e.primes[i] = p
	}
}

func (e *CREngine) initLimit() {
	e.limit = big.NewInt(1)
	e.phi = big.NewInt(1)
	for _, p := range e.primes {
		e.limit.Mul(e.limit, big.NewInt(p))
		e.phi.Mul(e.phi, big.NewInt(p-1))
	}
}

func (e *CREngine) initCoprimes() {

	e.coprimes = make([]*big.Int, e.size)
	bigprimes := make([]*big.Int, e.size)
	for i, p := range e.primes {
		bigprimes[i] = big.NewInt(p)
	}

	// coprimes [i] is the product of primes[j] for j != i, ...
	for i, pi := range bigprimes {

		cpi := big.NewInt(1)
		for j, pj := range bigprimes {
			if i != j {
				cpi.Mul(cpi, pj)
			}
		}
		// ... multiplied by its own inverse modulo prime[i]
		inv := new(big.Int)
		inv.ModInverse(cpi, pi)
		inv.Mul(inv, cpi)
		// ... modulo limit.
		e.coprimes[i] = inv.Mod(inv, e.limit)
	}
}

func (e *CREngine) String() string {
	sb := new(strings.Builder)
	fmt.Fprintf(sb, "\t\tSize\t%d\n", e.size)
	fmt.Fprintf(sb, "\t\tLimit\t%v\n", e.limit)
	fmt.Fprintf(sb, "\t\tPhi  \t%v\n", e.phi)
	fmt.Fprintln(sb, "\t\tPrimes\tCoprimes :")
	for i, p := range e.primes {
		fmt.Fprintf(sb, "%d\t%9d\t%v\n", i, p, e.coprimes[i])
	}
	return sb.String()
}

// Limit is the product of all the primes from the base.
func (e *CREngine) Limit() *big.Int {
	return e.limit
}

// Phi is the product of all the (prime - 1) in the base.
// For any non nul number a, we have a ^ phi = 1, modulo Limit.
func (e *CREngine) Phi() *big.Int {
	return e.phi
}
