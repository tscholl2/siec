package siec

import "math/big"

var (
	q, _ = new(big.Int).SetString("7fffffffffffffffffffffffffffddf40a09853f04c9246b1f1c11c8ad49dc91", 16)
)

type jacobiExtendedPoint struct {
	// https://hal-lirmm.ccsd.cnrs.fr/lirmm-00145805/document
	X, Y, Z, U, V, W *big.Int
}

func addExtended(P1, P2 jacobiExtendedPoint) (P3 jacobiExtendedPoint) {
	/*
		T1 = U1
		T2 = U2
		T3 = V1
		T4 = V2
		T5 = W1
		T6 = W2
		T7 = Y1
		T8 = Y2
		##############
		T9 = T7 * T8
		T7 = T7 + T3
		T8 = T8 + T4
		T3 = T3 * T4
		T7 = T7 * T8
		T7 = T7 - T9
		T7 = T7 - T3          # X3
		T4 = T1 * T2
		T8 = T5 * T6
		T1 = T1 + T5
		T2 = T2 + T6
		T5 = T1 * T2
		T5 = T5 - T4
		T5 = T5 - T8
		T4 = 3 * T4
		T1 = T8 - T4          # Z3
		T2 = T8 + T4
		T9 = T9 * T2
		T3 = 6 * T3
		T3 = T5 * T3
		T8 = T9 + T3          # Y3
		###############
		T2 = T7^2             # U3
		T4 = T1 * T7          # V3
		T9 = T1^2             # W3
		###############
		X3 = T7
		Z3 = T1
		Y3 = T8
		U3 = T2
		V3 = T4
		W3 = T9
	*/
	T1 := new(big.Int).Set(P1.U)
	T2 := new(big.Int).Set(P2.U)
	T3 := new(big.Int).Set(P1.V)
	T4 := new(big.Int).Set(P2.V)
	T5 := new(big.Int).Set(P1.W)
	T6 := new(big.Int).Set(P2.W)
	T7 := new(big.Int).Set(P1.Y)
	T8 := new(big.Int).Set(P2.Y)
	///////////
	T9 := new(big.Int).Mul(T7, T8)
	T9.Mod(T9, q)
	T7.Add(T7, T3)
	T8.Add(T8, T4)
	T3.Mul(T3, T4)
	T3.Mod(T3, q)
	T7.Mul(T7, T8)
	T7.Mod(T7, q)
	T7.Sub(T7, T9)
	T7.Sub(T7, T3)
	T4.Mul(T1, T2)
	T4.Mod(T4, q)
	T8.Mul(T5, T6)
	T8.Mod(T8, q)
	T1.Add(T1, T5)
	T2.Add(T2, T6)
	T5.Mul(T1, T2)
	T5.Mod(T5, q)
	T5.Sub(T5, T4)
	T5.Sub(T5, T8)
	T4.Mul(big.NewInt(3), T4) // TODO
	T1.Sub(T8, T4)
	T2.Add(T8, T4)
	T9.Mul(T9, T2)
	T9.Mod(T9, q)
	T3.Mul(big.NewInt(6), T3) // TODO
	T3.Mul(T5, T3)
	T3.Mod(T3, q)
	T8.Add(T9, T3)
	///////////////
	P3.X = T7
	P3.Z = T1
	P3.Y = T8
	///////////////
	T2.Exp(T7, big.NewInt(2), nil)
	T2.Mod(T2, q)
	T4.Mul(T1, T7)
	T4.Mod(T4, q)
	T9.Exp(T1, big.NewInt(2), nil)
	T9.Mod(T9, q)
	//////////////
	P3.U = T2
	P3.V = T4
	P3.W = T9
	return
}
