package main

import (
	"math/big"
	"reflect"
	"testing"
)

func BenchmarkNextSiec(b *testing.B) {
	M := new(big.Int)
	M.SetString("100000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000", 16)
	b.ResetTimer()
	for i := 1; i < b.N; i++ {
		nextSiec(M)
	}
}

func Test_nextSiec(t *testing.T) {
	type args struct {
		M *big.Int
	}
	tests := []struct {
		name  string
		args  args
		wantT *big.Int
		wantQ *big.Int
	}{
		{
			name:  "10^10",
			args:  args{M: big.NewInt(10000000000)},
			wantT: big.NewInt(200001),
			wantQ: big.NewInt(10000100003),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotT, gotQ := nextSiec(tt.args.M)
			if !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("nextSiec() gotT = %v, want %v", gotT, tt.wantT)
			}
			if !reflect.DeepEqual(gotQ, tt.wantQ) {
				t.Errorf("nextSiec() gotQ = %v, want %v", gotQ, tt.wantQ)
			}
		})
	}
}
