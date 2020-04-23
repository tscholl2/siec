package ff

import (
	"math/big"
	"math/bits"
)

var (
	// p as a *big.Int.
	pBI, _ = new(big.Int).SetString("7fffffffffffffffffffffffffffddf40a09853f04c9246b1f1c11c8ad49dc91", 16)
	// p as an Element.
	p = Element{2241686268122094737, 723255720879400043, 18446744073709542900, 9223372036854775807}
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

// Add a + b
func Add(a, b Element) Element {
	a = removeTopBit(a)
	b = removeTopBit(b)
	return add(a, b)
}

func add(a [4]uint64, b [4]uint64) (c [4]uint64) {
	var carry uint64
	for i := 0; i < 4; i++ {
		c[i], carry = bits.Add64(a[i], b[i], carry)
	}
	return
}

func removeTopBit(a [4]uint64) [4]uint64 {
	for i := 0; i < 2; i++ {
		topBit := a[3] >> 63
		a[3] &= 0x7fffffffffffffff
		// 2^255 = [4]uint64{16205057805587456879, 17723488352830151572, 8715, 0}
		var two255 = [4]uint64{16205057805587456879 * topBit, 17723488352830151572 * topBit, 8715 * topBit, 0}
		a = add(a, two255)
	}
	return a
}

func mul64by256(x uint64, y [4]uint64) (a [5]uint64) {
	var carry, lo uint64
	for i := 0; i < 3; i++ {
		a[i+1], lo = bits.Mul64(x, y[i])
		a[i], carry = bits.Add64(y[i], lo, carry)
	}
	return a
}

func reduce5(a [5]uint64) (b [4]uint64) {
	b0 := removeTopBit([4]uint64{a[0], a[1], a[2], a[3]})
	// 2^256 = [4]uint64{13963371537465362142, 17000232631950751529, 17431, 0}
	b1 := mul64by256(a[4], [4]uint64{13963371537465362142, 17000232631950751529, 17431, 0})
	// don't worry about removing the top bit of b1
	// because (2^256 % p) * (2^64 - 1) < 2^208 << 2^255
	b = add(b0, [4]uint64{b1[0], b1[1], b1[2], b1[3]})
	return
}

func mul(a, b [4]uint64) [8]uint64 {
	/*
		(a + b*2^64 + c*2^128 + d*2^192) * (e + f*2^64 + g*2^128 + h*2^192)

		ae,af,ag,ah = m00*2^000 + m01*2^064 + m02*2^128 + m03*2^192 + m04*2^256
		be,bf,bg,bh =             m10*2^064 + m11*2^128 + m12*2^192 + m13*2^256 + m14*2^320
		ce,cf,cg,ch =                         m20*2^128 + m21*2^192 + m22*2^256 + m23*2^320 + m24*2^384
		de,df,dg,dh =                                     m30*2^192 + m31*2^256 + m32*2^320 + m33*2^384 + m34*2^448
	*/
	var m [4][5]uint64
	for i := 0; i < 4; i++ {
		m[i] = mul64by256(a[i], b)
	}
	var r [8]uint64
	r[0], r[1], r[2], r[3], r[4] = m[0][0], m[0][1], m[0][2], m[0][3], m[0][4]
	for i := 1; i < 4; i++ {
		var carry uint64
		for j := 0; j < 5; j++ {
			r[i+j], carry = bits.Add64(r[i+j], m[i][j], carry)
		}
	}
	return r
}

// Mul a*b
func Mul(a, b Element) Element {
	r := mul(a, b)
	/*

	   // 2^320 = [4]uint64{0, 13963371537465362142, 17000232631950751529, 1089},
	   // 2^384 = [4]uint64{6946274766781380961, 1834951742433355726, 13963371537769226683, 486053787193498482},
	   // 2^448 = [4]uint64{1195137535819132669, 7639405055976263065, 2336821022378087212, 296249968807153639},
	*/
	r0 := removeTopBit(reduce5([5]uint64{r[0], r[1], r[2], r[3], r[4]}))
	r1 := removeTopBit(reduce5(mul64by256(r[5], [4]uint64{0, 13963371537465362142, 17000232631950751529, 1089})))
	r2 := removeTopBit(reduce5(mul64by256(r[6], [4]uint64{6946274766781380961, 1834951742433355726, 13963371537769226683, 486053787193498482})))
	r3 := removeTopBit(reduce5(mul64by256(r[7], [4]uint64{1195137535819132669, 7639405055976263065, 2336821022378087212, 296249968807153639})))

	s0 := removeTopBit(add(r0, r1))
	s1 := removeTopBit(add(r2, r3))

	return add(s0, s1)
}

// ElementToBigInt converts an element to a *big.Int.
func ElementToBigInt(a Element) (z *big.Int) {
	z = new(big.Int)
	for i := 3; i >= 0; i-- {
		z.Lsh(z, 64)
		z.Add(z, new(big.Int).SetUint64(a[i]))
	}
	return
}

// BigIntToElement converts a big.Int to an element.
func BigIntToElement(z *big.Int) (a Element) {
	z = new(big.Int).Mod(z, pBI) // Use a copy to avoid overwriting anything.
	mask := new(big.Int).SetUint64(0xffffffffffffffff)
	for i := 0; i < 4; i++ {
		a[i] = new(big.Int).And(z, mask).Uint64()
		z.Rsh(z, 64)
	}
	return
}
