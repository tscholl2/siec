package main

import (
	"math/big"
	"testing"
)

func TestDouble(t *testing.T) {
	c := SIEC255()
	x1, y1 := big.NewInt(5), big.NewInt(12)
	x, y := c.Double(x1, y1)
	u, _ := new(big.Int).SetString("6784692728748995825599862402855483522016546426567910438357042338075027826575", 10)
	v, _ := new(big.Int).SetString("14982863109320699114866362806305859444453206692004135551371801829915686450358", 10)
	if x.Cmp(u) != 0 {
		t.Errorf("SIEC255Params.Add() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("SIEC255Params.Add() gotY = %v, want %v", y, v)
	}
}

func TestScale(t *testing.T) {
	c := SIEC255()
	x, y := c.ScalarBaseMult(big.NewInt(2).Bytes())
	u, _ := new(big.Int).SetString("6784692728748995825599862402855483522016546426567910438357042338075027826575", 10)
	v, _ := new(big.Int).SetString("14982863109320699114866362806305859444453206692004135551371801829915686450358", 10)
	if x.Cmp(u) != 0 {
		t.Errorf("SIEC255Params.ScalarBaseMult() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("SIEC255Params.ScalarBaseMult() gotY = %v, want %v", y, v)
	}
}
