/*
 * @lc app=leetcode.cn id=20 lang=golang
 * @lcpr version=30204
 *
 * [20] 有效的括号
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 用栈模拟，输入端为{([时栈中将}]) 入栈，遇到}]) 时如果栈顶何其相同，那么出栈如果不同那么 return false
func isValid(s string) bool {
	stack := []rune{}
	// for _, val := range s {
	// 	switch val {
	// 	case '[':
	// 		stack = append(stack, ']')
	// 	case '(':
	// 		stack = append(stack, ')')
	// 	case '{':
	// 		stack = append(stack, '}')
	// 	}
	// 	if val == ')' || val == '}' || val == ']' {
	// 		if len(stack) > 0 && stack[len(stack)-1] == val {
	// 			stack = stack[:len(stack)-1]
	// 		} else {
	// 			return false
	// 		}
	// 	}
	// }
	//可以使用一个大 switch
	for _, char := range s {
		switch char {
		case '(', '[', '{': // 开括号：压入对应的闭括号
			if char == '(' {
				stack = append(stack, ')')
			} else if char == '[' {
				stack = append(stack, ']')
			} else { // char == '{'
				stack = append(stack, '}')
			}
		case ')', ']', '}': // 闭括号
			if len(stack) == 0 || stack[len(stack)-1] != char {
				return false // 栈空或栈顶不匹配
			}
			stack = stack[:len(stack)-1] // 出栈
			// default:
			// 如果题目保证字符串只包含括号，则不需要 default。
			// 如果可能包含其他字符且应视为无效，则可以在此 return false。
		}
	}
	//都一样
	return len(stack) == 0
}

// @lc code=end

/*
// @lcpr case=start
// "()"\n
// @lcpr case=end

// @lcpr case=start
// "()[]{}"\n
// @lcpr case=end

// @lcpr case=start
// "(]"\n
// @lcpr case=end

// @lcpr case=start
// "([])"\n
// @lcpr case=end

*/
