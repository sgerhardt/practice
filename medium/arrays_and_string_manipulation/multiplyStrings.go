package arrays_and_string_manipulation

import "fmt"

// Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.
//
// Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.
//
// Example 1:
//
// Input: num1 = "2", num2 = "3"
// Output: "6"
// Example 2:
//
// Input: num1 = "123", num2 = "456"
// Output: "56088"
// Remember how we do multiplication?
//
// Start from right to left, perform multiplication on every pair of digits, and add them together. Let's draw the process! From the following draft, we can immediately conclude:
//
// `num1[i] * num2[j]` will be placed at indices `[i + j`, `i + j + 1]`
func multiply(num1 string, num2 string) string {
	m := len(num1)
	n := len(num2)
	pos := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := (num1[i] - '0') * (num2[j] - '0')
			p1 := i + j
			p2 := i + j + 1
			sum := int(mul) + pos[p2]

			pos[p1] += sum / 10 // tens position
			pos[p2] = sum % 10  // ones position
		}
	}

	strVal := ""
	for i := 0; i < len(pos); i++ {
		if pos[i] == 0 && strVal == "" {
			continue
		}
		strVal += fmt.Sprint(pos[i])
	}
	if strVal == "" {
		return "0"
	}
	return strVal
}
