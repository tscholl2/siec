package siec

import (
	"crypto/elliptic"
	"math/big"
)

// Params returns the parameters for SIEC in Weierstrass a=-3 form
func Params() (params elliptic.CurveParams) {
	params.P, _ = new(big.Int).SetString("7fffffffffffffffffffffffffffddf40a09853f04c9246b1f1c11c8ad49dc91", 16)
	params.N, _ = new(big.Int).SetString("3fffffffffffffffffffffffffffeefaba09b5d37c42f6b9e90b9297cbef94d5", 16)
	params.B, _ = new(big.Int).SetString("0", 16)
	params.Gx, _ = new(big.Int).SetString("7fffffffffffffffffffffffffffddf40a09853f04c9246b1f1c11c8ad49dc90", 16)
	params.Gy, _ = new(big.Int).SetString("58779cd64069f261806367f8b481973c886a99152c9b80ac97a57cdf86750f9d", 16)
	params.BitSize = 255
	params.Name = "siec-w"
	return
}
