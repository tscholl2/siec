package ff

import (
	"math/big"
	"math/rand"
	"reflect"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	var a Element
	for i := 0; i < b.N; i++ {
		Add(a, a)
	}
}

func BenchmarkMul(b *testing.B) {
	var a Element
	for i := 0; i < b.N; i++ {
		Mul(a, a)
	}
}

func BenchmarkAddBigInt(b *testing.B) {
	z := randomBigInt(1)
	w := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Add(z, z)
		w.Mod(w, pBI)
	}
}

func BenchmarkMulBigInt(b *testing.B) {
	z := randomBigInt(1)
	y := randomBigInt(2)
	w := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w.Mul(z, y)
		w.Mod(w, pBI)
	}
}

func TestElementToBigInt(t *testing.T) {
	type args struct {
		a Element
	}
	tests := []struct {
		name  string
		args  args
		wantZ *big.Int
	}{
		{"0", args{Element{0, 0, 0, 0}}, big.NewInt(0)},
		{"1", args{Element{1, 0, 0, 0}}, big.NewInt(1)},
		{"2^64", args{Element{0, 1, 0, 0}}, new(big.Int).Lsh(big.NewInt(1), 64)},
		{"p", args{p}, pBI},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotZ := ElementToBigInt(tt.args.a); !reflect.DeepEqual(gotZ, tt.wantZ) {
				t.Errorf("ElementToBigInt() = %v, want %v", gotZ, tt.wantZ)
			}
		})
	}
}

func TestBigIntToElement(t *testing.T) {
	type args struct {
		z *big.Int
	}
	tests := []struct {
		name  string
		args  args
		wantA Element
	}{
		{"0", args{big.NewInt(0)}, Element{0, 0, 0, 0}},
		{"1", args{big.NewInt(1)}, Element{1, 0, 0, 0}},
		{"2^64", args{new(big.Int).Lsh(big.NewInt(1), 64)}, Element{0, 1, 0, 0}},
		{"p", args{pBI}, Element{0, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotA := BigIntToElement(tt.args.z); !reflect.DeepEqual(gotA, tt.wantA) {
				t.Errorf("BigIntToElement() = %v, want %v", gotA, tt.wantA)
			}
		})
	}
}

func Test_BigIntToElement_ElementToBigInt_random(t *testing.T) {
	for i := 0; i < 10; i++ {
		a := randomElement(i)
		n := ElementToBigInt(a)
		aa := BigIntToElement(n)
		if a != aa {
			t.Errorf("/%d BigIntToElement(ElementToBigInt(%v)) = %v, want %v", i, a, aa, a)
			t.FailNow()
		}
		nn := ElementToBigInt(aa)
		if n.Cmp(nn) != 0 {
			t.Errorf("/%d ElementToBigInt(BigIntToElement(%v)) = %v, want %v", i, n, nn, n)
			t.FailNow()
		}
	}
}

func randomElement(seed int) Element {
	r := rand.NewSource(int64(seed))
	return Element{uint64(r.Int63()), uint64(r.Int63()), uint64(r.Int63()), uint64(r.Int63())}
}

func randomBigInt(seed int) *big.Int {
	r := rand.New(rand.NewSource(int64(seed)))
	z := new(big.Int)
	n := new(big.Int).Lsh(big.NewInt(1), 256)
	z.Rand(r, n)
	return z
}
