package ff

import (
	"math/big"
	"testing"
)

func TestDoubleRandom(t *testing.T) {
	g := newGenerator(1)
	two := big.NewInt(2)
	p, _ := new(big.Int).SetString("28948022309329048855892746252183396360603931420023084536990047309120118726721", 10)
	for i := 0; i < 100; i++ {
		ia := g.next()
		a := FromBigInt(ia)
		ic := new(big.Int).Mod(new(big.Int).Mul(ia, two), p)
		c := double(a)
		if ic.Cmp(ToBigInt(c)) != 0 {
			t.Errorf("/%d double(%v) = %v, want %v", i, a, c, FromBigInt(ic))
			t.Errorf("/%d (alt) double(%v) = %v, want %v", i, ia, ToBigInt(c), ic)
			t.FailNow()
		}
	}
}

func BenchmarkFFDouble(b *testing.B) {
	g := newGenerator(1)
	a := FromBigInt(g.next())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		double(a)
	}
}

func BenchmarkFFDouble2(b *testing.B) {
	g := newGenerator(1)
	a := FromBigInt(g.next())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(a, a)
	}
}
