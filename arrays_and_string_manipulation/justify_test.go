package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "",
			args: args{
				words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
				maxWidth: 16,
			},
			want: []string{"This    is    an", "example  of text", "justification.  "},
		},
		{
			name: "",
			args: args{
				words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
				maxWidth: 16,
			},
			want: []string{"What   must   be", "acknowledgment  ", "shall be        "},
		},
		{
			name: "",
			args: args{
				words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
				maxWidth: 20,
			},
			want: []string{"Science  is  what we", "understand      well", "enough to explain to", "a  computer.  Art is", "everything  else  we", "do                  "},
		},
		//
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, fullJustify(tt.args.words, tt.args.maxWidth), "fullJustify(%v, %v)", tt.args.words, tt.args.maxWidth)
		})
	}
}
