/*
 * @lc app=leetcode.cn id=39 lang=golang
 * @lcpr version=30204
 *
 * [39] 组合总和
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func combinationSum(candidates []int, target int) (answer [][]int) {
	// 这次递归的边界条件是大于 target 或等于 target，大于的话直接返回，等于的话加入 answer 切片后返回
	answer = make([][]int, 0, 150)
	path := make([]int, 0, target/2)
	total := 0
	var dfs func(num, level int)
	length := len(candidates)
	dfs = func(num, level int) {
		if total >= target {
			if total == target {
				temp := make([]int, len(path))
				copy(temp, path)
				answer = append(answer, temp)
				return
			} else {
				return
			}
		}
		for i := num; i < length; i++ {
			v := candidates[i]
			path = append(path, v)
			total += v
			dfs(i, level+1)
			//永不复行
			//这样子就可以让下一层也可以访问当前数字的同时不能访问当前数字之前的数字，因为之前的数字已经访问完毕了可能性
			//i+1 不让访问自己
			//i 直接传入可以访问自己
			total -= v
			path = path[:level]
		}
		//排列没问题，现在是组合的回合了
		//组合要求访问过的不在访问
	}
	dfs(0, 0)
	return
}

// @lc code=end

/*
// @lcpr case=start
// [2,3,6,7]\n7\n
// @lcpr case=end

// @lcpr case=start
// [2,3,5]\n8\n
// @lcpr case=end

// @lcpr case=start
// [2]\n1\n
// @lcpr case=end

*/
