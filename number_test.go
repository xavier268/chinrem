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
	if !h.Equal(f) { // now, the NewCRISlice normalize !
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
				t.Fatalf("\nfailed inverting %v modulo %v", a, b.e.limit)
			}
		}
	}
}

func TestEngineChange(t *testing.T) {

	e1 := NewCREngine(5)
	e2 := NewCREngine(10)
	rd := rand.New(rand.NewSource(42))

	fmt.Println("Extending ...")

	for i := 0; i < 10; i++ { // extend
		r1 := e1.NewCRIRand(rd)
		r2 := r1.CloneE(e2)
		fmt.Println(r1, r2)
		if r1.ToBig().Cmp(r2.ToBig()) != 0 {
			t.Fatalf("extending should not change the big value, but it did\n%v -> %v", r1.ToBig(), r2.ToBig())
		}
	}
	fmt.Println("Truncating ...")

	for i := 0; i < 10; i++ { // truncate (should changes the value)
		r1 := e2.NewCRIRand(rd)
		r2 := r1.CloneE(e1)
		//fmt.Println(r1, r2)
		if r1.ToBig().Cmp(r2.ToBig()) != 0 {
			fmt.Printf("value changed:\t%v\t-> %v\n", r1, r2)
		} else {
			fmt.Printf("value stay same:\t%v\t-> %v\n", r1, r2)
		}
	}
	for i := 0; i < 15; i++ { // truncate small values
		r1 := e2.NewCRIInt64(int64(i))
		r2 := r1.CloneE(e1)
		//fmt.Println(r1, r2)
		if r1.ToBig().Cmp(r2.ToBig()) != 0 {
			fmt.Printf("value changed:\t%v\t-> %v\n", r1, r2)
		} else {
			fmt.Printf("value stay same:\t%v\t-> %v\n", r1, r2)
		}
	}

}

func TestCmpVisual(t *testing.T) {
	e := NewCREngine(5)
	rd := rand.New(rand.NewSource(42))

	for i := 1; i < 20; i++ {
		a := e.NewCRIRand(rd)
		b := e.NewCRIRand(rd)
		ab := a.Cmp(b)
		ba := b.Cmp(a)
		switch ab {
		case +1:
			fmt.Println(a, ">", b)
		case 0:
			fmt.Println(a, "=", b)
		case -1:
			fmt.Println(a, "<", b)
		default:
			t.Fatalf("unexpected value for Cmp : %d", ab)
		}
		if ab != -ba {
			t.Fatalf("Comparison did not reverse correctly ?")
		}
	}

	for i := 1; i < 20; i++ {
		a := e.NewCRIInt64(int64(i - 1))
		b := e.NewCRIInt64(int64(i))
		ab := a.Cmp(b)
		ba := b.Cmp(a)
		switch ab {
		case +1:
			fmt.Println(a, ">", b)
		case 0:
			fmt.Println(a, "=", b)
		case -1:
			fmt.Println(a, "<", b)
		default:
			t.Fatalf("unexpected value for Cmp : %d", ab)
		}
		if ab != -ba {
			t.Fatalf("Comparison did not reverse correctly ?")
		}
	}

	for i := 1; i < 20; i++ { // compare to 0
		a := e.NewCRIInt64(0)
		b := e.NewCRIRand(rd)
		ab := a.Cmp(b)
		ba := b.Cmp(a)
		switch ab {
		case +1:
			fmt.Println(a, ">", b)
		case 0:
			fmt.Println(a, "=", b)
		case -1:
			fmt.Println(a, "<", b)
		default:
			t.Fatalf("unexpected value for Cmp : %d", ab)
		}
		if ab != -ba {
			t.Fatalf("Comparison did not reverse correctly ?")
		}
	}

	for i := 1; i < 20; i++ { // compare to limit-1
		a := e.NewCRIInt64(-1)
		b := e.NewCRIRand(rd)
		ab := a.Cmp(b)
		ba := b.Cmp(a)
		switch ab {
		case +1:
			fmt.Println(a, ">", b)
		case 0:
			fmt.Println(a, "=", b)
		case -1:
			fmt.Println(a, "<", b)
		default:
			t.Fatalf("unexpected value for Cmp : %d", ab)
		}
		if ab != -ba {
			t.Fatalf("Comparison did not reverse correctly ?")
		}
	}
}

func TestQuoVisual(t *testing.T) {
	e := NewCREngine(3)
	var a, b, q *CRI

	a = e.NewCRIInt64(4)
	b = e.NewCRIInt64(2)
	q = e.NewCRI()

	err := q.Quo(a, b)
	fmt.Println(err, a.ToBig(), "/", b.ToBig(), "=", q.ToBig())

}

func TestQuoSystematic(t *testing.T) {
	e := NewCREngine(7)
	var a, b, q *CRI
	m := e.Limit().Int64()
	q = e.NewCRI()

	for i := int64(0); i < 10; i++ {
		for j := int64(0); j < 10; j++ {
			a, b = e.NewCRIInt64(i), e.NewCRIInt64(j)
			err := q.Quo(a, b)
			qq := q.ToBig().Int64()
			if j != 0 {
				if err != nil {
					fmt.Printf("%d/%d=%d\t%d/%d=%v[%d]", i, j, i/j, i, j, err, m)
				} else {
					fmt.Printf("%d/%d=%d\t%d/%d=%d[%d]", i, j, i/j, i, j, qq, m)
				}
			} else {
				if err != nil {
					fmt.Printf("%d/%d=NaN\t%d/%d=%v[%d]", i, j, i, j, err, m)
				} else {
					fmt.Printf("%d/%d=NaN\t%d/%d=%d[%d]", i, j, i, j, qq, m)
				}
			}
			if err == nil {
				fmt.Printf("\t\t Verify %d * %d = %d [%d]\n", j, qq, (j*qq)%m, m)
				if i != (j*qq)%m {
					t.Fatal("Verification failed !", (j*qq)%m)
				}
			} else {
				fmt.Println()
			}

		}
	}
}
