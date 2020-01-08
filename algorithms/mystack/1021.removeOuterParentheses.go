package mystack

// removeOuterParentheses remove the outer parentheses of S
func removeOuterParentheses(S string) string {
	var res string
	stack := NewStack()
	for _, c := range S {
		if c == ')' {
			stack.Pop()
		}
		if !stack.IsEmpty() {
			res += string(c)
		}
		if c == '(' {
			stack.Push(c)
		}
	}
	return res
}
