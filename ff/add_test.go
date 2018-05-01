package ff

import (
	"math/big"
	"testing"
)

func Test_add_random(t *testing.T) {
	for i := 0; i < 100; i++ {
		x := randomElement(2 * i)
		y := randomElement(2*i + 1)
		y[3] = y[3] & 0x7fffffffffffffff // clear last bit so garunteed it will fit in 256 bits
		got := add(x, y)
		want := BigIntToElement(new(big.Int).Add(ElementToBigInt(x), ElementToBigInt(y)))
		if got != want {
			t.Errorf("/%d add(%v,%v) = %v, want %v", i, x, y, got, want)
			t.FailNow()
		}
	}
}

func Benchmark_add(b *testing.B) {
	x := randomElement(1)
	y := randomElement(2)
	y[3] = y[3] & 0x7fffffffffffffff
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(x, y)
	}
}

func Benchmark_add_BI(b *testing.B) {
	x := randomElement(1)
	y := randomElement(2)
	y[3] = y[3] & 0x7fffffffffffffff
	A, B := ElementToBigInt(x), ElementToBigInt(y)
	C := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C.Add(A, B)
	}
}
