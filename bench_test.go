package siec

import (
	"crypto/elliptic"
	"testing"
)

func BenchmarkDouble(b *testing.B) {
	curve := SIEC255()
	x, y := curve.ScalarBaseMult([]byte{0x40, 0x0, 0x41})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.Double(x, y)
	}
}

func BenchmarkDoubleP256(b *testing.B) {
	curve := elliptic.P256()
	x, y := curve.ScalarBaseMult([]byte{0x40, 0x0, 0x41})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.Double(x, y)
	}
}

func BenchmarkDoubleP224(b *testing.B) {
	curve := elliptic.P224()
	x, y := curve.ScalarBaseMult([]byte{0x40, 0x0, 0x41})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.Double(x, y)
	}
}

func BenchmarkAdd(b *testing.B) {
	curve := SIEC255()
	x1, y1 := curve.ScalarBaseMult([]byte{0x40, 0x0, 0x41})
	x2, y2 := curve.ScalarBaseMult([]byte{0x40, 0x10, 0x41})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.Add(x1, y1, x2, y2)
	}
}

func BenchmarkAddP256(b *testing.B) {
	curve := elliptic.P256()
	x1, y1 := curve.ScalarBaseMult([]byte{0x40, 0x0, 0x41})
	x2, y2 := curve.ScalarBaseMult([]byte{0x40, 0x10, 0x41})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.Add(x1, y1, x2, y2)
	}
}

func BenchmarkAddP224(b *testing.B) {
	curve := elliptic.P224()
	x1, y1 := curve.ScalarBaseMult([]byte{0x40, 0x0, 0x41})
	x2, y2 := curve.ScalarBaseMult([]byte{0x40, 0x10, 0x41})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.Add(x1, y1, x2, y2)
	}
}

func BenchmarkScale(b *testing.B) {
	curve := SIEC255()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.ScalarBaseMult([]byte{0x40, 0x0, 0x41})
	}
}

func BenchmarkScaleP256(b *testing.B) {
	curve := elliptic.P256()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.ScalarBaseMult([]byte{0x40, 0x0, 0x41})
	}
}

func BenchmarkScaleP224(b *testing.B) {
	curve := elliptic.P224()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.ScalarBaseMult([]byte{0x40, 0x0, 0x41})
	}
}
