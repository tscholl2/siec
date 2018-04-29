package ff

import "math/big"

const (
	bit64  = uint64(1) << 63            // 2^63
	bit63  = uint64(1) << 62            // 2^62
	mask64 = uint64(0xffffffffffffffff) // 2^64-1
	mask63 = mask64 - bit64             // 2^63 - 1
)

var (
	// p
	pAsBigInt, _ = new(big.Int).SetString("28948022309329048855892746252183396360603931420023084536990047309120118726721", 10)
	// p
	pAsElement = Element{1126179130581057, 9223372036854775808, 33558592, 4611686018427387904}
	// (p-1)/2
	pMinusOneOver2AsBigInt, _ = new(big.Int).SetString("14474011154664524427946373126091698180301965710011542268495023654560059363360", 10)
	// (p-1)/2 : note this is the positive, so not in our canonical form
	pMinusOneOver2AsElement = Element{563089565290528, 4611686018427387904, 16779296, 2305843009213693952}
	// (p+1)/2
	pPlusOneOver2AsElement = Element{563089565290529, 4611686018427387904, 16779296, 2305843009213693952}
)

// Element represents a single element a ∈ 𝔽ₚ,
// where p = 28948022309329048855892746252183396360603931420023084536990047309120118726721.
// Elements are stored as an array of four 64 bit unsigned integers.
// Note that p is only 255 bits, so we use the top bit for the sign.
// Specifically,
//
//     element = {a,b,c,d} = (- if top bit of d is set else +)(a + b*2^64 + c*2^128 + (d&0x7fffffffffffffff)*2^192)
//
// The mask is because the top bit of d is 1 represents the sign of the element.
// This representation can store all integers in the range (-2^256,2^256).
//
// Normalized Elements are represented as the unique representative in the range
// [-(p-1)/2, (p-1)/2]. For example:
//
//     0 = {0,0,0,0}
//     1 = {1,0,0,0}
//     -1 = {1,0,0,0x08000000000000000}
//     1 + 2*2^64 + 3*2^128 = {1,2,3,0}
//
// Note: because p is only 255 bits and we use the top bit for sign, this means
// the second-to-top bit should always be empty in the normalized form.
type Element [4]uint64

// ToBigInt Converts an Element to a *big.Int.
func ToBigInt(a Element) *big.Int {
	z := new(big.Int).SetUint64(a[3] & mask63)
	z.Lsh(z, 64)
	z.Add(z, new(big.Int).SetUint64(a[2]))
	z.Lsh(z, 64)
	z.Add(z, new(big.Int).SetUint64(a[1]))
	z.Lsh(z, 64)
	z.Add(z, new(big.Int).SetUint64(a[0]))
	if a[3]&bit64 == 1 {
		z.Neg(z)
	}
	return z
}

// FromBigInt converts a *big.Int to an Element.
func FromBigInt(n *big.Int) (a Element) {
	z := new(big.Int).Mod(n, pAsBigInt)
	w := new(big.Int)
	low64 := new(big.Int).SetUint64(0xffffffffffffffff)
	a[0] = w.And(z, low64).Uint64()
	a[1] = w.And(z.Rsh(z, 64), low64).Uint64()
	a[2] = w.And(z.Rsh(z, 64), low64).Uint64()
	a[3] = w.And(z.Rsh(z, 64), low64).Uint64()
	if z.Sign() == -1 {
		a[3] = a[3] | bit64
	}
	return
}

// normalize converts an element from (-2^256,2^256) to [-(p-1)/2,(p-1)/2]
func normalize(a Element) Element {
	if a[3]&bit64 == 1 {
		a[3] = a[3] & mask63
		a = normalize(a)
		a[3] = a[3] | bit64
	}

	// Assume a >= 0.
	// if a >= p, return normalize(a - p).
	// if p > a > (p-1)/2, return -(p - a),
	// if (p-1)/2 >= a, return a

	// check bit254 for quick check to see if bigger than (p-1)/2
	for i := 3; i >= 0; i-- {
		if a[i] > pMinusOneOver2AsElement[i] {
			break
		}
		if a[i] < pMinusOneOver2AsElement[i] {
			return a
		}
	}
	var z uint64
	for i := 3; i >= 0; i-- {
		if a[i] > pMinusOneOver2AsElement[i]+z {
			a[i] = pMinusOneOver2AsElement[i] - (a[i] + z)
			z = 1
		} else {
			a[i] = pMinusOneOver2AsElement[i] - (a[i] + z)
			z = 0
		}
	}
	a[3] = a[3] | bit64
	return a
}
