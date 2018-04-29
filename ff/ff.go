package ff

import "math/big"

var pAsElement = Element{1126179130581057, 9223372036854775808, 33558592, 4611686018427387904}

type Element [4]uint64

func ToBigInt(a Element) *big.Int {
	z := new(big.Int).SetUint64(a[3])
	z.Lsh(z, 64)
	z.Add(z, new(big.Int).SetUint64(a[2]))
	z.Lsh(z, 64)
	z.Add(z, new(big.Int).SetUint64(a[1]))
	z.Lsh(z, 64)
	z.Add(z, new(big.Int).SetUint64(a[0]))
	return z
}

func FromBigInt(n *big.Int) (a Element) {
	z := new(big.Int).Set(n)
	w := new(big.Int)
	low64 := new(big.Int).SetUint64(0xffffffffffffffff)
	a[0] = w.And(z, low64).Uint64()
	a[1] = w.And(z.Rsh(z, 64), low64).Uint64()
	a[2] = w.And(z.Rsh(z, 64), low64).Uint64()
	a[3] = w.And(z.Rsh(z, 64), low64).Uint64()
	return
}

// TODO: force inline
func mod(c Element) Element {
	if isGreaterThanOrEqualToP(c) {
		var z uint64
		if pAsElement[0] > c[0]+z {
			c[0] = c[0] - (pAsElement[0] + z)
			z = 1
		} else {
			c[0] = c[0] - (pAsElement[0] + z)
			z = 0
		}
		if pAsElement[1] > c[1]+z {
			c[1] = c[1] - (pAsElement[1] + z)
			z = 1
		} else {
			c[1] = c[1] - (pAsElement[1] + z)
			z = 0
		}
		if pAsElement[2] > c[2]+z {
			c[2] = c[2] - (pAsElement[2] + z)
			z = 1
		} else {
			c[2] = c[2] - (pAsElement[2] + z)
			z = 0
		}
		c[3] = c[3] - (pAsElement[3] + z)
	}
	return c
}

// TODO: force inline
func isGreaterThanOrEqualToP(a Element) bool {
	for i := 3; i >= 0; i-- {
		if a[i] > pAsElement[i] {
			return true
		}
		if a[i] < pAsElement[i] {
			return false
		}
	}
	return true
}
