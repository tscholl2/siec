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
	M := new(big.Int).Sub(new(big.Int).Lsh(one, 256), big.NewInt(1))
	var mutex = &sync.Mutex{}
	var waitGroup sync.WaitGroup
	waitGroup.Add(255)
	for j := -1; j < 10; j++ {
		go func(j int) {
			defer waitGroup.Done()
			q := new(big.Int).Set(M)
			if j >= 0 {
				c := big.NewInt(1)
				c.Lsh(c, uint(j))
				q.Sub(q, c)
			}
			for {
				var t *big.Int
				q, t = prevSiec(q)
				N := new(big.Int).Add(q, big.NewInt(1))
				N.Sub(N, t)
				if N.ProbablyPrime(25) {
					mutex.Lock()
					results = append(results, [2]*big.Int{q, t})
					mutex.Unlock()
					return
				}
			}
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
		fmt.Printf("q = %d\nt = %d\n", r[0], r[1])
	}
}

func popCount(a *big.Int) (c uint) {
	n := a.BitLen()
	for i := 0; i <= n; i++ {
		c += a.Bit(i)
	}
	return
}
