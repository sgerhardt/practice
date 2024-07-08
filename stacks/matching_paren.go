package stacks

// The trick to many "parsing" questions like this is using a stack to track which brackets/phrases/etc are "open" as you go.
//
// So next time you get a parsing question, one of your first thoughts should be "use a stack!"
//
// In this problem, we can realize our stack would only hold '(' characters. So instead of storing each of those characters in a stack, we can store the number of items our stack would be holding.
//
// That gets us from O(n)O(n) space to O(1)O(1) space.
//
// It's pretty cool when you can replace a whole data structure with a single integer :)
func matchingParen(input string, pos int) int {
	openNestedParens := 0
	for idx := pos; idx < len(input); idx++ {
		if input[idx] == '(' {
			openNestedParens++
		} else if input[idx] == ')' {
			openNestedParens--
			if openNestedParens == 0 {
				return idx
			}
		}
	}
	return 0
}
