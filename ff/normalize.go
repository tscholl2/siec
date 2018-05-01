package ff

func normalize(a Element) Element {
	c := cmp(a, p)
	if c == -1 {
		return a
	}
	if c == 0 {
		return Element{0, 0, 0, 0}
	}
	return normalize(sub(a, p))
}
