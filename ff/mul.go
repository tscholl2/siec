package ff

func mul(a, b Element) (c Element) {
	d := a
	for i := 0; i < 4; i++ {
		for j := uint8(0); j < 64; j++ {
			if (b[i] >> j) > 0 {
				c = add(c, d)
			}
			d = double(d)
		}
	}
	return
}
