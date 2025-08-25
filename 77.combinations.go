/*
 * @lc app=leetcode.cn id=77 lang=golang
 * @lcpr version=30204
 *
 * [77] 组合
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func combinationsCount(n, k int) int {
	// C(n, k) == C(n, n-k)，我们计算较小的那一边
	if k > n/2 {
		k = n - k
	}
	if k == 0 {
		return 1
	}

	res := 1
	for i := 1; i <= k; i++ {
		res = res * (n - i + 1) / i
		//
	}
	return res
}
func combine(n int, k int) (answer [][]int) {
	// visited := make([]bool, n)
	//组合需要标记有哪些节点已经访问过了
	capacity := combinationsCount(n, k)
	answer = make([][]int, 0, capacity)
	path := make([]int, k)
	var dfs func(num, level int)
	dfs = func(num, level int) {
		//组合需要标记有哪些节点已经访问过了
		//怎么标记呢？
		//在 dfs 中缩小范围并且不影响外层递归
		//有时候我们可以通过回溯 visited 表来进行
		//这题使用 num 来慢慢拓展不可访问的数字范围
		if level == k {
			temp := make([]int, k)
			// 2. 将 path 当前的内容，完整地复制到 temp 中
			copy(temp, path)
			answer = append(answer, temp)
			return
		}

		for i := num; i <= n-(k-level-1); i++ {
			//最外层只需要遍历 n-k+1 个数字就行，最后一个 k-1 个不用遍历，因为剩余的根本无法填充满 k
			//随着层数增加，需要填充的 k 逐渐减少，减少量恰好和 level 一致
			//可以优化剪枝
			//当剩余数字不足以填满 k 时，放弃循环
			path[level] = i
			dfs(i+1, level+1)
		}
		return

	}
	dfs(1, 0)
	return
}

// @lc code=end

/*
// @lcpr case=start
// 4\n2\n
// @lcpr case=end

// @lcpr case=start
// 1\n1\n
// @lcpr case=end

*/
