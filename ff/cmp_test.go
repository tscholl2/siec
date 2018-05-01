package ff

import "testing"

func Test_cmp(t *testing.T) {
	type args struct {
		a [4]uint64
		b [4]uint64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1=1", args{[4]uint64{1, 0, 0, 0}, [4]uint64{1, 0, 0, 0}}, 0},
		{"1<2", args{[4]uint64{1, 0, 0, 0}, [4]uint64{2, 0, 0, 0}}, -1},
		{"2>1", args{[4]uint64{2, 0, 0, 0}, [4]uint64{1, 0, 0, 0}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cmp(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("cmp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cmp_random(t *testing.T) {
	for i := 0; i < 100; i++ {
		x := randomElement(2 * i)
		y := randomElement(2*i + 1)
		A, B := ElementToBigInt(x), ElementToBigInt(y)
		got := cmp(x, y)
		want := A.Cmp(B)
		if got != want {
			t.Errorf("/%d add(%v,%v) = %v, want %v", i, x, y, got, want)
			t.FailNow()
		}
	}
}

func Benchmark_cmp(b *testing.B) {
	x := randomElement(1)
	y := randomElement(2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cmp(x, y)
	}
}

func Benchmark_cmp_BI(b *testing.B) {
	x := ElementToBigInt(randomElement(1))
	y := ElementToBigInt(randomElement(2))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.Cmp(y)
	}
}
