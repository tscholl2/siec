package ff

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_add(t *testing.T) {
	type args struct {
		a Element
		b Element
	}
	tests := []struct {
		name string
		args args
		want Element
	}{
		{"1 + 1", args{Element{1, 0, 0, 0}, Element{1, 0, 0, 0}}, Element{2, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		a1 := randomElementNormalized(int64(2 * i))
		a2 := randomElementNormalized(int64(2*i + 1))
		A1 := ToBigInt(a1)
		A2 := ToBigInt(a2)
		B := new(big.Int).Add(A1, A2)
		if args, got, want := []Element{a1, a2}, add(a1, a2), FromBigInt(B); got != want {
			t.Errorf("/%d add(%v) = %v, want %v", i, args, got, want)
			t.FailNow()
		}
	}
}

func Benchmark_add(b *testing.B) {
	a1 := randomElement(1)
	a2 := randomElement(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(a1, a2)
	}
}

func Benchmark_addBI(b *testing.B) {
	a1 := ToBigInt(randomElement(1))
	a2 := ToBigInt(randomElement(2))
	c := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Add(a1, a2)
	}
}
