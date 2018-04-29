package ff

func double(a Element) Element {
	//	return add(a, a)
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
	// TODO: inline mod
	if isGreaterThanOrEqualToP(a) {
		var z uint64
		if pAsElement[0] > a[0]+z {
			a[0] = a[0] - (pAsElement[0] + z)
			z = 1
		} else {
			a[0] = a[0] - (pAsElement[0] + z)
			z = 0
		}
		if pAsElement[1] > a[1]+z {
			a[1] = a[1] - (pAsElement[1] + z)
			z = 1
		} else {
			a[1] = a[1] - (pAsElement[1] + z)
			z = 0
		}
		if pAsElement[2] > a[2]+z {
			a[2] = a[2] - (pAsElement[2] + z)
			z = 1
		} else {
			a[2] = a[2] - (pAsElement[2] + z)
			z = 0
		}
		a[3] = a[3] - (pAsElement[3] + z)
	}
	return a
}
