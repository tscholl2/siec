package siec

import (
	"math/big"
	"testing"
)

func Test_projectiveToAffine(t *testing.T) {
	curve := SIEC255()
	var x, y, X, Y, Z, u, v *big.Int
	X, Y, Z = big.NewInt(5), big.NewInt(12), big.NewInt(1)
	x, y = curve.projectiveToAffine(X, Y, Z)
	u, v = big.NewInt(5), big.NewInt(12)
	if x.Cmp(u) != 0 || y.Cmp(v) != 0 {
		t.Errorf("got: (0x%x,0x%x) wanted (0x%x,0x%x)", x, y, u, v)
	}
}

func Test_affineToProjective(t *testing.T) {
	curve := SIEC255()
	var x, y, X, Y, Z, U, V, W *big.Int
	x, y = big.NewInt(5), big.NewInt(12)
	X, Y, Z = curve.affineToProjective(x, y)
	U, V, W = big.NewInt(5), big.NewInt(12), big.NewInt(1)
	if X.Cmp(U) != 0 || Y.Cmp(V) != 0 || Z.Cmp(W) != 0 {
		t.Errorf("got: (0x%x,0x%x,0x%x) wanted (0x%x,0x%x,0x%x)", X, Y, Z, U, V, W)
	}
}

func Test_add2007bl(t *testing.T) {
	curve := SIEC255()
	var X1, Y1, Z1, X2, Y2, Z2, X3, Y3, Z3, U, V, W *big.Int
	X1 = big.NewInt(5)
	Y1 = big.NewInt(12)
	Z1 = big.NewInt(1)
	X2, _ = new(big.Int).SetString("f0000000000000000000000007803cf1e000000000000000000f00f3cb5e78f", 16)
	Y2, _ = new(big.Int).SetString("21200000000000000000000001090869624000000000000000021221a611b4b6", 16)
	Z2, _ = new(big.Int).SetString("1", 16)
	X3, Y3, Z3 = curve.add2007bl(X1, Y1, Z1, X2, Y2, Z2)
	X3.Mod(X3, curve.P)
	Y3.Mod(Y3, curve.P)
	Z3.Mod(Z3, curve.P)
	U, _ = new(big.Int).SetString("3db00000000000000000000001ed8faa2b600000000000000003db3ea9ac13f6", 16)
	V, _ = new(big.Int).SetString("334e00000000000000000000019a7d07349c000000000000000335141da62f6f", 16)
	W, _ = new(big.Int).SetString("1e000000000000000000000000f0079e3c000000000000000001e01e796bcf14", 16)
	if X3.Cmp(U) != 0 || Y3.Cmp(V) != 0 || Z3.Cmp(W) != 0 {
		t.Errorf("got: (0x%x,0x%x,0x%x) wanted (0x%x,0x%x,0x%x)", X3, Y3, Z3, U, V, W)
	}
}

func Test_dbl2009l(t *testing.T) {
	curve := SIEC255()
	var X1, Y1, Z1, X3, Y3, Z3, U, V, W *big.Int
	X1, _ = new(big.Int).SetString("f0000000000000000000000007803cf1e000000000000000000f00f3cb5e78f", 16)
	Y1, _ = new(big.Int).SetString("21200000000000000000000001090869624000000000000000021221a611b4b6", 16)
	Z1, _ = new(big.Int).SetString("1", 16)
	X3, Y3, Z3 = curve.dbl2009l(X1, Y1, Z1)
	X3.Mod(X3, curve.P)
	Y3.Mod(Y3, curve.P)
	Z3.Mod(Z3, curve.P)
	U, _ = new(big.Int).SetString("3a2e9fc0000000000000000001d183c44aecff80000000000003a325161bd1c7", 16)
	V, _ = new(big.Int).SetString("1e40efa67c0000000000000000f20f2bed3b24c8f80000000001e42db5c9fc62", 16)
	W, _ = new(big.Int).SetString("2400000000000000000000000120092448000000000000000002402491b492b", 16)
	if X3.Cmp(U) != 0 || Y3.Cmp(V) != 0 || Z3.Cmp(W) != 0 {
		t.Errorf("got: (0x%x,0x%x,0x%x) wanted (0x%x,0x%x,0x%x)", X3, Y3, Z3, U, V, W)
	}
	x, y := curve.projectiveToAffine(X3, Y3, Z3)
	u, v := new(big.Int), new(big.Int)
	u.SetString("1b3c1ab2fff8cc396d2c1528fd5e7229b58c06888decb98657312b2a0abc3ef9", 16)
	v.SetString("14a8371e5d745fe915613984af6451eb655cb3150a91caa9bafc5781f18d1257", 16)
	if x.Cmp(u) != 0 || y.Cmp(v) != 0 {
		t.Errorf("got: (0x%x,0x%x) wanted (0x%x,0x%x)", x, y, u, v)
	}
}

func Test_projectiveScalarBaseMult(t *testing.T) {
	curve := SIEC255()
	x, y := curve.projectiveScalarBaseMult([]byte{0x4})
	u, _ := new(big.Int).SetString("1b3c1ab2fff8cc396d2c1528fd5e7229b58c06888decb98657312b2a0abc3ef9", 16)
	v, _ := new(big.Int).SetString("14a8371e5d745fe915613984af6451eb655cb3150a91caa9bafc5781f18d1257", 16)
	if x.Cmp(u) != 0 || y.Cmp(v) != 0 {
		t.Errorf("got: (0x%x,0x%x) wanted (0x%x,0x%x)", x, y, u, v)
	}
}

func BenchmarkScale_projective(b *testing.B) {
	curve := SIEC255()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		curve.projectiveScalarBaseMult(hash(1))
	}
}
