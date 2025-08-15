/*
 * @lc app=leetcode.cn id=224 lang=golang
 * @lcpr version=30204
 *
 * [224] 基本计算器
 */

// @lcpr-template-start
package leetcode

import (
	"strconv"
	"strings"
	"unicode"
)

// @lcpr-template-end
// @lc code=start
// 分两步，一步构成逆波兰表达式，一步计算逆波兰表达式
func calculate(s string) int {
	// 先走第一步，转化得到逆波兰表达式
	//使用 rune 类型无疑是错误的
	//当为符号时，和栈顶符号进行优先级比较，如果当前符号优先级大于栈顶优先级，那么将其加入符号栈，如果小于或等于，将符号栈栈顶符号出栈到结果栈，将当前符号入栈到符号栈
	//一坨屎，只有加减和括号
	resultStack := []string{}
	symbolStack := []string{}
	var numBuilder strings.Builder
	//建立逆波兰表达式结果栈和用于辅助的符号栈
	for i := 0; i < len(s); i++ {
		char := rune(s[i])
		//将当前字符转化为 rune，用于判断是否为数字
		if unicode.IsDigit(char) {
			numBuilder.Reset()
			for i < len(s) && unicode.IsDigit(rune(s[i])) {
				numBuilder.WriteByte(s[i])
				//如果是数字，那么开始构建数字
				//直到非数字为止
				i++
			}
			resultStack = append(resultStack, numBuilder.String())
			i--
			//将数字一直遍历到非数字为止
			//i--防止跳过符号，因为外层循环还会有一次 i++
		} else {
			switch char {
			case ' ':
				continue
			case '+', '-':
				if len(symbolStack) > 0 && (symbolStack[len(symbolStack)-1] == "+" || symbolStack[len(symbolStack)-1] == "-") {
					resultStack = append(resultStack, symbolStack[len(symbolStack)-1])
					symbolStack = symbolStack[:len(symbolStack)-1]
					//出栈
					symbolStack = append(symbolStack, string(char))
				} else {
					symbolStack = append(symbolStack, string(char))
				}
			case '(':

				symbolStack = append(symbolStack, string(char))
				for i+1 < len(s) && s[i+1] == ' ' {
					i++
				}
				//往 (-中间加一个 0，不用做标记
				if s[i+1] == '-' {
					resultStack = append(resultStack, "0")
				}
			case ')':
				if symbolStack[len(symbolStack)-1] != "(" {
					resultStack = append(resultStack, symbolStack[len(symbolStack)-1])
					symbolStack = symbolStack[:len(symbolStack)-1]
				}
				symbolStack = symbolStack[:len(symbolStack)-1]
				//当遇到右括号时，将到左括号之前所有符号出栈到结果栈，抛弃左括号

			}

		}
	}
	if len(symbolStack) != 0 {
		resultStack = append(resultStack, symbolStack[0])
	}

	rpnStack := []int{}
	for _, val := range resultStack {
		topOfStack := len(rpnStack) - 1
		switch val {
		case "+", "-":
			operand2 := rpnStack[topOfStack]
			rpnStack = rpnStack[:topOfStack]
			var operand1 int
			if len(rpnStack) == 0 {
				operand1 = 0
			} else {
				operand1 = rpnStack[topOfStack-1]
				rpnStack = rpnStack[:topOfStack-1]
			}
			var result int
			if val == "+" {
				result = operand2 + operand1
			} else {
				result = operand1 - operand2
			}
			rpnStack = append(rpnStack, result)
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
// "1 + 1"\n
// @lcpr case=end

// @lcpr case=start
// " 2-1 + 2 "\n
// @lcpr case=end

// @lcpr case=start
// "(1+(4+5+2)-3)+(6+8)"\n
// @lcpr case=end

*/
