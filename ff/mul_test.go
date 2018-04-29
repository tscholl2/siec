package ff

import (
	"math/big"
	"testing"
)

func sTestMulRandom(t *testing.T) {
	g := newGenerator(1)
	p, _ := new(big.Int).SetString("28948022309329048855892746252183396360603931420023084536990047309120118726721", 10)
	for i := 0; i < 100; i++ {
		ia := g.next()
		a := FromBigInt(ia)
		ib := g.next()
		b := FromBigInt(ib)
		ic := new(big.Int).Mod(new(big.Int).Mul(ia, ib), p)
		c := mul(a, b)
		if ic.Cmp(ToBigInt(c)) != 0 {
			t.Errorf("/%d mul(%v,%v) = %v, want %v", i, a, b, c, FromBigInt(ic))
			t.Errorf("/%d (alt) mul(%v,%v) = %v, want %v", i, ia, ib, ToBigInt(c), ic)
			t.FailNow()
		}
	}
}

func BenchmarkBIMul(b *testing.B) {
	p, _ := new(big.Int).SetString("28948022309329048855892746252183396360603931420023084536990047309120118726721", 10)
	g := newGenerator(1)
	a1 := g.next()
	a2 := g.next()
	c := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Mod(c.Mul(a1, a2), p)
	}
}

func BenchmarkFFMul(b *testing.B) {
	g := newGenerator(1)
	a1 := FromBigInt(g.next())
	a2 := FromBigInt(g.next())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mul(a1, a2)
	}
}
