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
assert sum(p.bits()) == 14
assert len(p.bits()) == 255
assert p.is_prime()
assert N.is_prime()
E = EllipticCurve(GF(p),[0,19])
assert E.count_points() == N
```
