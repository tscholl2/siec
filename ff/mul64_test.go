package ff

import (
	"math/big"
	"math/rand"
	"reflect"
	"testing"
)

func Test_mul64(t *testing.T) {
	type args struct {
		x uint64
		y uint64
	}
	tests := []struct {
		name  string
		args  args
		wantZ [2]uint64
	}{
		{"2*1", args{2, 1}, [2]uint64{2, 0}},
		{"(2^64 - 1)(2^64 - 1)", args{0xffffffffffffffff, 0xffffffffffffffff}, [2]uint64{0x0000000000000001, 0xfffffffffffffffe}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotZ := mul64(tt.args.x, tt.args.y); !reflect.DeepEqual(gotZ, tt.wantZ) {
				t.Errorf("mul64() = %v, want %v", gotZ, tt.wantZ)
			}
		})
	}
}

func Test_mul64_random(t *testing.T) {
	r := rand.NewSource(1)
	for i := 0; i < 100; i++ {
		a := uint64(r.Int63())
		b := uint64(r.Int63())
		A := new(big.Int).SetUint64(a)
		B := new(big.Int).SetUint64(b)
		got := mul64(a, b)
		C1 := new(big.Int).Mul(A, B)
		want := [2]uint64{C1.Uint64(), C1.Rsh(C1, 64).Uint64()}
		if got != want {
			t.Errorf("/%d mul64(%v,%v) = %v, want %v", i, a, b, got, want)
			t.FailNow()
		}
	}
}

func Benchmark_mul64(b *testing.B) {
	r := rand.NewSource(1)
	x := uint64(r.Int63())
	y := uint64(r.Int63())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mul64(x, y)
	}
}

func Benchmark_mul64_BI(b *testing.B) {
	r := rand.NewSource(1)
	x := uint64(r.Int63())
	y := uint64(r.Int63())
	A := new(big.Int).SetUint64(x)
	B := new(big.Int).SetUint64(y)
	C := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C.Mul(A, B)
	}
}
