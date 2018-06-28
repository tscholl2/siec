# siec
Super-Isolated Elliptic Curve Implementation in Go

This package exports a super-isolated elliptic curve.
Over the base field 𝔽ₚ, the curve E does not admit any isogenies to other curves.

We can verify the curve properties in Sage.

```python
K.<isqrt3> = QuadraticField(-3)
pi = 2^127 + 2^25 + 2^12 + 2^6 + (1 - isqrt3)/2
p = ZZ(pi.norm())
N = ZZ((pi-1).norm())
E = EllipticCurve(GF(p),[0,19]) # E: y^2 = x^3 + 19
# p is a 255 bit prime with hamming weight 14
assert p.is_prime()
assert len(p.bits()) == 255
assert sum(p.bits()) == 14
# N is a 255 bit prime
assert N.is_prime()
assert len(N.bits()) == 255
# E has N points
assert E.count_points() == N
# The Frobenius endomorphism on E satisfies the same minimal polynomial as pi
assert E.frobenius_polynomial() == pi.minpoly()
# pi generates a maximal order
assert K.order([pi]).is_maximal()
# K has class number 1
assert K.class_number() == 1
```
