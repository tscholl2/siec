package ff

func mul256(x, y [4]uint64) (z [8]uint64) {
	/*
		x = x₁*2^B + x₀
		y = y₁*2^B + y₀
		z₂ = x₁y₁
		z₁ = (x₀-x₁)(y₁-y₀) + x₁y₁ + x₀y₀ = x₀y₁ + x₁y₀
		z₀ = x₀y₀
		xy = z₂*2^(2B) + z₁*2^B + z₀
	*/
	x0 := [2]uint64{x[0], x[1]}
	x1 := [2]uint64{x[2], x[3]}
	y0 := [2]uint64{y[0], y[1]}
	y1 := [2]uint64{y[2], y[3]}

	// compute z2
	z2 := mul128(x1, y1)
	// compute z0
	z0 := mul128(x0, y0)

	// compute z1
	var z1 [5]uint64
	// a = |(x₀-x₁)(y₁-y₀)|
	var a [4]uint64
	cx0x1 := cmp128(x0, x1)
	cy1y0 := cmp128(y1, y0)
	aIsNegative := (cx0x1 == -1 && cy1y0 == 1) || (cx0x1 == 1 && cy1y0 == -1)
	if aIsNegative {
		if cx0x1 == -1 {
			a = mul128(sub128(x1, x0), sub128(y1, y0))
		} else {
			a = mul128(sub128(x0, x1), sub128(y0, y1))
		}
	} else {
		if cx0x1 == -1 {
			a = mul128(sub128(x1, x0), sub128(y0, y1))
		} else {
			a = mul128(sub128(x0, x1), sub128(y1, y0))
		}
	}
	// Now add z₀ and z₂.
	var z0Plusz2 [5]uint64
	var k uint64
	for i := 0; i < 4; i++ {
		z0Plusz2[i] = z0[i] + z2[i] + k
		if (z0[i] == 0xffffffffffffffff && k == 1) || (z0[i]+k > 0xffffffffffffffff-z2[i]) {
			k = 1
		} else {
			k = 0
		}
	}
	z0Plusz2[4] = k
	k = 0
	// Now compute z₁.
	if aIsNegative {
		// z₁ = z₀ + z₂ - a
		for i := 0; i < 4; i++ {
			z1[i] = z0Plusz2[i] - (a[i] + k)
			if (z0Plusz2[i] == 0 && k == 1) || z0Plusz2[i]-k < a[i] {
				k = 1
			} else {
				k = 0
			}
		}
		z1[4] = z0Plusz2[4] - k
	} else {
		// z₁ = a + z₀ + z₂
		for i := 0; i < 4; i++ {
			z1[i] = a[i] + z0Plusz2[i] + k
			if (a[i] == 0xffffffffffffffff && k == 1) || a[i]+k > 0xffffffffffffffff-z0Plusz2[i] {
				k = 1
			} else {
				k = 0
			}
		}
		z1[4] = z0Plusz2[4] + k
	}

	/*
		a := mul128(x0, y1)
		b := mul128(x1, y0)
		for i := 0; i < 4; i++ {
			z1[i] = a[i] + b[i] + k
			if (a[i] == 0xffffffffffffffff && k == 1) || (a[i]+k) > 0xffffffffffffffff-b[i] {
				k = 1
			} else {
				k = 0
			}
		}
		z1[4] = k
	*/
	k = 0

	// compute z1

	z[0], z[1] = z0[0], z0[1]
	z[2] = z0[2] + z1[0]
	if z0[2] > 0xffffffffffffffff-z1[0] {
		k = 1
	}
	z[3] = z0[3] + z1[1] + k
	if (z0[3] == 0xffffffffffffffff && k == 1) || (z0[3]+k) > 0xffffffffffffffff-z1[1] {
		k = 1
	} else {
		k = 0
	}
	z[4] = z1[2] + z2[0] + k
	if (z1[2] == 0xffffffffffffffff && k == 1) || (z1[2]+k) > 0xffffffffffffffff-z2[0] {
		k = 1
	} else {
		k = 0
	}
	z[5] = z1[3] + z2[1] + k
	if (z1[3] == 0xffffffffffffffff && k == 1) || (z1[3]+k) > 0xffffffffffffffff-z2[1] {
		k = 1
	} else {
		k = 0
	}
	z[6] = z1[4] + z2[2] + k
	if (z1[4] == 0xffffffffffffffff && k == 1) || (z1[4]+k) > 0xffffffffffffffff-z2[2] {
		k = 1
	} else {
		k = 0
	}
	z[7] = z2[3] + k
	return
}

/*
func overflow3(x, y, k uint64) uint64 {
	if (x == 0xffffffffffffffff && k == 1) || (x+k) > 0xffffffffffffffff-y {
		return 1
	}
	return 0
}

func overflow(x, y uint64) uint64 {
	if x > 0xffffffffffffffff-y {
		return 1
	}
	return 0
}
*/

func cmp128(x, y [2]uint64) int {
	if x[1] > y[1] {
		return 1
	}
	if x[1] < y[1] {
		return -1
	}
	if x[0] > y[0] {
		return 1
	}
	if x[0] < y[0] {
		return -1
	}
	return 0
}

/*
func sub128(x, y [2]uint64) (z [2]uint64) {
	// x - y
	var k uint64
	z[0] = x[0] - y[0]
	if x[0] < y[0] {
		k = 1
	}
	z[1] = x[1] - (y[1] + k)
	return
}
*/
