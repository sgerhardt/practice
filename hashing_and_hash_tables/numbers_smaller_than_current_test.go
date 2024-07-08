package hashing_and_hash_tables

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_smallerNumbersThanCurrent(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "basic",
			args: []int{8, 1, 2, 2, 3},
			want: []int{4, 0, 1, 1, 3},
		}, {
			name: "basic2",
			args: []int{0, 1, 0, 0, 4, 4},
			want: []int{0, 3, 0, 0, 4, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, smallerNumbersThanCurrent(tt.args))
		})
	}
}
