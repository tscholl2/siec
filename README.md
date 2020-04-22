# siec
Super-Isolated Elliptic Curve Implementation in Go

This package exports a super-isolated elliptic curve.
Over the base field 𝔽ₚ, the curve E does not admit any isogenies to other curves.

We can verify the curve properties in Sage.

```python
# Curve

K.<I> = QuadraticField(-1)
pi = I - 0xb504f333f9de6484597d89b3754aa68c
q = ZZ(pi.norm()) # 7fffffffffffffffffffffffffffddf40a09853f04c9246b1f1c11c8ad49dc91
assert q.is_prime()
n = ZZ((pi-1).norm())
assert n%2 == 0 and ZZ(n/2).is_prime()
r = ZZ(n/2) # 3fffffffffffffffffffffffffffeefaba09b5d37c42f6b9e90b9297cbef94d5
E = EllipticCurve(GF(q),[-12,0])
assert E.frobenius_polynomial() == pi.minpoly()
assert q.nbits() == 255
assert K.order([pi]).is_maximal()
G = E([2,0x2d413cccfe779921165f626cdd52a9a30])
assert 2*G != 0
assert r*G == 0
i_mod_q = 0xb504f333f9de6484597d89b3754aa68c
assert i_mod_q^2 % q == -1 % q
i_mod_r = 0xb504f333f9de6484597d89b3754aa68d
assert i_mod_r^2 % r == -1 % r
assert i_mod_r*G == E([-G[0],i_mod_q*G[1]])
```
