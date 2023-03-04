package chinrem

import (
	"fmt"
	"math/big"
	"testing"
	"time"
)

// package level variable to prevent compiler optimization.
var b1, b2, b3, limit *big.Int
var e *CREngine
var c1, c2, c3 *CRI

func reset() {

	b1, b2, b3 = big.NewInt(14654), big.NewInt(654789), big.NewInt(0)

	e = NewCREngine(100) // a very, very large "limit"
	c1, c2, c3 = e.NewCRIBig(b1), e.NewCRIBig(b2), e.NewCRIBig(b3)
	limit = e.Limit()
}

func BenchmarkBvC(b *testing.B) {

	fmt.Println(time.Now())

	b.Run("big.mul", func(bb *testing.B) {
		reset()
		bb.ResetTimer()

		for i := 1; i < bb.N; i++ {
			b1.Mul(b2, b1)
		}
	})

	b.Run("chinrem.mul", func(bb *testing.B) {
		reset()
		bb.ResetTimer()

		for i := 1; i < bb.N; i++ {
			c1.Mul(c2, c1)
		}
	})

	b.Run("big.inv", func(bb *testing.B) {
		reset()
		bb.ResetTimer()

		for i := 1; i < bb.N; i++ {
			b1.ModInverse(b1, limit)
		}
	})

	b.Run("chinrem.inv", func(bb *testing.B) {
		reset()
		bb.ResetTimer()

		for i := 1; i < bb.N; i++ {
			c1.Inv(c1)
		}
	})

	b.Run("big.Exp", func(bb *testing.B) {
		reset()
		bb.ResetTimer()
		for i := 1; i < bb.N; i++ {
			b3.Exp(b1, b2, limit)
		}
	})

	b.Run("chinrem.Exp", func(bb *testing.B) {
		reset()
		bb.ResetTimer()

		for i := 1; i < bb.N; i++ {
			c3.Exp(c1, b2)
		}
	})

}
