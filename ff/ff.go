package ff

import (
	"math/big"
	"math/bits"
)

var (
	// p as a *big.Int.
	pBI, _ = new(big.Int).SetString("7fffffffffffffffffffffffffffddf40a09853f04c9246b1f1c11c8ad49dc91", 16)
	// p as an Element.
	p      = Element{2241686268122094737, 723255720879400043, 18446744073709542900, 576460752303423487}
	two256 = Element{13963371537465362142, 17000232631950751529, 17431, 0}
	two320 = Element{0, 13963371537465362142, 17000232631950751529, 1089}
	two384 = Element{6946274766781380961, 1834951742433355726, 13963371537769226683, 486053787193498482}
	two448 = Element{1195137535819132669, 7639405055976263065, 2336821022378087212, 296249968807153639}
	two512 = Element{15937601485410354464, 16159380161337193251, 11972309219966232048, 146051313898630730}
)

// Element represents a single element a ∈ 𝔽ₚ, where
//
//     p = 7fffffffffffffffffffffffffffddf40a09853f04c9246b1f1c11c8ad49dc91.
//
// Elements are represnted by an unsigned 256 bit integer.
// This is represented by an array of four 64 bit unsigned integers.
// Specifically,
//
//     element = [4]uint64{a,b,c,d} = (a + b*2^64 + c*2^128 + d*2^192)
//
// A representation is "normalized" if the it represents an integer in the interval [0,p).
type Element [4]uint64

// Add a = b + c
func (a Element) Add(b, c Element) {
	var carry uint64
	for i := 0; i < 4; i++ {
		a[i], carry = bits.Add64(b[i], c[i], carry)
	}
	// TODO: normalize
}

// Mul a = b*c
func (a Element) Mul(b, c Element) {
	// bits.Mul64(x, y uint64) (hi, lo uint64)
	/*
		(a + b*2^64 + c*2^128 + d*2^192) * (e + f*2^64 + g*2^128 + h*2^192)

	*/
}

// Inv a = 1/b
func (a Element) Inv(b Element) {

}

// Sub c = b - c
func (a Element) Sub(b, c Element) {
	var borrow uint64
	for i := 0; i < 4; i++ {
		a[i], borrow = bits.Sub64(b[i], c[i], borrow)
	}
	// TODO: check borrow and sub p if necessary?
	/*
	   func Sub64(x, y, borrow uint64) (diff, borrowOut uint64)
	   Sub64 returns the difference of x, y and borrow: diff = x - y - borrow.
	   The borrow input must be 0 or 1; otherwise the behavior is undefined.
	   The borrowOut output is guaranteed to be 0 or 1.
	*/
}

// Neg a = -b
func (a Element) Neg(b Element) {
	b.Sub(p, a)
}

// Cmp compares a to b as integers in [0,p).
func (a Element) Cmp(b Element) {

}

// ElementToBigInt converts an element to a *big.Int.
func ElementToBigInt(a Element) (z *big.Int) {
	arr := a[:]
	if len(arr) == 0 {
		return new(big.Int)
	}
	z = new(big.Int).SetUint64(arr[len(arr)-1])
	for i := len(arr) - 2; i >= 0; i-- {
		z.Add(z.Lsh(z, 64), new(big.Int).SetUint64(arr[i]))
	}
	return z
}

// BigIntToElement converts a big.Int to an element.
func BigIntToElement(z *big.Int) (a Element) {
	z = new(big.Int).Mod(z, pBI) // Use a copy to avoid overwriting anything.
	var arr []uint64
	low64 := new(big.Int).SetUint64(0xffffffffffffffff)
	for z.BitLen() > 0 {
		arr = append(arr, new(big.Int).And(z, low64).Uint64())
		z.Rsh(z, 64)
	}
	for len(arr) < 4 {
		arr = append(arr, 0)
	}
	return Element{arr[0], arr[1], arr[2], arr[3]}
}
