package ff

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_mul(t *testing.T) {
	type args struct {
		x Element
		y Element
	}
	tests := []struct {
		name  string
		args  args
		wantZ Element
	}{
		{"1*1", args{Element{1, 0, 0, 0}, Element{1, 0, 0, 0}}, Element{1, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotZ := mul(tt.args.x, tt.args.y); !reflect.DeepEqual(gotZ, tt.wantZ) {
				t.Errorf("mul() = %v, want %v", gotZ, tt.wantZ)
			}
		})
	}
}

func Benchmark_mul(b *testing.B) {
	x := randomElement(1)
	y := randomElement(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mul(x, y)
	}
}

func Benchmark_mul_BI(b *testing.B) {
	x := randomElement(1)
	y := randomElement(2)
	A, B := ElementToBigInt(x), ElementToBigInt(y)
	C := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C.Mod(C.Mul(A, B), pBI)
	}
}
