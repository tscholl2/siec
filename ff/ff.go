package ff

import (
	"math/big"
)

var (
	// p as a *big.Int.
	pBI, _ = new(big.Int).SetString("28948022309329048855892746252183396360603931420023084536990047309120118726721", 10)
	// p as an Element.
	p = Element{1126179130581057, 9223372036854775808, 33558592, 4611686018427387904}
)

// Element represents a single element a ∈ 𝔽ₚ, where
//
//     p = 28948022309329048855892746252183396360603931420023084536990047309120118726721.
//
// Elements are represnted by an unsigned 256 bit integer.
// This is represented by an array of four 64 bit unsigned integers.
// Specifically,
//
//     element = [4]uint64{a,b,c,d} = (a + b*2^64 + c*2^128 + d*2^192)
//
// A representation is "normalized" if the it represents an integer in the interval [0,p).
type Element [4]uint64

// ElementToBigInt converts an element to a *big.Int.
func ElementToBigInt(a Element) (z *big.Int) {
	return uint64ArrayToBigInt(a[:])
}

// BigIntToElement converts a big.Int to an element.
func BigIntToElement(z *big.Int) (a Element) {
	arr := bigIntToUint64Array(z)
	for len(arr) < 4 {
		arr = append(arr, 0)
	}
	return Element{arr[0], arr[1], arr[2], arr[3]}
}

// Transform a *big.Int to []uint64 in little-endian.
func bigIntToUint64Array(z *big.Int) (arr []uint64) {
	z = new(big.Int).Set(z) // Use a copy to avoid overwriting anything.
	low64 := new(big.Int).SetUint64(0xffffffffffffffff)
	for z.BitLen() > 0 {
		arr = append(arr, new(big.Int).And(z, low64).Uint64())
		z.Rsh(z, 64)
	}
	return
}

func uint64ArrayToBigInt(arr []uint64) (z *big.Int) {
	if len(arr) == 0 {
		return new(big.Int)
	}
	z = new(big.Int).SetUint64(arr[len(arr)-1])
	for i := len(arr) - 2; i >= 0; i-- {
		z.Add(z.Lsh(z, 64), new(big.Int).SetUint64(arr[i]))
	}
	return z
}
