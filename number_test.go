package chinrem

import (
	"fmt"
	"math/big"
	"math/rand"
	"testing"
)

func TestCRIVisual(t *testing.T) {

	e := NewCREngine(10)
	fmt.Println("engine", e)

	a := e.NewCRI()
	fmt.Println("a", a)

	b := e.NewCRIInt64(0)
	fmt.Println("b", b)

	if !b.Equal(a) {
		t.FailNow()
	}

	c := e.NewCRIInt64(300)
	fmt.Println("c", c)

	d := e.NewCRIBig(big.NewInt(300))
	fmt.Println("d", d)

	if !c.Equal(d) {
		t.Fail()
	}

	f := e.NewCRISlice([]int64{0, 0, 0, 6, 3, 1, 11, 15, 1, 10})
	fmt.Println("f", f, f.ToBig())
	if !f.Equal(d) {
		t.Fail()
	}
	if f.ToBig().Int64() != 300 {
		t.Fail()
	}

	g := f.Clone()
	if !f.Equal(g) {
		t.Fail()
	}
	g.Minus(g)
	if f.Equal(g) {
		t.Fail()
	}
	if g.ToBig().Int64() != -300+e.limit.Int64() {
		t.Fail()
	}
	g.Minus(g)
	if !f.Equal(g) {
		t.Fail()
	}

	h := e.NewCRISlice([]int64{300, 300, 300, 300, 300, 300, 300, 300, 300, 300})
	if h.Equal(f) {
		t.Fail()
	}
	h.Normalize()
	if !h.Equal(f) {
		t.Fail()
	}

	for i := 0; i < 10; i++ {
		r := e.NewCRIRand(rand.New(rand.NewSource(42 * int64(i))))
		fmt.Println("rand", r, r.ToBig())
		kk := r.Clone()
		kk.Normalize()
		if !kk.Equal(r) {
			t.FailNow()
		}
	}

}

func TestGcd(t *testing.T) {

	data := []int64{ // a,b,gcd(a,b)
		4, 5, 1,
		15, 3, 3,
		1, 1, 1,
		0, 5, 5,
		0, 0, 0,
		-2, 5, 1,
		-2, -5, 1,
		0, -3, 3,
		-15, 20, 5,
		30, -10, 10,
	}

	for i := 0; i < len(data); i += 3 {

		a, b, g := data[i], data[i+1], data[i+2]
		gg, u, v := gcd(a, b)
		if g != gg {
			t.Fatalf("gcd of %d and %d returned %d but expected %d", a, b, gg, g)
		}
		ggg := a*u + b*v
		if ggg != g {
			t.Fatalf("bezout equation invalid : %d*%d+%d*%d = %d (expected %d)", a, u, b, v, ggg, g)
		}

		// Switching a and b, and running same test
		a, b = b, a

		gg, u, v = gcd(a, b)
		if g != gg {
			t.Fatalf("gcd of %d and %d returned %d but expected %d", a, b, gg, g)
		}
		ggg = a*u + b*v
		if ggg != g {
			t.Fatalf("bezout equation invalid : %d*%d+%d*%d = %d (expected %d)", a, u, b, v, ggg, g)
		}
	}
}

func TestInv(t *testing.T) {

	e := NewCREngine(5)
	rd := rand.New(rand.NewSource(42))
	b := e.NewCRI()

	for i := 0; i < 100; i++ {
		a := e.NewCRIRand(rd)
		err := b.Inv(a)
		if err == nil {
			b.Mul(a, b)
			if !b.IsOne() {
				fmt.Println(b)
				t.Fatalf("\nfailed inversing %v modulo %v", a, b.e.limit)
			}
		} else {
			//fmt.Println(a, err)
		}
	}

}
