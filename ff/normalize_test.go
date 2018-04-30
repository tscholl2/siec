package ff

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_normalize(t *testing.T) {
	type args struct {
		a Element
	}
	tests := []struct {
		name string
		args Element
		want Element
	}{
		{"0", Element{0, 0, 0, 0}, Element{0, 0, 0, 0}},
		{"1", Element{1, 0, 0, 0}, Element{1, 0, 0, 0}},
		{"-1", Element{1, 0, 0, bit64}, Element{1, 0, 0, bit64}},
		{"p -> 0", pFF, Element{0, 0, 0, 0}},
		{"(p-1)/2 -> (p-1)/2", Element{563089565290528, 4611686018427387904, 16779296, 2305843009213693952}, Element{563089565290528, 4611686018427387904, 16779296, 2305843009213693952}},
		{"(p+1)/2 -> -(p-1)/2", Element{563089565290529, 4611686018427387904, 16779296, 2305843009213693952}, Element{563089565290528, 4611686018427387904, 16779296, 11529215046068469760}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalize(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_normalizeRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		a := randomElement(int64(i))
		A := ToBigInt(a)
		B := new(big.Int).Mod(A, pBI)
		if B.Cmp(pMinusOneOver2BI) > 0 {
			B.Sub(B, pBI)
		}
		if args, got, want := a, normalize(a), FromBigInt(B); got != want {
			t.Errorf("/%d normalize(%v) = %v, want %v", i, args, got, want)
			t.FailNow()
		}
	}
}

func Benchmark_normalize(b *testing.B) {
	a := randomElement(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		normalize(a)
	}
}
