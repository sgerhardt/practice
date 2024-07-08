package math

import (
	"errors"
	"strconv"
	"strings"
)

type operator struct {
	operand           string
	operatorIndexHigh int
	operatorIndexLow  int
}

func simpleArithmetic(s string) (int, error) {
	vals := strings.Split(s, " ")
	var arithmeticNums []int
	var parens []string
	var operators []operator
	for idx, val := range vals {
		switch val {
		case "(":
			// push onto stack
			parens = append(parens, "(")
		case ")":
			// pop from stack
			if parens[len(parens)-1] == "(" {
				parens = append(parens, parens[:len(parens)-1]...)
			} else {
				return 0, errors.New("error: no matching opening bracket")
			}
		case "+":

			operators = append(operators, operator{
				operand:           "+",
				operatorIndexHigh: idx + 1,
				operatorIndexLow:  idx - 1,
			})
		case "-":
			operators = append(operators, operator{
				operand:           "-",
				operatorIndexHigh: idx + 1,
				operatorIndexLow:  idx - 1,
			})
		default:
			v, err := strconv.Atoi(val)
			if err != nil {
				return 0, err
			}
			arithmeticNums = append(arithmeticNums, v)
		}
	}

	sum := 0
	for _, num := range arithmeticNums {
		sum += num
	}
	return sum, nil
}
