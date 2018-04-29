package ff

func add(a, b Element) (c Element) {
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
	return normalize(c)
}
