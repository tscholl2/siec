package main

import (
	"math/big"
)

var (
	_D  = []*big.Int{big.NewInt(3), big.NewInt(4), big.NewInt(7), big.NewInt(8), big.NewInt(11), big.NewInt(19), big.NewInt(43), big.NewInt(67), big.NewInt(163)}
	one = big.NewInt(1)
	two = big.NewInt(2)
)

func test(t *big.Int, q *big.Int) bool {
	for _, d := range _D {
		q.Exp(t, two, nil)
		q.Add(q, d)
		if q.Bits()[0]&3 != 0 {
			continue
		}
		q.Rsh(q, 2)
		// TODO: check if q is a prime power
		if q.ProbablyPrime(20) {
			return true
		}
	}
	return false
}

func nextSiec(M *big.Int) (q, t *big.Int) {
	t = new(big.Int)
	q = new(big.Int)
	t.Lsh(M, 2)
	t.Sub(t, _D[8])
	t.Sqrt(t)
	t.Sub(t, one)
	for {
		if test(t, q) && q.Cmp(M) > 0 {
			return
		}
		t.Add(t, one)
	}
}

func prevSiec(M *big.Int) (q, t *big.Int) {
	t = new(big.Int)
	q = new(big.Int)
	t.Lsh(M, 2)
	t.Sub(t, _D[0])
	t.Sqrt(t)
	t.Add(t, one)
	for {
		if test(t, q) && q.Cmp(M) < 0 {
			return
		}
		t.Sub(t, one)
	}
}
