package ff

func neg(a Element) Element {
	var z uint64
	for j := 0; j < 3; j++ {
		if a[j] > pAsElement[j]+z {
			a[j] = pAsElement[j] - (a[j] + z)
			z = 1
		} else {
			a[j] = pAsElement[j] - (a[j] + z)
			z = 0
		}
	}
	a[3] = pAsElement[3] - (a[3] + z)
	return a
}
