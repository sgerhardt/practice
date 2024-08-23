package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountingSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "basic case",
			args: args{arr: []int{1, 5, 5, 3, 2}},
			want: []int{1, 2, 3, 5, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, CountingSort(tt.args.arr), "CountingSort(%v)", tt.args.arr)
		})
	}
}
