package ff

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_double(t *testing.T) {
	type args struct {
		a Element
	}
	tests := []struct {
		name  string
		args  args
		wantB Element
	}{
		{"1", args{Element{1, 0, 0, 0}}, Element{2, 0, 0, 0}},
		{"2^64", args{Element{0, 1, 0, 0}}, Element{0, 2, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotB := double(tt.args.a); !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("double() = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func Test_double_random(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := randomElement(i)
		a[3] = a[3] & 0x7fffffffffffffff // clear last bit so garunteed it will fit in 256 bits
		got := double(a)
		want := BigIntToElement(new(big.Int).Lsh(ElementToBigInt(a), 1))
		if got != want {
			t.Errorf("/%d double(%v) = %v, want %v", i, a, got, want)
			t.FailNow()
		}
	}
}

func Benchmark_double(b *testing.B) {
	a := randomElement(1)
	a[3] = a[3] & 0x7fffffffffffffff
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		double(a)
	}
}

func Benchmark_double_BI(b *testing.B) {
	a := ElementToBigInt(randomElement(1))
	z := new(big.Int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		z.Lsh(a, 1)
	}
}

func Benchmark_double_add(b *testing.B) {
	a := randomElement(1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(a, a)
	}
}
