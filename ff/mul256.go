package ff

func mul256(x, y [4]uint64) (z [8]uint64) {
	/*
		B = 128
		x = x₁*2^B + x₀
		y = y₁*2^B + y₀
		z₂ = x₁y₁
		z₁ = (x₀-x₁)(y₁-y₀) + x₁y₁ + x₀y₀ = x₀y₁ + x₁y₀
		z₀ = x₀y₀
		x*y = z₂*2^(2B) + z₁*2^B + z₀
	*/
	x0 := [2]uint64{x[0], x[1]}
	x1 := [2]uint64{x[2], x[3]}
	y0 := [2]uint64{y[0], y[1]}
	y1 := [2]uint64{y[2], y[3]}
	z0 := mul128(x0, y0)
	z2 := mul128(x1, y1)
	x0y1 := mul128(x0, y1)
	x1y0 := mul128(x1, y0)
	var k1, k2 uint64
	z[0] = z0[0]
	z[1] = z0[1]
	z[2], k1 = addThree64(z0[2], x0y1[0], k1)
	z[2], k2 = addTwo64(z[2], x1y0[0])
	z[3], k1 = addThree64(z0[3], x0y1[1], k1+k2)
	z[3], k2 = addTwo64(z[3], x1y0[1])
	z[4], k1 = addThree64(z2[0], x0y1[2], k1+k2)
	z[4], k2 = addTwo64(z[4], x1y0[2])
	z[5], k1 = addThree64(z2[1], x0y1[3], k1+k2)
	z[5], k2 = addTwo64(z[5], x1y0[3])
	z[6], k1 = addTwo64(z2[2], k1+k2)
	z[7], _ = addTwo64(z2[3], k1)
	return
}

func addThree64(x, y, z uint64) (sum, carry uint64) {
	sum = x + y + z
	if (y > 0xffffffffffffffff-x) || (z > 0xffffffffffffffff-(x+y)) {
		carry = 1
	}
	return
}

func addTwo64(x, y uint64) (sum, carry uint64) {
	sum = x + y
	if y > 0xffffffffffffffff-x {
		carry = 1
	}
	return
}
