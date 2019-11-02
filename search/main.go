package main

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"sync"
)

func main() {
	// build results
	var results [][2]*big.Int
	M := new(big.Int).Lsh(one, 256)
	var mutex = &sync.Mutex{}
	var waitGroup sync.WaitGroup
	waitGroup.Add(255)
	for j := 0; j < 255; j++ {
		go func(j int) {
			defer waitGroup.Done()
			Mj := new(big.Int).Lsh(one, uint(j))
			Mj.Sub(M, Mj)
			q1, t1 := nextSiec(Mj)
			q2, t2 := prevSiec(Mj)
			Mj = new(big.Int).Lsh(one, uint(j))
			Mj.Sub(M, Mj)
			q3, t3 := nextSiec(Mj)
			q4, t4 := prevSiec(Mj)
			mutex.Lock()
			results = append(
				results,
				[2]*big.Int{q1, t1},
				[2]*big.Int{q2, t2},
				[2]*big.Int{q3, t3},
				[2]*big.Int{q4, t4},
			)
			mutex.Unlock()
		}(j)
	}
	waitGroup.Wait()
	// filter results
	var newResults [][2]*big.Int
	for _, r := range results {
		N := new(big.Int).Add(r[0], big.NewInt(1))
		N.Sub(N, r[1])
		if N.ProbablyPrime(20) {
			newResults = append(newResults, r)
		}
	}
	results = newResults
	// rank results
	sort.Slice(results, func(i int, j int) bool {
		n1 := results[i][0].BitLen()
		c1 := popCount(results[i][0])
		n2 := results[j][0].BitLen()
		c2 := popCount(results[j][0])
		return math.Abs(float64(n1-2*int(c1))) < math.Abs(float64(n2-2*int(c2)))
	})
	// print
	for _, r := range results {
		fmt.Println("q = %d\nt = %d", r[0], r[1])
	}
}

func popCount(a *big.Int) (c uint) {
	n := a.BitLen()
	for i := 0; i <= n; i++ {
		c += a.Bit(i)
	}
	return
}
