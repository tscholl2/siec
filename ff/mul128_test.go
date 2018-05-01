package ff

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_mul128(t *testing.T) {
	type args struct {
		x [2]uint64
		y [2]uint64
	}
	tests := []struct {
		name  string
		args  args
		wantZ [4]uint64
	}{
		{"2*1", args{[2]uint64{2, 0}, [2]uint64{1, 0}}, [4]uint64{2, 0, 0, 0}},
		{"(2^128 - 1)(2^128 - 1)", args{[2]uint64{0xffffffffffffffff, 0xffffffffffffffff}, [2]uint64{0xffffffffffffffff, 0xffffffffffffffff}}, [4]uint64{0x1, 0, 0xfffffffffffffffe, 0xffffffffffffffff}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotZ := mul128(tt.args.x, tt.args.y); !reflect.DeepEqual(gotZ, tt.wantZ) {
				t.Errorf("mul128() = %v, want %v", gotZ, tt.wantZ)
			}
		})
	}
}

func Test_mul128_random(t *testing.T) {
	for i := 0; i < 100; i++ {
		x := [2]uint64{randArr(2 * i)[0], randArr(2 * i)[1]}
		y := [2]uint64{randArr(2*i + 1)[0], randArr(2*i + 1)[1]}
		got := mul128(x, y)
		want := biArr(new(big.Int).Mul(arrBI([4]uint64{x[0], x[1], 0, 0}), arrBI([4]uint64{y[0], y[1], 0, 0})))
		if got != want {
			t.Errorf("/%d mul128(%v,%v) = %v, want %v", i, x, y, got, want)
			t.FailNow()
		}
	}
}

func Benchmark_mul128(b *testing.B) {
	x := [2]uint64{randArr(1)[0], randArr(1)[1]}
	y := [2]uint64{randArr(2)[0], randArr(2)[1]}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mul128(x, y)
	}
}

func Benchmark_mul128_2(b *testing.B) {
	x := [2]uint64{randArr(1)[0], randArr(1)[1]}
	y := [2]uint64{randArr(2)[0], randArr(2)[1]}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mul128_2(x, y)
	}
}

func Benchmark_mul128_BI(b *testing.B) {
	x := [4]uint64{randArr(1)[0], randArr(1)[1], 0, 0}
	y := [4]uint64{randArr(2)[0], randArr(2)[1], 0, 0}
	A, B := arrBI(x), arrBI(y)
	C := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C.Mul(A, B)
	}
}
