package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_arrayPairSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{nums: []int{1, 4, 3, 2}},
			want: 4,
		}, {
			name: "",
			args: args{nums: []int{6, 2, 6, 5, 1, 2}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, arrayPairSum(tt.args.nums), "arrayPairSum(%v)", tt.args.nums)
		})
	}
}
