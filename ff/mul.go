package ff

func mul(a, b Element) (c Element) {
	// TODO: implement
	A, B := ToBigInt(a), ToBigInt(b)
	return FromBigInt(A.Mod(A.Mul(A, B), pBI))
}

// mul64 multiplies 2 unsigned 64 bit integers and returns a 128 bit unsigned integer.
func mul64(x, y uint64) (z [2]uint64) {
	/*
		B = 2^32

		x = x₁*B + x₀
		y = y₁*B + y₀

		z₂ = x₁y₁
		z₁ = x₁y₀ + x₀y₁
		z₀ = x₀y₀

		xy = z₂*2^(2B) + z₁*2^B + z₀
	*/
	x1, x0 := x>>32, x&0xffffffff
	y1, y0 := y>>32, y&0xffffffff
	z[0], z[1] = x0*y0, x1*y1
	a := x1 * y0
	b := x0 * y1
	c := a + b
	// w = z1*2^B
	w := [2]uint64{c << 32, c >> 32}
	if a > 0xffffffffffffffff-b {
		w[1] = w[1] | bit33
	}
	// z + w
	if z[0] > 0xffffffffffffffff-w[0] {
		w[1]++
	}
	z[0] = z[0] + w[0]
	z[1] = z[1] + w[1]
	return
}

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
	z1 := add128(mul64(x[1], y[0]), mul64(x[0], y[1])) // TODO: use karatsuba, check overflow
	z1 = add128(add128(z1, z2), z0)                    // check overflow
	z[2], z[3] = z2[0], z2[1]
	z[0], z[1] = z0[0], z0[1]
	// add z1
	var k uint64
	if z[1] == 0xffffffffffffffff && z1[0] > 0 {
		k = 1
	}
	z[1] = z[1] + z1[0]
	// check for overflow with z[2]+z1[1]+k
	if (z[2] == 0xffffffffffffffff && (z1[1] > 0 || k == 1)) || (z[2]+k > 0xffffffffffffffff-z1[1]) {
		z[3]++
	}
	z[2] = z[2] + z1[1] + k
	return
}

func sub128(x, y [2]uint64) (z [2]uint64) {
	z[0] = x[0] - y[0]
	z[1] = x[1] - y[1]
	if x[0] < y[0] {
		z[1]--
	}
	return
}

func sub256(x, y [4]uint64) (z [4]uint64) {
	// TODO
	return
}

func add128(x, y [2]uint64) (z [2]uint64) {
	z[0] = x[0] + y[0]
	z[1] = x[1] + y[1]
	if x[0] > 0xffffffffffffffff-y[0] {
		z[1]++
	}
	return
}

func add256(x, y [4]uint64) (z [4]uint64) {
	// TODO
	return
}

/*
func mul(a, b Element) (c Element) {
	d := a
	for i := 0; i < 4; i++ {
		for j := uint8(0); j < 64; j++ {
			if (b[i]>>j)&1 == 1 {
				c = add(c, d)
			}
			d = double(d)
		}
	}
	return
}

func mul2(a, b Element) Element {
	return karatsuba(a, b)
}

		x = x₁B + x₀
		y = y₁B + y₀

		z₂ = x₁y₁
		z₁ = (x₀-x₁)(y₁-y₀) + x₁y₁ + x₀y₀
		z₀ = x₀y₀

		xy = z₂2^(2B) + z₁2^B + z₀
	}

func karatsuba(a, b Element) (c Element) {
	var x1, x0, y1, y0 Element
	var B uint8
	if a[3] > 0 || a[2] > 0 || b[3] > 0 || b[2] > 0 {
		x1 = Element{a[2], a[3], 0, 0}
		x0 = Element{a[0], a[1], 0, 0}
		y1 = Element{b[2], b[3], 0, 0}
		y0 = Element{b[0], b[1], 0, 0}
		B = 128
	} else {
		if a[1] > 0 || b[1] > 0 {
			x1 = Element{a[1], 0, 0, 0}
			x0 = Element{a[0], 0, 0, 0}
			y1 = Element{b[1], 0, 0, 0}
			y0 = Element{b[0], 0, 0, 0}
			B = 64
		} else {
			if (a[0]>>32) > 0 || (b[0]>>32) > 0 {
				x1 = Element{a[0] >> 32, 0, 0, 0}
				x0 = Element{a[0] & 0xffffffff, 0, 0, 0}
				y1 = Element{b[0] >> 32, 0, 0, 0}
				y0 = Element{b[0] & 0xffffffff, 0, 0, 0}
				B = 32
			} else {
				return Element{a[0] * b[0], 0, 0, 0}
			}
		}
	}
	z2 := karatsuba(x1, y1)
	z0 := karatsuba(x0, y0)
	z1 := add(add(karatsuba(add(x0, neg(x1)), add(x1, neg(y0))), z2), z0)
	for i := uint8(0); i < 2*B; i++ {
		z2 = double(z2)
	}
	for i := uint8(0); i < B; i++ {
		z2 = double(z1)
	}
	return add(z2, add(z1, z0))
}
*/
