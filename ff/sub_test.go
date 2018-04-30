package ff

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_sub(t *testing.T) {
	type args struct {
		a Element
		b Element
	}
	tests := []struct {
		name string
		args args
		want Element
	}{
		{"1 - 1", args{Element{1, 0, 0, 0}, Element{1, 0, 0, 0}}, Element{0, 0, 0, 0}},
		{"12 - 1", args{Element{12, 0, 0, 0}, Element{1, 0, 0, 0}}, Element{11, 0, 0, 0}},
		{"p - (p-1)", args{Element{1126179130581057, 9223372036854775808, 33558592, 4611686018427387904}, Element{1126179130581056, 9223372036854775808, 33558592, 4611686018427387904}}, Element{1, 0, 0, 0}},
		{"1 - p", args{Element{1, 0, 0, 0}, Element{1126179130581057, 9223372036854775808, 33558592, 4611686018427387904}}, Element{1126179130581056, 9223372036854775808, 33558592, 13835058055282163712}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sub(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subRandom(t *testing.T) {
	for i := 3; i < 100; i++ {
		a1 := normalize(randomElement(int64(2 * i)))
		a2 := normalize(randomElement(int64(2*i + 1)))
		B := new(big.Int).Sub(ToBigInt(a1), ToBigInt(a2))
		if args, got, want := []Element{a1, a2}, sub(a1, a2), FromBigInt(B); got != want {
			t.Errorf("/%d sub(%v) = %v, want %v", i, args, got, want)
			t.FailNow()
		}
	}
}

func Benchmark_sub(b *testing.B) {
	a1 := randomElement(1)
	a2 := randomElement(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sub(a1, a2)
	}
}

func Benchmark_subBI(b *testing.B) {
	a1 := ToBigInt(randomElement(1))
	a2 := ToBigInt(randomElement(2))
	c := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Sub(a1, a2)
	}
}
