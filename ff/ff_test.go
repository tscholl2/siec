package ff

import (
	"math/big"
	"math/rand"
	"testing"
)

func TestToAndFromBigInt(t *testing.T) {
	i1 := big.NewInt(1)
	i64 := new(big.Int).Add(new(big.Int).Lsh(i1, 64), i1)
	i64p1 := new(big.Int).Add(i64, i1)
	i128p64p1 := new(big.Int).Add(new(big.Int).Add(new(big.Int).Lsh(i1, 128), i1), i64p1)
	tests := []struct {
		name string
		args *big.Int
	}{
		{"1", i1},
		{"1<<64", i64},
		{"1<<64+1", i64p1},
		{"1<<128 + 1<<64 + 1", i128p64p1},
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
	g := newGenerator(1)
	for i := 0; i < 10; i++ {
		i1 := g.next()
		e1 := FromBigInt(i1)
		i2 := ToBigInt(e1)
		e2 := FromBigInt(i2)
		if i1.Cmp(i2) != 0 {
			t.Errorf("ToBigInt(FromBigInt(%v)) = %v, want %v", i1, i2, i1)
			t.FailNow()
		}
		if e2 != e1 {
			t.Errorf("FromBigInt(ToBigInt(%v)) = %v, want %v", e1, e2, e1)
			t.FailNow()
		}
	}
}

func Test_isGreaterThanOrEqualToP(t *testing.T) {
	p, _ := new(big.Int).SetString("28948022309329048855892746252183396360603931420023084536990047309120118726721", 10)
	i1 := big.NewInt(1)
	pm1 := new(big.Int).Sub(p, i1)
	pp1 := new(big.Int).Add(p, i1)
	type args struct {
		a Element
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{FromBigInt(i1)}, false},
		{"p-1", args{FromBigInt(pm1)}, false},
		{"p", args{FromBigInt(p)}, true},
		{"p+1", args{FromBigInt(pp1)}, true},
		{"2*(p-1)", args{FromBigInt(new(big.Int).Add(pm1, pm1))}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isGreaterThanOrEqualToP(tt.args.a); got != tt.want {
				t.Errorf("isGreaterThanOrEqualToP() = %v, want %v", got, tt.want)
			}
		})
	}
}

type intGenerator struct {
	source rand.Source
}

func newGenerator(seed int64) intGenerator {
	return intGenerator{rand.NewSource(seed)}
}

func (b intGenerator) next() *big.Int {
	p, _ := new(big.Int).SetString("28948022309329048855892746252183396360603931420023084536990047309120118726721", 10)
	z := big.NewInt(b.source.Int63())
	z = z.Lsh(z, 64).Add(z, big.NewInt(b.source.Int63()))
	z = z.Lsh(z, 64).Add(z, big.NewInt(b.source.Int63()))
	z = z.Lsh(z, 64).Add(z, big.NewInt(b.source.Int63()))
	z = z.Mod(z, p)
	return z
}
