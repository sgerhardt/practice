package arrays_and_string_manipulation

import "testing"

func Test_iterCols(t *testing.T) {
	tests := []struct {
		name string
	}{{name: "test"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterCols()
		})
	}
}

func Test_iterRows(t *testing.T) {
	tests := []struct {
		name string
	}{{name: "test"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterRows()
		})
	}
}

func Test_iterRowsWithRanges(t *testing.T) {
	tests := []struct {
		name string
	}{{name: "test"}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iterRowsWithRanges()
		})
	}
}
