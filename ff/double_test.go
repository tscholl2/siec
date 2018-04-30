package ff

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_double(t *testing.T) {
	tests := []struct {
		name string
		args Element
		want Element
	}{
		{"1", Element{1, 0, 0, 0}, Element{2, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := double(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("double() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doubleRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := normalize(randomElement(int64(2 * i)))
		B := new(big.Int).Add(ToBigInt(a), ToBigInt(a))
		if args, got, want := []Element{a}, double(a), FromBigInt(B); got != want {
			t.Errorf("/%d double(%v) = %v, want %v", i, args, got, want)
			t.FailNow()
		}
	}
}

func Benchmark_double(b *testing.B) {
	a := randomElement(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		double(a)
	}
}
func Benchmark_double2(b *testing.B) {
	a := randomElement(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(a, a)
	}
}
