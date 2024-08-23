package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{s: "MCMXCIV"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, romanToInt(tt.args.s), "romanToInt(%v)", tt.args.s)
		})
	}
}
