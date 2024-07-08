package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getClicks(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "",
			args: args{input: []string{
				"5,yahoo.com",
				"3,sports.yahoo.com"}},
			want: map[string]string{"sports.yahoo.com": "3", "yahoo.com": "8", "com": "8"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getClicks(tt.args.input), "getClicks(%v)", tt.args.input)
		})
	}
}
