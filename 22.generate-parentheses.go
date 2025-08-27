/*
 * @lc app=leetcode.cn id=22 lang=golang
 * @lcpr version=30204
 *
 * [22] 括号生成
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func generateParenthesis(n int) (answer []string) {
	length := 2 * n
	path := make([]rune, length)
	//括号的基本条件就是左括号数量大于等于右括号
	left, right := 0, 0
	brackets := [2]rune{'(', ')'}
	var dfs func(level int)
	dfs = func(level int) {
		if left > n {
			return
		}
		if level == length {
			if left == right {
				answer = append(answer, string(path))
				return
			} else {
				return
			}

		}
		if left < right {
			return
		}
		for index, v := range brackets {
			if index == 0 {
				left++
				path[level] = v
				dfs(level + 1)
				left--
			} else {
				right++
				path[level] = v
				dfs(level + 1)
				right--
			}
		}
	}
	dfs(0)
	return
}

// @lc code=end

/*
// @lcpr case=start
// 3\n
// @lcpr case=end

// @lcpr case=start
// 1\n
// @lcpr case=end

*/
