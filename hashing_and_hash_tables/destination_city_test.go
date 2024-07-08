package hashing_and_hash_tables

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_destCity(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{[][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}},
			want: "Sao Paulo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, destCity(tt.args.paths), "destCity(%v)")
		})
	}
}
