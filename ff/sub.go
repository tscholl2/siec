package ff

func sub(a, b Element) (c Element) {
	// return FromBigInt(new(big.Int).Sub(ToBigInt(a), ToBigInt(b)))
	// c = a - b
	s := isNegative(a)
	if s != isNegative(b) {
		a[3] = a[3] & mask63
		if s {
			// if a < 0 and b >= 0, then a - b = -(|a|+b)
			return neg(add(a, b))
		}
		// if a >= 0 and b < 0, then a - b = a+|b|
		b[3] = b[3] & mask63
		return add(a, b)
	}
	// a, b have same sign and it is recorded in s,
	// so we can take abs
	a[3] = a[3] & mask63
	b[3] = b[3] & mask63
	// if a >= b then use a - b, else -(b - a)
	for i := 3; i >= 0; i-- {
		if a[i] > b[i] {
			goto subtract
		}
		if a[i] < b[i] {
			a, b = b, a
			s = !s
			goto subtract
		}
	}
	// if a == b, return 0
	return
subtract:
	// We can now assume that a >= b >= 0.
	var z uint64
	for i := 0; i < 4; i++ {
		c[i] = a[i] - (b[i] + z)
		if a[i] < b[i]+z {
			z = 1
		} else {
			z = 0
		}
	}
	if s {
		return neg(c)
	}
	return
}
