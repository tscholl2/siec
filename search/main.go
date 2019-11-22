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
	waitGroup.Add(1)
	for j := 0; j < 1; j++ {
		go func(j int) {
			defer waitGroup.Done()
			M2 := new(big.Int).Set(M)
			c := new(big.Int).Lsh(one, uint(130))
			M2.Sub(M2, c)
			for {
				q, t := prevSiec(M2)
				N := new(big.Int).Add(q, big.NewInt(1))
				N.Sub(N, t)
				N2 := new(big.Int).Add(q, big.NewInt(1))
				N2.Add(N2, t)
				if true || N.ProbablyPrime(25) || N2.ProbablyPrime(25) {
					mutex.Lock()
					results = append(results, [2]*big.Int{q, t})
					mutex.Unlock()
					return
				}
			}
		}(j)
	}
	waitGroup.Wait()
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
		fmt.Printf("q = %x\nt = %x\n", r[0], r[1])
	}
}

func popCount(a *big.Int) (c uint) {
	n := a.BitLen()
	for i := 0; i <= n; i++ {
		c += a.Bit(i)
	}
	return
}
