_=gp.eval("""
next_siec(M) =
{
  default(realprecision, 300);
  local(DISCS,a,b,d,q,t);
  DISCS = [3, 4, 7, 8, 11, 19, 43, 67, 163];
  if(M < 1681,error('M must be at least 1681'));
  a = floor(sqrt(4*M - 163));
  b = 2*a;
  for(t=a,b,for(i=1,length(DISCS),
    d = DISCS[i];
    q = t^2 + d;
    if(q%4 != 0,next);
    q = q/4;
    if (q<=M,next());
    if(ispseudoprimepower(q),return([q,t]));
  ););
  error('SIEC not found');
}
""".replace("\n"," "))
_=gp.eval("""
prev_siec(M) =
{
  default(realprecision, 300);
  local(DISCS,a,b,d,q,t);
  DISCS = [3, 4, 7, 8, 11, 19, 43, 67, 163];
  if(M < 1681,error('M must be at least 1681'));
  a = ceil(sqrt(4*M - 3));
  b = 82;
  forstep(t=a,b,-1,forstep(i=length(DISCS),1,-1,
    d = DISCS[i];
    q = t^2 + d;
    if(q%4 != 0,next);
    q = q/4;
    if (q>=M,next());
    if(ispseudoprimepower(q),return([q,t]));
  ););
  error('SIEC not found');
}
""".replace("\n"," "))

def next_siec(M):
    if M < 1621:
        raise UnimplementedError
    t = floor(sqrt(4*M - 163))
    while True:
        for d in [3,4,7,8,11,19,43,67,163]:
            q = t^2 + d
            if q%4 != 0:
                continue
            q //= 4;
            p,k = q.perfect_power()
            if q > M and p.is_prime(proof=False):
                return q,t
        t += 1

def prev_siec(M):
    if M < 1621:
        raise UnimplementedError
    t = ceil(sqrt(4*M - 3))
    while True:
        for d in [3,4,7,8,11,19,43,67,163][::-1]:
            q = t^2 + d
            if q%4 != 0:
                continue
            q //= 4;
            p,k = q.perfect_power()
            if q < M and p.is_prime(proof=False):
                return q,t
        t -= 1

assert list(prev_siec(2^80)) == [ZZ(a) for a in gp.prev_siec(2^80)]
assert list(next_siec(2^80)) == [ZZ(a) for a in gp.next_siec(2^80)]

q = 2^256
i = 0
while True:
    i += 1
    t,q = prev_siec(q)
    d = t^2 - 4*q
    if abs(d) > 4:
        N = [q + 1 - t, q + 1 + t]
    else:
        K = QuadraticField(d)
        pi = K.prime_above(q).gens_reduced()[0]
        N = [(pi*u - 1).norm() for u in K.roots_of_unity()]
    if any(ZZ(n).is_prime(proof=False) for n in N):
        print "FOUND SOMETHING"
        print q
        print sum(q.bits())
        break





q = 115792089237316195423570985008687907420430813942206838514045083354703965072401
A,B = 3888,0
E = EllipticCurve(GF(q),[A,B])
P = E([0,0])
assert P.order() == 2

# http://hyperelliptic.org/EFD/g1p/auto-jquartic-xxyzzr.html

p,_ = P.xy()
def E_to_C(Q):
    if Q == 0:
        return (0,1,1)
    if Q == P:
        return (0,-1,1)
    x,y = P.xy()
    return (2*(x-p), (2*x+p)*(x-p)^2-y^2, y)

e = -(3*p^2+4*A)/16
d = 3*p/4
# C: y^2 = ex^4 + 2ax^2z^2 + z^4
for Q in [E(0),E(P)]+[E.random_point() for _ in range(100)]:
    x,y,z = E_to_C(E.random_point())
    assert y^2 == e*x^4 + 2*a*x^2*z^2 + z^4