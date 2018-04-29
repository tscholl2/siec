package ff

import (
	"math/big"
	"math/rand"
	"testing"
)

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

func randomElementNormalized(seed int64) Element {
	r := rand.NewSource(seed)
	return normalize(Element{uint64(r.Int63()), uint64(r.Int63()), uint64(r.Int63()), uint64(r.Int63())})
}
