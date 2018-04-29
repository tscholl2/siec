package ff

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

/*
		x = x₁B + x₀
		y = y₁B + y₀

		z₂ = x₁y₁
		z₁ = (x₀-x₁)(x₁-y₀) + x₁y₁ + x₀y₀
		z₀ = x₀y₀

		xy = z₂B^2 + z₁B + z₀
	}
*/
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
	for i := uint8(0); i < B; i++ {
		z2 = double(z2)
	}
	for i := uint8(0); i < B/2; i++ {
		z2 = double(z1)
	}
	return add(z2, add(z1, z0))
}
