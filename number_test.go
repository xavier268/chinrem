package chinrem

import (
	"fmt"
	"testing"
)

func TestPrimesVisual(t *testing.T) {

	primes = []int64{2, 3}
	n := 5

	extendPrime(n)
	//fmt.Printf("%#v\n", primes)
	t.Log(primes)
	if len(primes) != n {
		t.FailNow()
	}

	n = 100
	extendPrime(n)
	//fmt.Printf("%#v\n", primes)
	t.Log(primes)
	if len(primes) != n {
		t.FailNow()
	}
	// No op
	extendPrime(n - 20)
	//fmt.Printf("%#v\n", primes)
	t.Log(primes)
	if len(primes) != n {
		t.FailNow()
	}

	if primes[99] != 541 {
		t.FailNow()
	}

	n = 10000
	extendPrime(n)
	//fmt.Printf("%#v\n", primes)
	if len(primes) != n {
		t.Log(primes)
		t.FailNow()
	}
}

func TestCRIVisual(t *testing.T) {

	c := NewCRI(5)
	fmt.Println(c)
	fmt.Println(c.Max())
	if c.Max().Int64() != 2310 {
		t.Fail()
	}

}
