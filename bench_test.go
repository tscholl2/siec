package siec

import (
	"crypto/elliptic"
	"crypto/sha256"
	"testing"

	"github.com/tscholl2/siec/edwards25519"
)

func BenchmarkDouble(b *testing.B) {
	curve := SIEC255()
	x, y := curve.ScalarBaseMult(hash(1))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.Double(x, y)
	}
}

func BenchmarkDoubleP256(b *testing.B) {
	curve := elliptic.P256()
	x, y := curve.ScalarBaseMult(hash(1))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.Double(x, y)
	}
}

func BenchmarkDoubleP224(b *testing.B) {
	curve := elliptic.P224()
	x, y := curve.ScalarBaseMult(hash(1))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.Double(x, y)
	}
}

func BenchmarkAdd(b *testing.B) {
	curve := SIEC255()
	x1, y1 := curve.ScalarBaseMult(hash(1))
	x2, y2 := curve.ScalarBaseMult(hash(2))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.Add(x1, y1, x2, y2)
	}
}

func BenchmarkAddP256(b *testing.B) {
	curve := elliptic.P256()
	x1, y1 := curve.ScalarBaseMult(hash(1))
	x2, y2 := curve.ScalarBaseMult(hash(2))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.Add(x1, y1, x2, y2)
	}
}

func BenchmarkAddP224(b *testing.B) {
	curve := elliptic.P224()
	x1, y1 := curve.ScalarBaseMult(hash(1))
	x2, y2 := curve.ScalarBaseMult(hash(2))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.Add(x1, y1, x2, y2)
	}
}

func BenchmarkScale(b *testing.B) {
	curve := SIEC255()
	k := hash(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.ScalarBaseMult(k)
	}
}

func BenchmarkScaleP256(b *testing.B) {
	curve := elliptic.P256()
	k := hash(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.ScalarBaseMult(k)
	}
}

func BenchmarkScaleP224(b *testing.B) {
	curve := elliptic.P224()
	k := hash(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.ScalarBaseMult(k)
	}
}

func BenchmarkScaleEd25519(b *testing.B) {
	arr := hash(1)
	arr[0] &= 248
	arr[31] &= 127
	arr[31] |= 64
	var A edwards25519.ExtendedGroupElement
	var hBytes [32]byte
	copy(hBytes[:], arr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		edwards25519.GeScalarMultBase(&A, &hBytes)
	}
}

func hash(i int) []byte {
	arr := sha256.Sum256([]byte{byte(i)})
	return arr[:]
}
