/*
 * @lc app=leetcode.cn id=238 lang=golang
 * @lcpr version=30204
 *
 * [238] 除自身以外数组的乘积
 *
 * https://leetcode.cn/problems/product-of-array-except-self/description/
 *
 * algorithms
 * Medium (77.31%)
 * Likes:    2005
 * Dislikes: 0
 * Total Accepted:    664.2K
 * Total Submissions: 858.1K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。
 *
 * 题目数据 保证 数组 nums 之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。
 *
 * 请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
 *
 *
 *
 * 示例 1:
 *
 * 输入：nums = [1,2,3,4]
 * 输出：[24,12,8,6]
 *
 *
 * 示例 2:
 *
 * 输入：nums = [-1,1,0,-3,3]
 * 输出：[0,0,9,0,0]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= nums.length <= 10^5
 * -30 <= nums[i] <= 30
 * 输入 保证 数组 answer[i] 在  32 位 整数范围内
 *
 *
 *
 *
 * 进阶：你可以在 O(1) 的额外空间复杂度内完成这个题目吗？（出于对空间复杂度分析的目的，输出数组 不被视为 额外空间。）
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func productExceptSelf(nums []int) []int {
	pre := make([]int, len(nums))
	pre[0] = 1
	for i, _ := range nums {
		if i > 0 {
			pre[i] = pre[i-1] * nums[i-1]
			//计算不包含自身的前缀乘积
		}
	}
	maxLimit := len(nums) - 1
	for i := maxLimit; i >= 0; i-- {
		if i == maxLimit {
			continue
		}
		nums[i] *= nums[i+1]
		//获取从尾到头的包含自身的乘积
		//加个 1 就不包含自身了
	}
	for i := 0; i < maxLimit; i++ {
		pre[i] = nums[i+1] * pre[i]
		//将两个乘积相乘获得目标乘积
	}
	return pre
	// 三个 On 循环，O1 空间
	//这叫前后缀分解
	//可以进行优化
	//思路是先算不包含自身的后缀，当算前缀时顺便算出结果，不需要开辟额外的数组空间，用一个前缀值累乘就行了
	// n := len(nums)
	// suf := make([]int, n)
	// suf[n-1] = 1
	// for i := n - 2; i >= 0; i-- {
	// 	suf[i] = nums[i+1] * suf[i+1]
	// 	//不包含自身的后缀乘积
	// }
	// pre := 1
	// for i, val := range nums {
	// 	// 此时 pre 为 nums[0] 到 nums[i-1] 的乘积，直接乘到 suf[i] 中

	// 	suf[i] *= pre
	// 	pre *= val
	// }
	// return suf
	// // 三个 On 循环，O1 空间
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [-1,1,0,-3,3]\n
// @lcpr case=end

*/
