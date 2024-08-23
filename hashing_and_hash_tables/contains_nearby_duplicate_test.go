package hashing_and_hash_tables

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 2, 3, 1},
				k:    3,
			},
			want: true,
		}, {
			name: "",
			args: args{
				nums: []int{1, 0, 1, 1},
				k:    1,
			},
			want: true,
		}, {
			name: "",
			args: args{
				nums: []int{1, 2, 3, 1, 1, 2, 3},
				k:    0,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, containsNearbyDuplicate(tt.args.nums, tt.args.k), "containsNearbyDuplicate(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
