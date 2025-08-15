/*
 * @lc app=leetcode.cn id=150 lang=golang
 * @lcpr version=30204
 *
 * [150] 逆波兰表达式求值
 */

// @lcpr-template-start
package leetcode

import "strconv"

// @lcpr-template-end
// @lc code=start
func evalRPN(tokens []string) int {
	rpnStack := []int{}
	for _, val := range tokens {
		topOfStack := len(rpnStack) - 1
		switch val {
		case "-":
			newNum := rpnStack[topOfStack-1] - rpnStack[topOfStack]
			rpnStack = rpnStack[:topOfStack-1]
			rpnStack = append(rpnStack, newNum)

		case "+":
			newNum := rpnStack[topOfStack-1] + rpnStack[topOfStack]
			rpnStack = rpnStack[:topOfStack-1]
			rpnStack = append(rpnStack, newNum)

		case "/":
			newNum := rpnStack[topOfStack-1] / rpnStack[topOfStack]
			rpnStack = rpnStack[:topOfStack-1]
			rpnStack = append(rpnStack, newNum)

		case "*":
			newNum := rpnStack[topOfStack-1] * rpnStack[topOfStack]
			rpnStack = rpnStack[:topOfStack-1]
			rpnStack = append(rpnStack, newNum)
		default:
			newNum, _ := strconv.Atoi(val)
			rpnStack = append(rpnStack, newNum)
		}

	}
	return rpnStack[0]
}

// @lc code=end

/*
// @lcpr case=start
// ["2","1","+","3","*"]\n
// @lcpr case=end

// @lcpr case=start
// ["4","13","5","/","+"]\n
// @lcpr case=end

// @lcpr case=start
// ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]\n
// @lcpr case=end

*/
