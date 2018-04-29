package ff

func neg(a Element) Element {
	if a[3]&mask63 == 0 && a[2] == 0 && a[1] == 0 && a[0] == 0 {
		return Element{}
	}
	a[3] = a[3] ^ bit64
	return a
}
