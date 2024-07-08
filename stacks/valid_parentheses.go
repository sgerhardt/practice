package stacks

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
//
// Example 1:
//
// Input: s = "()"
// Output: true
// Example 2:
//
// Input: s = "()[]{}"
// Output: true
// Example 3:
//
// Input: s = "(]"
// Output: false
func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}

	var parens []string

	for _, char := range s {
		c := string(char)
		if c == `(` {
			parens = append(parens, c)
			continue
		} else if c == `{` {
			parens = append(parens, c)
			continue
		} else if c == `[` {
			parens = append(parens, c)
			continue
		}

		if len(parens) == 0 {
			// if parens is 0, and we're still processing characters, we must have a closing bracket without an open...
			return false
		}
		// peek last item to see if it's open bracket
		lastChar := parens[len(parens)-1:][0]

		if lastChar == `(` && c == `)` || lastChar == `{` && c == `}` || lastChar == `[` && c == `]` {
			// pop it off the stack
			parens = parens[:len(parens)-1]
		} else {
			// could not find matching paren...
			return false
		}

	}
	return len(parens) == 0
}
