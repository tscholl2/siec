package ff

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_neg(t *testing.T) {
	tests := []struct {
		name string
		args Element
		want Element
	}{
		{"1", Element{1, 0, 0, 0}, Element{1, 0, 0, bit64}},
		{"1", Element{1, 0, 0, bit64}, Element{1, 0, 0, 0}},
		{"0", Element{0, 0, 0, 0}, Element{0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := neg(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_negRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		a := randomElement(int64(i))
		A := ToBigInt(a)
		B := new(big.Int).Neg(A)
		if args, got, want := a, neg(a), FromBigInt(B); got != want {
			t.Errorf("/%d neg(%v) = %v, want %v", i, args, got, want)
			t.FailNow()
		}
	}
}

func Benchmark_neg(b *testing.B) {
	a := randomElement(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		neg(a)
	}
}
