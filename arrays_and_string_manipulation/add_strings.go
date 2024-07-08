package arrays_and_string_manipulation

import (
	"strconv"
	"strings"
)

//Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 as a string.
//
//You must solve the problem without using any built-in library for handling large integers (such as BigInteger).
//You must also not convert the inputs to integers directly.
//
//
//
//Example 1:
//
//Input: num1 = "11", num2 = "123"
//Output: "134"
//Example 2:
//
//Input: num1 = "456", num2 = "77"
//Output: "533"
//Example 3:
//
//Input: num1 = "0", num2 = "0"
//Output: "0"
//

func addStrings(num1 string, num2 string) string {

	num1Index := len(num1) - 1
	num2Index := len(num2) - 1

	carry := int32(0)
	var num1Ints []int32
	for _, c := range num1 {
		num1Ints = append(num1Ints, c)
	}

	var num2Ints []int32
	for _, c := range num2 {
		num2Ints = append(num2Ints, c)
	}

	sb := strings.Builder{}
	for num1Index >= 0 || num2Index >= 0 || carry == 1 {
		num1Val := int32(0)
		if num1Index >= 0 {
			// convert to integer by offsetting ascii value
			num1Val = num1Ints[num1Index] - '0'
			num1Index--
		}

		num2Val := int32(0)
		if num2Index >= 0 {
			// convert to integer by offsetting ascii value
			num2Val = num2Ints[num2Index] - '0'
			num2Index--
		}

		sum := num1Val + num2Val + carry
		carry = sum / 10
		val := sb.String()
		sb.Reset()

		// note, we write out a single digit at a time, which is why we mod 10 and use a carry variable...
		sb.WriteString(strconv.Itoa(int(sum)%10) + val)
	}
	return sb.String()
}
