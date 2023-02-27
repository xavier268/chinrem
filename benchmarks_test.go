package chinrem

import (
	"fmt"
	"math/big"
	"testing"
	"time"
)

// package level variable to prevent compiler optimization.
var b1, b2, limit *big.Int
var e *CREngine
var c1, c2 *CRI

func reset() {

	b1, b2 = big.NewInt(14654), big.NewInt(654789)
	e = NewCREngine(100) // a very, very large "limit"
	c1, c2 = e.NewCRIBig(b1), e.NewCRIBig(b2)
	limit = e.Limit()
}

func BenchmarkBigVersusChinrem(b *testing.B) {

	fmt.Println(time.Now())

	b.Run("big.Int-mul", func(bb *testing.B) {
		reset()
		bb.ResetTimer()

		for i := 1; i < bb.N; i++ {
			b1.Mul(b2, b1)
		}
	})

	b.Run("chinrem.CRI-mul", func(bb *testing.B) {
		reset()
		bb.ResetTimer()

		for i := 1; i < bb.N; i++ {
			c1.Mul(c2, c1)
		}
	})

	b.Run("big.Int-inv", func(bb *testing.B) {
		reset()
		bb.ResetTimer()

		for i := 1; i < bb.N; i++ {
			b1.ModInverse(b1, limit)
		}
	})

	b.Run("chinrem.CRI-inv", func(bb *testing.B) {
		reset()
		bb.ResetTimer()

		for i := 1; i < bb.N; i++ {
			c1.Inv(c1)
		}
	})

}