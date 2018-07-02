package ff

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_mul256(t *testing.T) {
	type args struct {
		x [4]uint64
		y [4]uint64
	}
	tests := []struct {
		name  string
		args  args
		wantZ [8]uint64
	}{
		{"2*1", args{[4]uint64{2, 0, 0, 0}, [4]uint64{1, 0, 0, 0}}, [8]uint64{2, 0, 0, 0, 0, 0, 0, 0}},
		{
			"(2^128 - 1)(2^128 - 1)",
			args{
				[4]uint64{0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff},
				[4]uint64{0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff},
			},
			[8]uint64{1, 0, 0, 0, 0xfffffffffffffffe, 0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotZ := mul256(tt.args.x, tt.args.y); !reflect.DeepEqual(gotZ, tt.wantZ) {
				t.Errorf("mul256() = %v, want %v", gotZ, tt.wantZ)
			}
		})
	}
}

func Test_mul256_random(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := randomElement(i)
		A := ElementToBigInt(a)
		b := randomElement(i + 1)
		B := ElementToBigInt(b)
		c := mul256(a, b)
		C := new(big.Int).Mul(A, B)
		if C.Cmp(uint64ArrayToBigInt(c[:])) != 0 {
			t.Errorf("mul256() = %v, want %v", c, bigIntToUint64Array(C))
		}
	}
}

func Benchmark_mul256(b *testing.B) {
	x := randomElement(1)
	y := randomElement(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mul256(x, y)
	}
}

func Benchmark_mul256_BI(b *testing.B) {
	x := randomElement(1)
	y := randomElement(2)
	A, B := ElementToBigInt(x), ElementToBigInt(y)
	C := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C.Mul(A, B)
	}
}
