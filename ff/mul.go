package ff

var (
	negative2ToThe256ModP = Element{0x1001040c208104, 0x0, 0x8004102, 0x0}
)

func mul(x, y Element) (z Element) {
	/*
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
	// Compute z₀,z₂
	z2 := to320(mul128(x1, y1))
	z0 := to320(mul128(x0, y0))
	// a = |(x₀-x₁)(y₁-y₀)|
	var a [5]uint64
	x0IsLessThanx1 := cmp(Element{x0[0], x0[1], 0, 0}, Element{x1[0], x1[1], 0, 0}) == -1
	y1IsLessThany0 := cmp(Element{y1[0], y1[1], 0, 0}, Element{y0[0], y0[1], 0, 0}) == -1
	aIsNegative := x0IsLessThanx1 != y1IsLessThany0
	if aIsNegative {
		if x0IsLessThanx1 {
			a = to320(mul128(sub128(x1, x0), sub128(y1, y0)))
		} else {
			a = to320(mul128(sub128(x0, x1), sub128(y0, y1)))
		}
	} else {
		if x0IsLessThanx1 {
			a = to320(mul128(sub128(x1, x0), sub128(y0, y1)))
		} else {
			a = to320(mul128(sub128(x0, x1), sub128(y1, y0)))
		}
	}
	// Now add z₀ and z₂.
	z0Plusz2 := add320(z0, z2)
	// Now compute z₁ = z₂ + z₀ +/- a.
	var z1 [5]uint64
	if aIsNegative {
		// z₁ = z₀ + z₂ - a
		// Note that because z₁ = x₀y₁ + x₁y₀ ≥ 0, we know that this subtraction is ok.
		z1 = sub320(z0Plusz2, a)
	} else {
		// z₁ = a + z₀ + z₂
		z1 = add320(z0Plusz2, a)
	}
	// z = z₂*2^(256) + z₁*2^(128) + z₀
	// TODO:
	// z5 is z as a 320 bit number
	z5 := add320(z0, [5]uint64{0, 0, z1[0], z1[1], 0})
	//z5 = add320(z5, mulBy256([5]uint64{z1[2], z1[3], 0, 0, 0}))
	//z5 = add320(z5, mulBy256(z2))
	// TODO: turn z5 into z
	z[0] = z5[0]
	return
}

func mulBy256(a [5]uint64) [5]uint64 {
	//
	// (-2^256 % p) = {0x1001040c208104,0,0x8004102}
	// (2^384 % p)  = {0x30c30d2479030608, 0x800c, 0x30030c24618300}
	//
	// -{a0,a1,a2,a3}*2^256 = {a0,a1}*(-2^256) - {a2,a3}*2^384
	//                      = ~~~~~~~ζ~~~~~~~~ - ~~~~~~ω~~~~~~
	//
	// ζ = {a0,a1}*{0x1001040c208104,0,0x8004102}
	//   = a0*{0x1001040c208104,0,0x8004102} + {0,a1}*0x1001040c208104 + (a1*0x8004102)*2^192
	//     ~~~~~~~~~~~~~~~ζ₀~~~~~~~~~~~~~~~~   ~~~~~~~~~~ζ₁~~~~~~~~~~~
	//   = ζ₀ + ζ₁ + (a1*0x8004102={b0,b1})*2^192
	//   = ζ₀ + ζ₁ + {0,0,0,b0} + b1*2^256
	//               ~~~~ζ₂~~~~
	//   = ζ₀ + ζ₁ + ζ₂ + b1*2^256
	//   = ζ₀ + ζ₁ + ζ₂ - b1*{0x1001040c208104,0,0x8004102}
	//                    ~~~~~~~~~~~~~~~ζ₃~~~~~~~~~~~~~~~~
	//   = ζ₀ + ζ₁ + ζ₂ - ζ₃
	//
	// ω = {a2,a3}*{0x30c30d2479030608,0x800c,0x30030c24618300}
	//   = a2*{0x30c30d2479030608,0x800c,0x30030c24618300} + {0,a3}*{0x30c30d2479030608,0x800c} + {0,a3}*{0,0,0x30030c24618300}
	//     ~~~~~~~~~~~~~~~~~~~~~~ω₀~~~~~~~~~~~~~~~~~~~~~~~ + ~~~~~~~~~~~~~~~~~~ω₁~~~~~~~~~~~~~~
	//   = ω₀ + ω₁ + {0,a3}*{0,0,0x30030c24618300}
	//   = ω₀ + ω₁ + (a3*0x30030c24618300={c1,c2})*2^192
	//   = ω₀ + ω₁ + {0,0,0,c1} + c2*2^256
	//   = ω₀ + ω₁ + ~~~~ω₂~~~~ - c2*{0x1001040c208104,0,0x8004102}
	//   = ω₀ + ω₁ + ω₂ - c2*{0x1001040c208104,0,0x8004102}
	//                    ~~~~~~~~~~~~~~ω₃~~~~~~~~~~~~~~~~~
	//   = ω₀ + ω₁ + ω₂ - ω₃
	//
	var ζ [4]Element
	// ζ₀
	ζ0a := mul64(a[0], 0x1001040c208104)
	ζ0b := mul64(a[0], 0x8004102)
	ζ[0][0], ζ[0][1], ζ[0][2], ζ[0][3] = ζ0a[0], ζ0a[1], ζ0b[0], ζ0b[1]
	// ζ₁
	ζ1a := mul64(a[1], 0x1001040c208104)
	ζ[1][1], ζ[1][2] = ζ1a[0], ζ1a[1]
	// ζ₂
	b := mul64(a[1], 0x8004102)
	ζ[2][3] = b[0]
	// ζ₃
	ζ3a := mul64(b[1], 0x1001040c208104)
	ζ3b := mul64(b[1], 0x8004102)
	ζ[3][0], ζ[3][1], ζ[3][2], ζ[3][3] = ζ3a[0], ζ3a[1], ζ3b[0], ζ3b[1]
	// TODO: think about normalization and check overflow
	ζall := add(normalize(ζ[0]), normalize(ζ[1]))
	ζall = add(ζall, normalize(ζ[2]))
	ζall = sub(ζall, ζ[3])
	if a[2] == 0 && a[3] == 0 {
		return [5]uint64{0, 0, 0, 0, 0} //TODO: ζall
	}
	return [5]uint64{0, 0, 0, 0, 0}
}

func sub128(x, y [2]uint64) (z [2]uint64) {
	// x - y
	var k uint64
	for i := 0; i < 2; i++ {
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

func to320(x [4]uint64) (z [5]uint64) {
	z[0], z[1], z[2], z[3] = x[0], x[1], x[2], x[3]
	return
}

func add320(x, y [5]uint64) (z [5]uint64) {
	var k uint64
	for i := 0; i < 5; i++ {
		z[i] = x[i] + y[i] + k
		if (x[i] == 0xffffffffffffffff && k == 1) || x[i]+k > 0xffffffffffffffff-y[i] {
			k = 1
		} else {
			k = 0
		}
	}
	return
}

func sub320(x, y [5]uint64) (z [5]uint64) {
	var k uint64
	for i := 0; i < 5; i++ {
		z[i] = x[i] - (y[i] + k)
		if (x[i] == 0 && k == 1) || x[i]-k < y[i] {
			k = 1
		} else {
			k = 0
		}
	}
	return
}

func cmp320(a [5]uint64, b [5]uint64) int {
	for i := 4; i >= 0; i-- {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}
	return 0
}
