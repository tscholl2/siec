package ff

func sub(a, b Element) (c Element) {
	// TODO: implement
	A, B := ToBigInt(a), ToBigInt(b)
	return FromBigInt(A.Sub(A, B))
}
