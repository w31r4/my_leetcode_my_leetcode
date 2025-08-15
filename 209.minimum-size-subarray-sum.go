/*
 * @lc app=leetcode.cn id=209 lang=golang
 * @lcpr version=30204
 *
 * [209] 长度最小的子数组
 *
 * https://leetcode.cn/problems/minimum-size-subarray-sum/description/
 *
 * algorithms
 * Medium (47.03%)
 * Likes:    2438
 * Dislikes: 0
 * Total Accepted:    1M
 * Total Submissions: 2.1M
 * Testcase Example:  '7\n[2,3,1,2,4,3]'
 *
 * 给定一个含有 n 个正整数的数组和一个正整数 target。
 *
 * 找出该数组中满足其总和大于等于 target 的长度最小的 子数组 [numsl, numsl+1, ..., numsr-1, numsr]
 * ，并返回其长度。如果不存在符合条件的子数组，返回 0。
 *
 *
 *
 * 示例 1：
 *
 * 输入：target = 7, nums = [2,3,1,2,4,3]
 * 输出：2
 * 解释：子数组 [4,3] 是该条件下的长度最小的子数组。
 *
 *
 * 示例 2：
 *
 * 输入：target = 4, nums = [1,4,4]
 * 输出：1
 *
 *
 * 示例 3：
 *
 * 输入：target = 11, nums = [1,1,1,1,1,1,1,1]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= target <= 10^9
 * 1 <= nums.length <= 10^5
 * 1 <= nums[i] <= 10^4
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 如果你已经实现 O(n) 时间复杂度的解法，请尝试设计一个 O(n log(n)) 时间复杂度的解法。
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 好好读题非常重要!!!
func minSubArrayLen(target int, nums []int) int {
	// 遍历 nums 数组，维护一个哈希表和一个当前哈希表中数字总和 sum
	//当 sum+ 当前值<target 时，将当前值加入哈希表，如果 sum+ 当前值>target 时，去除最小数，直到去除当前最小数会导致 sum 小于 target 为止
	//读题!!!!
	//使找到总和大于等于 target 的长度的最小的子数组
	//子数组!!!
	//设置左右指针，当右指针指过的数字之和 sum 小于目标值时，不断向右扩张，每扫过一次给 ans+1
	//当右指针指过数字之和大于目标值时，将左指针的值减去，如果减去后不小于目标值，将 ans-1，左指针右移一位
	left, right, ans, minAns := 0, 0, 0, 0
	isFirstTime := true
	length := len(nums)
	sum := 0
	if nums[0] >= target {
		return 1
	}
	for ; right < length; right++ {
		sum += nums[right]
		ans++
		//右端点在没有出现 sum<target 的情况下持续右移
		//其实 ans=right-left+1，为实时窗口长度
		if sum < target {
			continue
		}

		for sum >= target {
			//当 sum>target 时
			if isFirstTime {
				minAns = ans
				isFirstTime = false
				//为 minAns 初次赋值
			}
			minAns = min(ans, minAns)
			//minAns 始终代表出现过的达到目标的最小窗口长度，该窗口数组的 sum>target
			sum -= nums[left]
			//将窗口左移一位：减去窗口左端点值，同时左端点 +1
			left++
			ans--
			//此时的窗口长度也要减一
		}
	}
	return minAns

}

// @lc code=end

/*
// @lcpr case=start
// 7\n[2,3,1,2,4,3]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[1,4,4]\n
// @lcpr case=end

// @lcpr case=start
// 11\n[1,1,1,1,1,1,1,1]\n
// @lcpr case=end

*/
