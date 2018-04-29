package ff

import (
	"math/big"
	"math/rand"
	"testing"
)

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

func TestNormalize(t *testing.T) {
	tests := []struct {
		name string
		args *big.Int
		want *big.Int
	}{
		{"1", bigIntFromString("1"), bigIntFromString("1")},
		{"p-1", bigIntFromString("28948022309329048855892746252183396360603931420023084536990047309120118726720"), bigIntFromString("14474011154664524427946373126091698180301965710011542268495023654560059363360")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := FromBigInt(tt.args)
			want := FromBigInt(tt.want)
			got := normalize(args)
			if got != want {
				t.Errorf("normalize(%v) = %v, want %v ", args, got, want)
			}
		})
	}
}

func TestNormalizeRandom(t *testing.T) {
	g := newGenerator(1)
	for i := 0; i < 10; i++ {
		n := g.next()
		ai := n.Mod(g.next(), pAsBigInt)
		a := FromBigInt(ai)
		b := normalize(a)
		if FromBigInt(n) != b {
			t.Errorf("/%d normalize(%v) = %v, want %v", i, a, b, FromBigInt(n))
			t.FailNow()
		}
	}
}

func Test_isGreaterThanOrEqualToP(t *testing.T) {
	i1 := big.NewInt(1)
	pm1 := new(big.Int).Sub(pAsBigInt, i1)
	pp1 := new(big.Int).Add(pAsBigInt, i1)
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
		{"p", args{FromBigInt(pAsBigInt)}, true},
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
	z := big.NewInt(b.source.Int63())
	z = z.Lsh(z, 64).Add(z, big.NewInt(b.source.Int63()))
	z = z.Lsh(z, 64).Add(z, big.NewInt(b.source.Int63()))
	z = z.Lsh(z, 64).Add(z, big.NewInt(b.source.Int63()))
	z = z.Mod(z, pAsBigInt)
	if z.Cmp(pMinusOneOver2AsBigInt) > 0 {
		z.Sub(z, pAsBigInt)
	}
	return z
}

func bigIntFromString(s string) *big.Int {
	n, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("invalid number: " + s)
	}
	return n
}
