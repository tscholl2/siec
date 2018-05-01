package ff

// cmp returns 1 if a > b, 0 if a = b, and -1 if a < b.
func cmp(a [4]uint64, b [4]uint64) int {
	for i := 3; i >= 0; i-- {
		if a[i] > b[i] {
			return 1
		} else if a[i] < b[i] {
			return -1
		}
	}
	return 0
}
