package stacks

// You're working with an intern that keeps coming to you with JavaScript code that won't run because the braces, brackets, and parentheses are off.
// To save you both some time, you decide to write a braces/brackets/parentheses validator.
//
// Let's say:
//
// '(', '{', '[' are called "openers."
// ')', '}', ']' are called "closers."
// Write an efficient method that tells us whether or not an input string's openers and closers are properly nested.
//
// Examples:
//
// "{ [ ] ( ) }" should return true
// "{ [ ( ] ) }" should return false
// "{ [ }" should return false
func bracketsValid(chars []string) bool {
	bracketStack := StackString{}
	for _, c := range chars {
		switch c {
		case `(`:
			bracketStack.push(`(`)
		case `{`:
			bracketStack.push(`{`)
		case `[`:
			bracketStack.push(`[`)

		case `)`:
			poppedElement := bracketStack.pop()
			if poppedElement != `(` {
				return false
			}
		case `}`:
			poppedElement := bracketStack.pop()
			if poppedElement != `{` {
				return false
			}
		case `]`:
			poppedElement := bracketStack.pop()
			if poppedElement != `[` {
				return false
			}
		}
	}
	return true
}
