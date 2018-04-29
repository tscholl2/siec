package ff

// normalize converts an element from (-2^256,2^256) to [-(p-1)/2,(p-1)/2]
func normalize(a Element) Element {
	// If a < 0, then return -normalize(-a).
	if isNegative(a) {
		a[3] = a[3] & mask63
		a = normalize(a)
		a[3] = a[3] ^ bit64
		return a
	}
	//
	// Assume a >= 0.
	//
	// If a >= p, return normalize(a - p).
	for i := 3; i >= 0; i-- {
		if a[i] > pFF[i] {
			return normalize(sub(a, pFF))
		}
		if a[i] < pFF[i] {
			goto lessThanP
		}
	}
	return Element{0, 0, 0, 0}
lessThanP:
	// If p > a > (p-1)/2, return  a - p,
	/*
		a > (p-1)/2
		a >= (p+1)/2
		a-p >= (p+1)/2 - p
		= (-p+1)/2
		= -(p-1)/2
	*/
	for i := 3; i >= 0; i-- {
		if a[i] > pMinusOneOver2FF[i] {
			return sub(a, pFF)
		}
		if a[i] < pMinusOneOver2FF[i] {
			break
		}
	}
	// If (p-1)/2 >= a, return a
	return a
}
