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

	z := NewCRI(5)
	fmt.Println(z)
	fmt.Println(z.Limit())
	if z.Limit().Int64() != 2310 {
		t.FailNow()
	}

	a := NewCRI(1000)
	fmt.Println(a.Limit())
	if a.Limit().Int64() == 2310 {
		t.FailNow()
	}

	b := a.Clone()
	if !a.Equal(b) {
		t.Fail()
	}
	b.Minus()
	if !a.Equal(b) {
		t.Fail()
	}
	b.Minus()
	if !a.Equal(b) {
		t.Fail()
	}

	c := NewCRIInt64(5, 1000)
	d := NewCRISlice([]int64{0, 1, 0, 6, 10})
	e := NewCRISlice([]int64{1000, 1000, 1000, 1000, 1000})
	f := c.Clone()

	if !c.Equal(f) {
		t.Fail()
	}
	f.Minus()
	if c.Equal(f) {
		t.Fail()
	}
	f.Minus()
	if !c.Equal(f) {
		t.Fail()
	}

	if d.Equal(e) {
		t.FailNow()
	}
	e.Normalize()
	if !d.Equal(e) {
		t.FailNow()
	}
	if !c.Equal(d) {
		t.FailNow()
	}

}
