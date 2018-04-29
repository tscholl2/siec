package ff

func add(a, b Element) (c Element) {
	// TODO: implement
	A, B := ToBigInt(a), ToBigInt(b)
	return FromBigInt(A.Add(A, B))
	/*
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
	*/
}
