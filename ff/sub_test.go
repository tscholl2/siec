package ff

import (
	"math/big"
	"testing"
)

func Test_sub_random(t *testing.T) {
	for i := 0; i < 100; i++ {
		x := randomElement(2 * i)
		y := randomElement(2*i + 1)
		x[3] = x[3] | 0x8000000000000000 // set top bit so x > y
		y[3] = y[3] & 0x7fffffffffffffff // clear last bit so y < x
		got := sub(x, y)
		want := BigIntToElement(new(big.Int).Sub(ElementToBigInt(x), ElementToBigInt(y)))
		if got != want {
			t.Errorf("/%d sub(%v,%v) = %v, want %v", i, x, y, got, want)
			t.FailNow()
		}
	}
}

func Benchmark_sub(b *testing.B) {
	x := randomElement(1)
	y := randomElement(2)
	y[3] = y[3] & 0x7fffffffffffffff
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sub(x, y)
	}
}

func Benchmark_sub_BI(b *testing.B) {
	x := randomElement(1)
	y := randomElement(2)
	y[3] = y[3] & 0x7fffffffffffffff
	A, B := ElementToBigInt(x), ElementToBigInt(y)
	C := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C.Sub(A, B)
	}
}
