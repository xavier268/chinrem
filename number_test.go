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
	g.Minus()
	if f.Equal(g) {
		t.Fail()
	}
	if g.ToBig().Int64() != -300+e.limit.Int64() {
		t.Fail()
	}
	g.Minus()
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
