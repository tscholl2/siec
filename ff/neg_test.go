package ff

import (
	"math/big"
	"testing"
)

func TestNegRandom(t *testing.T) {
	g := newGenerator(1)
	p, _ := new(big.Int).SetString("28948022309329048855892746252183396360603931420023084536990047309120118726721", 10)
	for i := 0; i < 100; i++ {
		ia := g.next()
		a := FromBigInt(ia)
		ic := new(big.Int).Mod(new(big.Int).Neg(ia), p)
		c := neg(a)
		if ic.Cmp(ToBigInt(c)) != 0 {
			t.Errorf("/%d neg(%v) = %v, want %v", i, a, c, FromBigInt(ic))
			t.Errorf("/%d (alt) neg(%v) = %v, want %v", i, ia, ToBigInt(c), ic)
			t.FailNow()
		}
	}
}
