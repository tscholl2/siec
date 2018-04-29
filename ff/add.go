package ff

func add(a, b Element) (c Element) {
	var z uint64
	c[0] = a[0] + b[0] + z
	// check if overflowed
	if (a[0] == 0xffffffffffffffff && (b[0] > 0 || z == 1)) || (a[0]+z > 0xffffffffffffffff-b[0]) {
		z = 1
	} else {
		z = 0
	}
	c[1] = a[1] + b[1] + z
	// check if overflowed
	if (a[1] == 0xffffffffffffffff && (b[1] > 0 || z == 1)) || (a[1]+z > 0xffffffffffffffff-b[1]) {
		z = 1
	} else {
		z = 0
	}
	c[2] = a[2] + b[2] + z
	// check if overflowed
	if (a[2] == 0xffffffffffffffff && (b[2] > 0 || z == 1)) || (a[2]+z > 0xffffffffffffffff-b[2]) {
		z = 1
	} else {
		z = 0
	}
	c[3] = a[3] + b[3] + z
	// inline mod is much faster?
	if isGreaterThanOrEqualToP(c) {
		z = 0
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
	return
}
