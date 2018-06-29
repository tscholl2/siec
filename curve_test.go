package siec

import (
	"crypto/elliptic"
	"crypto/rand"
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
		t.Errorf("SIEC255Params.Double() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("SIEC255Params.Double() gotY = %v, want %v", y, v)
	}
	x, y = curve.Double(x, y)
	u, _ = new(big.Int).SetString("12318642006867402687195826566147291859634823582672295191656499276835526033145", 10)
	v, _ = new(big.Int).SetString("9343467693237486709905252998911952863134805995110526737200728195882424275543", 10)
	if x.Cmp(u) != 0 {
		t.Errorf("SIEC255Params.Double() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("SIEC255Params.Double() gotY = %v, want %v", y, v)
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
		t.Errorf("SIEC255Params.Add() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("SIEC255Params.Add() gotY = %v, want %v", y, v)
	}
}

func TestScale(t *testing.T) {
	curve := SIEC255()
	x, y := curve.ScalarBaseMult([]byte{0x40})
	u, _ := new(big.Int).SetString("22784956368772284587014129354783824301122808969944262244690262244567645543628", 10)
	v, _ := new(big.Int).SetString("15122902039027963115899162435007755480175091927763482548168795824636596047001", 10)
	if x.Cmp(u) != 0 {
		t.Errorf("SIEC255Params.ScalarBaseMult() gotX = %v, want %v", x, u)
	}
	if y.Cmp(v) != 0 {
		t.Errorf("SIEC255Params.ScalarBaseMult() gotY = %v, want %v", y, v)
	}
}

func TestScaleP256(t *testing.T) {
	// Sage:
	// P256 = EllipticCurve(GF(115792089210356248762697446949407573530086143415290314195533631308867097853951),[-3,0x5ac635d8aa3a93e7b3ebbd55769886bc651d06b0cc53b0f63bce3c3e27d2604b])
	// G = P256([0x6b17d1f2e12c4247f8bce6e563a440f277037d812deb33a0f4a13945d898c296,0x4fe342e2fe1a7f9b8ee7eb4a7c0f9e162bce33576b315ececbb6406837bf51f5])
	// 0x40*G
	curve := elliptic.P256()
	x, y := curve.ScalarBaseMult([]byte{0x40})
	u, _ := new(big.Int).SetString("4534198767316794591643245143622298809742628679895448054572722918996032022405", 10)
	v, _ := new(big.Int).SetString("38538856030597174617352966265796180312895426960288118979288294421866280361154", 10)
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
			gotX, gotY := LiftX(tt.args.X)
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
	var x1, y1, x2, y2 *big.Int
	x1, _ = new(big.Int).SetString("5", 10)
	y1, _ = new(big.Int).SetString("12", 10)
	x2, y2 = Decompress(Compress(x1, y1))
	if x2.Cmp(x1) != 0 || y2.Cmp(y1) != 0 {
		t.Errorf("(%v,%v) did not compress/decompress correctly, got (%v,%v)", x1, y1, x2, y2)
	}
	x1, _ = new(big.Int).SetString("6784692728748995825599862402855483522016546426567910438357042338075027826575", 10)
	y1, _ = new(big.Int).SetString("14982863109320699114866362806305859444453206692004135551371801829915686450358", 10)
	x2, y2 = Decompress(Compress(x1, y1))
	if x2.Cmp(x1) != 0 || y2.Cmp(y1) != 0 {
		t.Errorf("(%v,%v) did not compress/decompress correctly, got (%v,%v)", x1, y1, x2, y2)
	}
}

func TestCompressRandom(t *testing.T) {
	curve := SIEC255()
	b := make([]byte, 32)
	for i := 0; i < 10; i++ {
		rand.Read(b)
		x1, y1 := curve.ScalarBaseMult(b)
		x2, y2 := Decompress(Compress(x1, y1))
		if x2.Cmp(x1) != 0 || y2.Cmp(y1) != 0 {
			t.Errorf("(%v,%v) did not compress/decompress correctly, got (%v,%v)", x1, y1, x2, y2)
		}
	}
}
