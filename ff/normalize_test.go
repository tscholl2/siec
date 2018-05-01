package ff

import (
	"reflect"
	"testing"
)

func Test_normalize(t *testing.T) {
	type args struct {
		a Element
	}
	tests := []struct {
		name string
		args args
		want Element
	}{
		{"1", args{Element{1, 0, 0, 0}}, Element{1, 0, 0, 0}},
		{"0", args{Element{0, 0, 0, 0}}, Element{0, 0, 0, 0}},
		{"p", args{p}, Element{0, 0, 0, 0}},
		{"2p", args{Element{0x8008206104082, 0, 0x4002081, 0x8000000000000000}}, Element{0, 0, 0, 0}},
		{"2p+1", args{Element{0x8008206104083, 0, 0x4002081, 0x8000000000000000}}, Element{1, 0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := normalize(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("normalize() = %v, want %v", got, tt.want)
			}
		})
	}
}
