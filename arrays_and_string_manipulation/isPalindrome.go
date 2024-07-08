package arrays_and_string_manipulation

import "strings"

func isPalindrome(s string) bool {
	s = strip(s)
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	s = strings.ToLower(s)
	j := len(s) - 1
	for i := 0; i <= len(s)/2; i++ {
		if s[i] != s[j] {
			return false
		}
		j--
	}
	return true
}

func strip(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') {
			result.WriteByte(b)
		}
	}
	return result.String()
}
