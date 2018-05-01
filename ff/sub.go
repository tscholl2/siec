package ff

func sub(x, y Element) (z Element) {
	// x - y
	var k uint64
	for i := 0; i < 4; i++ {
		z[i] = x[i] - (y[i] + k)
		// if x[i] - (y[i]+k) < 0
		if (x[i] == 0 && k == 1) || x[i]-k < y[i] {
			k = 1
		} else {
			k = 0
		}
	}
	return
}
