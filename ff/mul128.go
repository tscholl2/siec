package ff

// mul128 multiplies 2 unsigned 128bit integers and returns a 256 bit unsigned integer.
func mul128(x, y [2]uint64) (z [4]uint64) {
	/*
		x = x₁B + x₀
		y = y₁B + y₀
		z₂ = x₁y₁
		z₁ = (x₀-x₁)(y₁-y₀) + x₁y₁ + x₀y₀ = x₀y₁ + x₁y₀
		z₀ = x₀y₀
		xy = z₂*2^(2B) + z₁*2^B + z₀
	*/
	z2 := mul64(x[1], y[1])
	z0 := mul64(x[0], y[0])

	z1a := mul64(x[1], y[0])
	z1b := mul64(x[0], y[1])

	// z1 = z1a + z1b
	// z1 := add256([4]uint64{0, z1a[0], z1a[1], 0}, [4]uint64{0, z1b[0], z1b[1], 0})
	var z1 [4]uint64
	var k uint64
	z1[1] = z1a[0] + z1b[0]
	if z1a[0] > 0xffffffffffffffff-z1b[0] {
		k = 1
	}
	z1[2] = z1a[1] + z1b[1] + k
	if (z1a[1] == 0xffffffffffffffff && k == 1) || z1a[1]+k > 0xffffffffffffffff-z1b[1] {
		z1[3] = 1
	}

	/*
	   		[ z0[0], z0[1], z2[0], z2[1] ]
	   +  [   0  , z1[1], z1[2], z1[3] ]
	   __________________________________
	                                   z
	*/

	z[0] = z0[0]
	z[1] = z0[1] + z1[1]
	if z0[1] > 0xffffffffffffffff-z1[1] {
		k = 1
	} else {
		k = 0
	}
	z[2] = z2[0] + z1[2] + k
	if (z2[0] == 0xffffffffffffffff && k == 1) || z2[0]+k > 0xffffffffffffffff-z1[2] {
		k = 1
	} else {
		k = 0
	}
	z[3] = z2[1] + z1[3] + k
	return
}

// mul128 multiplies 2 unsigned 128bit integers and returns a 256 bit unsigned integer.
func mul128_2(x, y [2]uint64) (z [4]uint64) {
	/*
		x = x₁*2^B + x₀
		y = y₁*2^B + y₀
		z₂ = x₁y₁
		z₁ = (x₀-x₁)(y₁-y₀) + x₁y₁ + x₀y₀ = x₀y₁ + x₁y₀
		z₀ = x₀y₀
		xy = z₂*2^(2B) + z₁*2^B + z₀
	*/
	z2 := mul64(x[1], y[1])
	z0 := mul64(x[0], y[0])

	var a [2]uint64 // a = |(x₀-x₁)(y₁-y₀)|
	aIsNegative := (x[0] < x[1] && y[1] > y[0]) || (x[0] > x[1] && y[1] < y[0])
	if aIsNegative {
		if x[0] < x[1] {
			a = mul64(x[1]-x[0], y[1]-y[0])
		} else {
			a = mul64(x[0]-x[1], y[0]-y[1])
		}
	} else {
		if x[0] < x[1] {
			a = mul64(x[1]-x[0], y[0]-y[1])
		} else {
			a = mul64(x[0]-x[1], y[1]-y[0])
		}
	}
	// now add z0 and z2
	var z0Plusz2 [4]uint64
	var k uint64
	z0Plusz2[0] = z0[0] + z2[0]
	if z0[0] > 0xffffffffffffffff-z2[0] {
		k = 1
	}
	z0Plusz2[1] = z0[1] + z2[1] + k
	if (z0[1] == 0xffffffffffffffff && k > 0) || z0[1]+k > 0xffffffffffffffff-z2[1] {
		z0Plusz2[2] = 1
	}
	k = 0
	var z1 [4]uint64
	// now add (or subtract) a to z0+z2
	if aIsNegative {
		// z1 = z0+z2 - a
		for i := 0; i < 2; i++ {
			z1[i] = z0Plusz2[i] - (a[i] + k)
			if (z0Plusz2[i] == 0 && k == 1) || z0Plusz2[i]-k < a[i] {
				k = 1
			} else {
				k = 0
			}
		}
		z1[2] = z0Plusz2[2] - k
	} else {
		// z1 = a + z0+z2
		for i := 0; i < 2; i++ {
			z1[i] = a[i] + z0Plusz2[i] + k
			if (a[i] == 0xffffffffffffffff && k == 1) || a[i]+k > 0xffffffffffffffff-z0Plusz2[i] {
				k = 1
			} else {
				k = 0
			}
		}
		z1[2] = z0Plusz2[2] + k
	}
	k = 0
	/*
	   		[ z0[0], z0[1], z2[0], z2[1] ]
	   +  [   0  , z1[0], z1[1], z1[2] ]
	   __________________________________
	                                   z
	*/
	z[0] = z0[0]
	z[1] = z0[1] + z1[0]
	if z0[1] > 0xffffffffffffffff-z1[0] {
		k = 1
	} else {
		k = 0
	}
	z[2] = z2[0] + z1[1] + k
	if (z2[0] == 0xffffffffffffffff && k == 1) || z2[0]+k > 0xffffffffffffffff-z1[1] {
		k = 1
	} else {
		k = 0
	}
	z[3] = z2[1] + z1[2] + k
	return
}
