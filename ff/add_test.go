package ff

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_add(t *testing.T) {
	p, _ := new(big.Int).SetString("28948022309329048855892746252183396360603931420023084536990047309120118726721", 10)
	i1 := big.NewInt(1)
	i2 := big.NewInt(2)
	e1 := FromBigInt(i1)
	e2 := FromBigInt(i2)
	pm1 := new(big.Int).Sub(p, i1)
	epm1 := FromBigInt(pm1)
	dpm1 := new(big.Int).Mod(new(big.Int).Add(pm1, pm1), p)
	edpm1 := FromBigInt(dpm1)
	type args struct {
		a Element
		b Element
	}
	tests := []struct {
		name  string
		args  args
		wantC Element
	}{
		{"1+1", args{e1, e1}, e2},
		{"p-1 + 2", args{epm1, e2}, e1},
		{"p-1 + p-1", args{epm1, epm1}, edpm1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotC := add(tt.args.a, tt.args.b); !reflect.DeepEqual(gotC, tt.wantC) {
				t.Errorf("add() = %v, want %v", gotC, tt.wantC)
			}
		})
	}
}

func TestAddRandom(t *testing.T) {
	g := newGenerator(1)
	p, _ := new(big.Int).SetString("28948022309329048855892746252183396360603931420023084536990047309120118726721", 10)
	for i := 0; i < 100; i++ {
		ia := g.next()
		a := FromBigInt(ia)
		ib := g.next()
		b := FromBigInt(ib)
		ic := new(big.Int).Mod(new(big.Int).Add(ia, ib), p)
		c := add(a, b)
		if ic.Cmp(ToBigInt(c)) != 0 {
			t.Errorf("/%d add(%v,%v) = %v, want %v", i, a, b, c, FromBigInt(ic))
			t.Errorf("/%d (alt) add(%v,%v) = %v, want %v", i, ia, ib, ToBigInt(c), ic)
			t.FailNow()
		}
	}
}

func BenchmarkFFAdd(b *testing.B) {
	g := newGenerator(1)
	a1 := FromBigInt(g.next())
	a2 := FromBigInt(g.next())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(a1, a2)
	}
}

func BenchmarkBIAdd1(b *testing.B) {
	g := newGenerator(1)
	a1 := g.next()
	a2 := g.next()
	c := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Add(a1, a2)
	}
}

func BenchmarkBIAdd2(b *testing.B) {
	g := newGenerator(1)
	a1 := g.next()
	a2 := g.next()
	c := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Add(a1, a2)
	}
}
