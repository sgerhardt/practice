package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_wordSeek(t *testing.T) {
	type args struct {
		puzzle [][]string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "simples horizontal case",
			want: []string{"apple"},
			args: args{puzzle: [][]string{{`b`, `a`, `p`, `p`, `l`, `e`}}},
		},
		{
			name: "simples vertical case",
			want: []string{"apple"},
			args: args{puzzle: [][]string{
				{`a`, `a`, `b`, `p`, `x`, `e`},
				{`b`, `p`, `b`, `p`, `y`, `e`},
				{`b`, `p`, `b`, `p`, `z`, `e`},
				{`b`, `l`, `b`, `p`, `j`, `e`},
				{`b`, `e`, `b`, `p`, `m`, `e`},
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, wordSeek(tt.args.puzzle))
		})
	}
}

func Test_wordSeekFromFile(t *testing.T) {
	tests := []struct {
		name string
		want []string
	}{
		{
			name: "basic horizontal",
			want: []string{"cat"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, wordSeekFromFile(), "wordSeekFromFile()")
		})
	}
}
