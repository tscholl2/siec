package ff

func add(x, y Element) (z Element) {
	// x + y
	var k uint64
	for i := 0; i < 4; i++ {
		z[i] = x[i] + y[i] + k
		// if x[i] + y[i]+ k > 0xffffffffffffffff
		if (x[i] == 0xffffffffffffffff && k == 1) || x[i]+k > 0xffffffffffffffff-y[i] {
			k = 1
		} else {
			k = 0
		}
	}
	return
}
