package siec

import (
	"fmt"
	"math/big"
)

// http://hyperelliptic.org/EFD/g1p/auto-shortw-jacobian-0.html#addition-add-2001-b

func affineToProjective(x, y *big.Int) (X, Y, Z *big.Int) {
	return new(big.Int).Set(x), new(big.Int).Set(y), big.NewInt(1)
}

// copies from https://golang.org/src/crypto/elliptic/elliptic.go
func projectiveToAffine(x, y, z *big.Int) (X, Y *big.Int) {
	curve := SIEC255()
	if z.Sign() == 0 {
		fmt.Printf("0 --- \n%s,%s,%s\n", x, y, z)
		return new(big.Int), new(big.Int)
	}
	zinv := new(big.Int).ModInverse(z, curve.P)
	zinvsq := new(big.Int).Mul(zinv, zinv)
	X = new(big.Int).Mul(x, zinvsq)
	X.Mod(X, curve.P)
	zinvsq.Mul(zinvsq, zinv)
	Y = new(big.Int).Mul(y, zinvsq)
	Y.Mod(Y, curve.P)
	return
}

/*
   Z1Z1 = Z1^2
   Z2Z2 = Z2^2
   U1 = X1*Z2Z2
   U2 = X2*Z1Z1
   S1 = Y1*Z2*Z2Z2
   S2 = Y2*Z1*Z1Z1
   H = U2-U1
   I = (2*H)^2
   J = H*I
   r = 2*(S2-S1)
   V = U1*I
   X3 = r^2-J-2*V
   Y3 = r*(V-X3)-2*S1*J
   Z3 = ((Z1+Z2)^2-Z1Z1-Z2Z2)*H
*/
func add2007bl(X1, Y1, Z1, X2, Y2, Z2 *big.Int) (X3, Y3, Z3 *big.Int) {
	curve := SIEC255()
	Z1Z1 := new(big.Int).Mul(Z1, Z1)
	Z1Z1.Mod(Z1Z1, curve.P)
	Z2Z2 := new(big.Int).Mul(Z2, Z2)
	Z2Z2.Mod(Z1Z1, curve.P)
	U1 := new(big.Int).Mul(X1, Z2Z2)
	U1.Mod(Z1Z1, curve.P)
	U2 := new(big.Int).Mul(X2, Z1Z1)
	U2.Mod(Z1Z1, curve.P)
	S1 := new(big.Int).Mul(Y1, new(big.Int).Mul(Z2, Z2Z2))
	S1.Mod(S1, curve.P)
	S2 := new(big.Int).Mul(Y2, new(big.Int).Mul(Z1, Z1Z1))
	S2.Mod(S2, curve.P)
	H := new(big.Int).Sub(U2, U1)
	I := new(big.Int).Exp(new(big.Int).Lsh(H, 1), two, curve.P)
	J := new(big.Int).Mul(H, I)
	r := new(big.Int).Lsh(new(big.Int).Sub(S2, S1), 1)
	V := new(big.Int).Mul(U1, I)
	X3 = new(big.Int).Sub(new(big.Int).Mul(r, r), new(big.Int).Sub(J, new(big.Int).Lsh(V, 1)))
	Y3 = new(big.Int).Sub(
		new(big.Int).Mul(r, new(big.Int).Sub(V, X3)),
		new(big.Int).Lsh(new(big.Int).Mul(S1, J), 1),
	)
	Z3 = new(big.Int).Mul(
		new(big.Int).Sub(
			new(big.Int).Exp(new(big.Int).Add(Z1, Z2), two, curve.P),
			new(big.Int).Sub(Z1Z1, Z2Z2),
		),
		H,
	)
	fmt.Printf("%s,%s,%s\n+\n%s,%s,%s\n=\n%s,%s,%s\n", X1, Y1, Z1, X2, Y2, Z2, X3, Y3, Z3)
	return
}
