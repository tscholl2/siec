package ff

import (
	"math/big"
	"math/rand"
	"reflect"
	"testing"
)

func TestFromBigInt(t *testing.T) {
	type args struct {
		n *big.Int
	}
	tests := []struct {
		name string
		args args
		want Element
	}{
		{"0", args{big.NewInt(0)}, Element{0, 0, 0, 0}},
		{"1", args{big.NewInt(1)}, Element{1, 0, 0, 0}},
		{"2^64", args{new(big.Int).Lsh(big.NewInt(1), 64)}, Element{0, 1, 0, 0}},
		{"p", args{pBI}, pFF},
		{"-1", args{big.NewInt(-1)}, Element{1, 0, 0, bit64}},
		{"-p", args{new(big.Int).Neg(pBI)}, Element{1126179130581057, 9223372036854775808, 33558592, 13835058055282163712}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotA := FromBigInt(tt.args.n); !reflect.DeepEqual(gotA, tt.want) {
				t.Errorf("FromBigInt() = %v, want %v", gotA, tt.want)
			}
		})
	}
}

func TestToBigInt(t *testing.T) {
	type args struct {
		a Element
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{"0", args{Element{0, 0, 0, 0}}, new(big.Int)},
		{"-p", args{Element{1126179130581057, 9223372036854775808, 33558592, 13835058055282163712}}, new(big.Int).Neg(pBI)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBigInt(tt.args.a); got.Cmp(tt.want) != 0 {
				t.Errorf("ToBigInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToAndFromBigInt(t *testing.T) {
	tests := []struct {
		name string
		args *big.Int
	}{
		{"1", bigIntFromString("1")},
		{"1<<64", bigIntFromString("18446744073709551616")},
		{"1<<64+1", bigIntFromString("18446744073709551617")},
		{"1<<128 + 1<<64 + 1", bigIntFromString("25108406941546723055343157692830665664409421777856138051584")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := tt.args
			e := FromBigInt(i)
			ii := ToBigInt(e)
			ee := FromBigInt(ii)
			if e != ee {
				t.Errorf("e != ee, got %v, want %v ", e, ee)
			}
			if i.Cmp(ii) != 0 {
				t.Errorf("i != ii, got %v, want %v", i, ii)
			}
		})
	}
}

func TestToAndFromBigIntRandom(t *testing.T) {
	for i := 0; i < 10; i++ {
		i1 := ToBigInt(randomElement(int64(i)))
		e1 := FromBigInt(i1)
		i2 := ToBigInt(e1)
		e2 := FromBigInt(i2)
		if i1.Cmp(i2) != 0 {
			t.Errorf("/%d ToBigInt(FromBigInt(%v)) = %v, want %v", i, i1, i2, i1)
			t.FailNow()
		}
		if e2 != e1 {
			t.Errorf("/%d FromBigInt(ToBigInt(%v)) = %v, want %v", i, e1, e2, e1)
			t.FailNow()
		}
	}
}

func bigIntFromString(s string) *big.Int {
	n, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("invalid number: " + s)
	}
	return n
}

func randomElement(seed int64) Element {
	r := rand.NewSource(seed)
	return Element{uint64(r.Int63()), uint64(r.Int63()), uint64(r.Int63()), uint64(r.Int63())}
}
