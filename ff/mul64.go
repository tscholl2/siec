package ff

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
		w[1] = w[1] | 0x100000000
	}
	// z + w
	if z[0] > 0xffffffffffffffff-w[0] {
		w[1]++
	}
	z[0] = z[0] + w[0]
	z[1] = z[1] + w[1]
	return
}
