package ff

import (
	"math/big"
	"math/rand"
	"reflect"
	"testing"
)

func Test_mul(t *testing.T) {
	type args struct {
		a Element
		b Element
	}
	tests := []struct {
		name string
		args args
		want Element
	}{
		{"1 * 1", args{Element{1, 0, 0, 0}, Element{1, 0, 0, 0}}, Element{1, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mul(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_add256_random(t *testing.T) {
	for i := 0; i < 100; i++ {
		x := randArr(2 * i)
		y := randArr(2*i + 1)
		y[3] = y[3] & mask63 // clear last bit so garunteed it will fit in 256 bits
		got := add256(x, y)
		want := biArr(new(big.Int).Add(arrBI(x), arrBI(y)))
		if got != want {
			t.Errorf("/%d add256(%v,%v) = %v, want %v", i, x, y, got, want)
			t.FailNow()
		}
	}
}

func Benchmark_add256(b *testing.B) {
	x := randArr(1)
	y := randArr(2)
	y[3] = y[3] & mask63
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add256(x, y)
	}
}

func Benchmark_add256_BI(b *testing.B) {
	x := randArr(1)
	y := randArr(2)
	y[3] = y[3] & mask63
	A, B := arrBI(x), arrBI(y)
	C := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C.Add(A, B)
	}
}

func Test_sub256_random(t *testing.T) {
	for i := 0; i < 100; i++ {
		x := randArr(2 * i)
		y := randArr(2*i + 1)
		x[3] = x[3] | bit64  // set top bit so x > y
		y[3] = y[3] & mask63 // clear last bit so y < x
		got := sub256(x, y)
		want := biArr(new(big.Int).Sub(arrBI(x), arrBI(y)))
		if got != want {
			t.Errorf("/%d sub256(%v,%v) = %v, want %v", i, x, y, got, want)
			t.FailNow()
		}
	}
}

func Benchmark_sub256(b *testing.B) {
	x := randArr(1)
	y := randArr(2)
	x[3] = x[3] | bit64
	y[3] = y[3] & mask63
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sub256(x, y)
	}
}

func Benchmark_sub256_BI(b *testing.B) {
	x := randArr(1)
	y := randArr(2)
	x[3] = x[3] | bit64
	y[3] = y[3] & mask63
	A, B := arrBI(x), arrBI(y)
	C := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		C.Sub(A, B)
	}
}

func randArr(seed int) [4]uint64 {
	r := rand.NewSource(int64(seed))
	return [4]uint64{uint64(r.Int63()), uint64(r.Int63()), uint64(r.Int63()), uint64(r.Int63())}
}

func arrBI(x [4]uint64) *big.Int {
	a := new(big.Int).SetUint64(x[3])
	a.Lsh(a, 64)
	a.Add(a, new(big.Int).SetUint64(x[2]))
	a.Lsh(a, 64)
	a.Add(a, new(big.Int).SetUint64(x[1]))
	a.Lsh(a, 64)
	a.Add(a, new(big.Int).SetUint64(x[0]))
	return a
}

func biArr(n *big.Int) (x [4]uint64) {
	x[0] = n.Uint64()
	n = new(big.Int).Set(n)
	n.Rsh(n, 64)
	x[1] = n.Uint64()
	n.Rsh(n, 64)
	x[2] = n.Uint64()
	n.Rsh(n, 64)
	x[3] = n.Uint64()
	return
}

func Benchmark_mul(b *testing.B) {
	a1 := randomElement(1)
	a2 := randomElement(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mul(a1, a2)
	}
}
func Benchmark_mulBI(b *testing.B) {
	a1 := ToBigInt(randomElement(1))
	a2 := ToBigInt(randomElement(2))
	c := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Mod(c.Mul(a1, a2), pBI)
	}
}

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
