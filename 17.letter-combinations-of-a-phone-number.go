/*
 * @lc app=leetcode.cn id=17 lang=golang
 * @lcpr version=30204
 *
 * [17] 电话号码的字母组合
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func letterCombinations(digits string) (answer []string) {
	length := len(digits)
	if length == 0 {
		return
	}
	letterMap := [8][4]rune{{'a', 'b', 'c'}, {'d', 'e', 'f'}, {'g', 'h', 'i'}, {'j', 'k', 'l'}, {'m', 'n', 'o'}, {'p', 'q', 'r', 's'}, {'t', 'u', 'v'}, {'w', 'x', 'y', 'z'}}
	lineString := make([]rune, length)
	var dfs func(level int)
	// 在调用 dfs 之前
	capacity := 1
	for _, digit := range digits {
		num := digit - '2'
		if num == 5 || num == 7 { // '7' (pqrs) or '9' (wxyz)
			capacity *= 4
		} else {
			capacity *= 3
		}
	}
	answer = make([]string, 0, capacity) // 创建容量为 capacity 的切片
	//长度为 0
	//提前分配好空间，减少 append 开销
	dfs = func(level int) {
		if level == length {
			answer = append(answer, string(lineString))
			return
		}
		nowNum := digits[level] - '2'
		for _, v := range letterMap[nowNum] {
			if v != 0 {
				lineString[level] = v
				dfs(level + 1)
			}
		}
	}
	dfs(0)
	return
}

// @lc code=end

/*
// @lcpr case=start
// "23"\n
// @lcpr case=end

// @lcpr case=start
// ""\n
// @lcpr case=end

// @lcpr case=start
// "2"\n
// @lcpr case=end

*/
