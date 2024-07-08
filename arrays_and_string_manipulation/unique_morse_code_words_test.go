package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_uniqueMorseRepresentations(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{words: []string{"ab"}},
			want: 1,
		}, {
			name: "",
			args: args{words: []string{"rwjje", "aittjje", "auyyn", "lqtktn", "lmjwn"}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, uniqueMorseRepresentations(tt.args.words), "uniqueMorseRepresentations(%v)", tt.args.words)
		})
	}
}
