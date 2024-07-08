package arrays_and_string_manipulation

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	weight := 0
	h := fmt.Sprintf("%08b", num)
	for _, c := range h {
		if string(c) == "1" {
			weight++
		}
	}
	return weight
}
