package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fewestCoins(t *testing.T) {
	type args struct {
		value int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{value: 51},
			want: []int{2, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, fewestCoins(tt.args.value), "fewestCoins(%v)", tt.args.value)
		})
	}
}
