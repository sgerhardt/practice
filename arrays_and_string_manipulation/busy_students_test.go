package arrays_and_string_manipulation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_busyStudent(t *testing.T) {
	type args struct {
		startTime []int
		endTime   []int
		queryTime int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				startTime: []int{1, 2, 3},
				endTime:   []int{3, 2, 7},
				queryTime: 4,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, busyStudent(tt.args.startTime, tt.args.endTime, tt.args.queryTime), "busyStudent(%v, %v, %v)", tt.args.startTime, tt.args.endTime, tt.args.queryTime)
		})
	}
}
