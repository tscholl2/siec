package ff

func add(a, b Element) (c Element) {
	// return FromBigInt(new(big.Int).Add(ToBigInt(a), ToBigInt(b)))
	// c = a + b
	if isNegative(a) != isNegative(b) {
		b[3] = b[3] ^ bit64
		return sub(a, b)
	}
	var z uint64
	for i := 0; i < 4; i++ {
		c[i] = a[i] + b[i] + z
		// check if overflowed
		if (a[i] == mask64 && (b[i] > 0 || z == 1)) || (a[i]+z > mask64-b[i]) {
			z = 1
		} else {
			z = 0
		}
	}
	c[3] = c[3] | (a[3] & bit64)
	return
}
