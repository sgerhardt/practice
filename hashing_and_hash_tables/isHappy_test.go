package hashing_and_hash_tables

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "no cycle",
			args: args{n: 19},
			want: true,
		},
		{
			name: "has a cycle",
			args: args{n: 2},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isHappy(tt.args.n), "isHappy(%v)", tt.args.n)
		})
	}
}
