package ff

func double(a Element) Element {
	// TODO: optimize
	// return add(a, a)
	if isNegative(a) {
		return neg(double(neg(a)))
	}
	a[3] = a[3] << 1
	if a[2]>>63 == 1 {
		a[3] = a[3] + 1
	}
	a[2] = a[2] << 1
	if a[1]>>63 == 1 {
		a[2] = a[2] + 1
	}
	a[1] = a[1] << 1
	if a[0]>>63 == 1 {
		a[1] = a[1] + 1
	}
	a[0] = a[0] << 1
	return a
}
