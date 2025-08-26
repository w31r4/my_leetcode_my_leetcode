/*
 * @lc app=leetcode.cn id=46 lang=golang
 * @lcpr version=30204
 *
 * [46] 全排列
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func getCap(n int) int {
	cap := 1
	for i := 1; i <= n; i++ {
		cap *= i
	}
	return cap
}
func permute(nums []int) (answer [][]int) {
	length := len(nums)
	cap := getCap(length)
	//阶乘分配空间量
	answer = make([][]int, 0, cap)
	//提前分配好空间
	path := make([]int, length)
	visited := make([]bool, length)
	var dfs func(level int)
	dfs = func(level int) {
		if level == length {
			temp := make([]int, length)
			copy(temp, path)
			answer = append(answer, temp)
			return
		}
		for i, v := range nums {
			if !visited[i] {
				path[level] = v
				visited[i] = true
				dfs(level + 1)
				visited[i] = false
				//回溯 visited，为下一轮循环做准备
			}
		}
	}
	dfs(0)
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
