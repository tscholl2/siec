package ff

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_sub(t *testing.T) {
	type args struct {
		a Element
		b Element
	}
	tests := []struct {
		name string
		args args
		want Element
	}{
		{"1 - 1", args{Element{1, 0, 0, 0}, Element{1, 0, 0, 0}}, Element{0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sub(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sub() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Benchmark_sub(b *testing.B) {
	a1 := randomElement(1)
	a2 := randomElement(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sub(a1, a2)
	}
}

func Benchmark_subBI(b *testing.B) {
	a1 := ToBigInt(randomElement(1))
	a2 := ToBigInt(randomElement(2))
	c := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Sub(a1, a2)
	}
}
