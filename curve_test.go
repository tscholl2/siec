package siec

import (
	"crypto/elliptic"
	"math/big"
	"reflect"
	"testing"
)

func TestDouble(t *testing.T) {
	curve := SIEC255()
	x1, y1 := curve.Gx, curve.Gy
	x, y := curve.Double(x1, y1)
	u, _ := new(big.Int).SetString("6784692728748995825599862402855483522016546426567910438357042338075027826575", 10)
	v, _ := new(big.Int).SetString("14982863109320699114866362806305859444453206692004135551371801829915686450358", 10)
	if x.Cmp(u) != 0 {
		t.Errorf("curve.Double() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("curve.Double() gotY = %v, want %v", y, v)
	}
	x, y = curve.Double(x, y)
	u, _ = new(big.Int).SetString("12318642006867402687195826566147291859634823582672295191656499276835526033145", 10)
	v, _ = new(big.Int).SetString("9343467693237486709905252998911952863134805995110526737200728195882424275543", 10)
	if x.Cmp(u) != 0 {
		t.Errorf("curve.Double() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("curve.Double() gotY = %v, want %v", y, v)
	}
}

func TestAdd(t *testing.T) {
	curve := SIEC255()
	x1, _ := new(big.Int).SetString("5", 10)
	y1, _ := new(big.Int).SetString("12", 10)
	x2, _ := new(big.Int).SetString("6784692728748995825599862402855483522016546426567910438357042338075027826575", 10)
	y2, _ := new(big.Int).SetString("14982863109320699114866362806305859444453206692004135551371801829915686450358", 10)
	x, y := curve.Add(x1, y1, x2, y2)
	u, _ := new(big.Int).SetString("15508762693928266726769396085241920452527964114034020358226683933290603861217", 10)
	v, _ := new(big.Int).SetString("14902066091748681388520350407103675753252739566337731413744921722214701412361", 10)
	if x.Cmp(u) != 0 {
		t.Errorf("curve.Add() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("curve.Add() gotY = %v, want %v", y, v)
	}
}

func TestIsOnCurve(t *testing.T) {
	curve := SIEC255()
	x, _ := new(big.Int).SetString("5", 10)
	y, _ := new(big.Int).SetString("12", 10)
	if !curve.IsOnCurve(x, y) {
		t.Errorf("curve.IsOnCurve() got = %v, want %v", false, true)
	}
	y.Add(y, big.NewInt(1))
	if curve.IsOnCurve(x, y) {
		t.Errorf("curve.IsOnCurve() got = %v, want %v", true, false)
	}
	x, _ = new(big.Int).SetString("6784692728748995825599862402855483522016546426567910438357042338075027826575", 10)
	y, _ = new(big.Int).SetString("14982863109320699114866362806305859444453206692004135551371801829915686450358", 10)
	if !curve.IsOnCurve(x, y) {
		t.Errorf("curve.IsOnCurve() got = %v, want %v", false, true)
	}
	x, _ = new(big.Int).SetString("15508762693928266726769396085241920452527964114034020358226683933290603861217", 10)
	y, _ = new(big.Int).SetString("14902066091748681388520350407103675753252739566337731413744921722214701412361", 10)
	if !curve.IsOnCurve(x, y) {
		t.Errorf("curve.IsOnCurve() got = %v, want %v", false, true)
	}
}

func TestScale(t *testing.T) {
	curve := SIEC255()
	x, y := curve.ScalarBaseMult([]byte{0x40})
	u, _ := new(big.Int).SetString("22784956368772284587014129354783824301122808969944262244690262244567645543628", 10)
	v, _ := new(big.Int).SetString("15122902039027963115899162435007755480175091927763482548168795824636596047001", 10)
	if x.Cmp(u) != 0 {
		t.Errorf("curve.ScalarBaseMult() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("curve.ScalarBaseMult() gotY = %v, want %v", y, v)
	}
}

func TestScaleP256(t *testing.T) {
	// Sage:
	// P256 = EllipticCurve(GF(115792089210356248762697446949407573530086143415290314195533631308867097853951),[-3,0x5ac635d8aa3a93e7b3ebbd55769886bc651d06b0cc53b0f63bce3c3e27d2604b])
	// G = P256([0x6b17d1f2e12c4247f8bce6e563a440f277037d812deb33a0f4a13945d898c296,0x4fe342e2fe1a7f9b8ee7eb4a7c0f9e162bce33576b315ececbb6406837bf51f5])
	// 0x4001*G
	curve := elliptic.P256()
	x, y := curve.ScalarBaseMult([]byte{0x40, 0x1})
	u, _ := new(big.Int).SetString("72695894326801153147216665794252886088922051068954702128373097767796982305416", 10)
	v, _ := new(big.Int).SetString("115052900738981503025783750032562471330247412874189631362047844261308430634379", 10)
	if x.Cmp(u) != 0 {
		t.Errorf("P256.ScalarBaseMult() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("P256.ScalarBaseMult() gotY = %v, want %v", y, v)
	}
}

func TestLiftX(t *testing.T) {
	curve := SIEC255()
	type args struct {
		X *big.Int
	}
	tests := []struct {
		name  string
		args  args
		wantX *big.Int
		wantY *big.Int
	}{
		{"generator lifts correctly", args{curve.Gx}, curve.Gx, curve.Gy},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotX, gotY := curve.liftX(tt.args.X)
			if !reflect.DeepEqual(gotX, tt.wantX) {
				t.Errorf("LiftX() gotX = %v, want %v", gotX, tt.wantX)
			}
			if !reflect.DeepEqual(gotY, tt.wantY) {
				t.Errorf("LiftX() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}

func TestCompress(t *testing.T) {
	curve := SIEC255()
	var x1, y1, x2, y2 *big.Int
	x1, _ = new(big.Int).SetString("5", 10)
	y1, _ = new(big.Int).SetString("12", 10)
	x2, y2 = curve.Decompress(curve.Compress(x1, y1))
	if x2.Cmp(x1) != 0 || y2.Cmp(y1) != 0 {
		t.Errorf("(%v,%v) did not compress/decompress correctly, got (%v,%v)", x1, y1, x2, y2)
	}
	x1, _ = new(big.Int).SetString("6784692728748995825599862402855483522016546426567910438357042338075027826575", 10)
	y1, _ = new(big.Int).SetString("14982863109320699114866362806305859444453206692004135551371801829915686450358", 10)
	x2, y2 = curve.Decompress(curve.Compress(x1, y1))
	if x2.Cmp(x1) != 0 || y2.Cmp(y1) != 0 {
		t.Errorf("(%v,%v) did not compress/decompress correctly, got (%v,%v)", x1, y1, x2, y2)
	}
}

func TestCompressRandom(t *testing.T) {
	curve := SIEC255()
	for i := 0; i < 10; i++ {
		x1, y1 := curve.ScalarBaseMult(hash(i))
		x2, y2 := curve.Decompress(curve.Compress(x1, y1))
		if x2.Cmp(x1) != 0 || y2.Cmp(y1) != 0 {
			t.Errorf("(%v,%v) did not compress/decompress correctly, got (%v,%v)", x1, y1, x2, y2)
		}
	}
}

func TestScalarMult2Random(t *testing.T) {
	curve := SIEC255()
	for i := 0; i < 10; i++ {
		x, y := curve.ScalarBaseMult(hash(i))
		u, v := curve.scalarMult2(curve.Gx, curve.Gy, hash(i))
		if x.Cmp(u) != 0 || y.Cmp(v) != 0 {
			t.Errorf("wanted (%v,%v), got (%v,%v)", x, y, u, v)
		}
	}
}

func TestEndomorphism(t *testing.T) {
	curve := SIEC255()
	x, y := curve.Gx, curve.Gy
	u, _ := new(big.Int).SetString("28948022309329048855892746252183396359753225502720738378331610790540530405116", 10)
	v, _ := new(big.Int).SetString("28948022309329048855892746252183396360603931420023084536990047309120118726709", 10)
	x, y = curve.phi(x, y)
	if x.Cmp(u) != 0 {
		t.Errorf("curve.ScalarBaseMult() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("curve.ScalarBaseMult() gotY = %v, want %v", y, v)
	}
	u, _ = new(big.Int).SetString("850705917302346158658436518579588321600", 10)
	v, _ = new(big.Int).SetString("12", 10)
	x, y = curve.phi(x, y)
	if x.Cmp(u) != 0 {
		t.Errorf("curve.ScalarBaseMult() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("curve.ScalarBaseMult() gotY = %v, want %v", y, v)
	}
	u, _ = new(big.Int).SetString("5", 10)
	v, _ = new(big.Int).SetString("28948022309329048855892746252183396360603931420023084536990047309120118726709", 10)
	x, y = curve.phi(x, y)
	if x.Cmp(u) != 0 {
		t.Errorf("curve.ScalarBaseMult() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("curve.ScalarBaseMult() gotY = %v, want %v", y, v)
	}
	u, _ = new(big.Int).SetString("28948022309329048855892746252183396359753225502720738378331610790540530405116", 10)
	v, _ = new(big.Int).SetString("12", 10)
	x, y = curve.phi(x, y)
	if x.Cmp(u) != 0 {
		t.Errorf("curve.ScalarBaseMult() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("curve.ScalarBaseMult() gotY = %v, want %v", y, v)
	}
	u, _ = new(big.Int).SetString("850705917302346158658436518579588321600", 10)
	v, _ = new(big.Int).SetString("28948022309329048855892746252183396360603931420023084536990047309120118726709", 10)
	x, y = curve.phi(x, y)
	if x.Cmp(u) != 0 {
		t.Errorf("curve.ScalarBaseMult() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("curve.ScalarBaseMult() gotY = %v, want %v", y, v)
	}
	u, _ = new(big.Int).SetString("5", 10)
	v, _ = new(big.Int).SetString("12", 10)
	x, y = curve.phi(x, y)
	if x.Cmp(u) != 0 {
		t.Errorf("curve.ScalarBaseMult() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("curve.ScalarBaseMult() gotY = %v, want %v", y, v)
	}
}
