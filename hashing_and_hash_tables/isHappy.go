package hashing_and_hash_tables

import (
	"strconv"
	"strings"
)

func isHappy(n int) bool {
	digits := strconv.Itoa(n)

	seenNumbers := map[int]struct{}{}
	for v := squareDigits(strings.Split(digits, "")); v != 1; v = squareDigits(strings.Split(digits, "")) {
		if _, ok := seenNumbers[v]; ok {
			return false
		}
		seenNumbers[v] = struct{}{}
		digits = strconv.Itoa(v)
	}

	return true
}

func squareDigits(digits []string) int {
	sumOfSquaredDigits := 0
	for _, d := range digits {
		val, _ := strconv.Atoi(d)
		square := val * val
		sumOfSquaredDigits += square
	}
	return sumOfSquaredDigits
}
