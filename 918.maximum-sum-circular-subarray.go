/*
 * @lc app=leetcode.cn id=918 lang=golang
 * @lcpr version=30204
 *
 * [918] 环形子数组的最大和
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func maxSubarraySumCircular(nums []int) (answer int) {
	// length := len(nums)
	// newNums := make([]int, 0, length*2)
	// newNums = append(newNums, nums...)
	// newNums = append(newNums, nums...)
	// answer = nums[0]
	// for i := 0; i < length; i++ {
	// 	nums1 := newNums[i : i+length]
	// 	maxSubSum := 0
	// 	preSubSum := 0
	// 	for _, val := range nums1 {
	// 		maxSubSum = max(preSubSum+val, val)
	// 		preSubSum = maxSubSum
	// 		answer = max(answer, maxSubSum)
	// 	}
	// }
	//N 平方，时间超了
	maxSubSum := 0
	preSubSum := 0
	maxOrigin := nums[0]
	sum := 0
	for _, val := range nums {
		maxSubSum = max(preSubSum+val, val)
		preSubSum = maxSubSum
		maxOrigin = max(maxOrigin, maxSubSum)
		sum += val
	}
	minSubSum := 0
	preMinSum := 0
	maxRing := 0
	for _, val := range nums {
		minSubSum = min(preMinSum+val, val)
		preMinSum = minSubSum
		maxRing = min(maxRing, minSubSum)
	}
	if maxRing == sum {
		return maxOrigin
	}
	maxRing = sum - maxRing
	answer = max(maxOrigin, maxRing)
	return
}

// @lc code=end

/*
// @lcpr case=start
// [1,-2,3,-2]\n
// @lcpr case=end

// @lcpr case=start
// [5,-3,5]\n
// @lcpr case=end

// @lcpr case=start
// [3,-2,2,-3]\n
// @lcpr case=end

*/
