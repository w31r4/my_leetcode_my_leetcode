/*
 * @lc app=leetcode.cn id=15 lang=golang
 * @lcpr version=30204
 *
 * [15] 三数之和
 *
 * https://leetcode.cn/problems/3sum/description/
 *
 * algorithms
 * Medium (39.13%)
 * Likes:    7442
 * Dislikes: 0
 * Total Accepted:    2.3M
 * Total Submissions: 5.8M
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * 给你一个整数数组 nums，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j !=
 * k，同时还满足 nums[i] + nums[j] + nums[k] == 0。请你返回所有和为 0 且不重复的三元组。
 *
 * 注意：答案中不可以包含重复的三元组。
 *
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [-1,0,1,2,-1,-4]
 * 输出：[[-1,-1,2],[-1,0,1]]
 * 解释：
 * nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
 * nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
 * nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
 * 不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
 * 注意，输出的顺序和三元组的顺序并不重要。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [0,1,1]
 * 输出：[]
 * 解释：唯一可能的三元组和不为 0。
 *
 *
 * 示例 3：
 *
 * 输入：nums = [0,0,0]
 * 输出：[[0,0,0]]
 * 解释：唯一可能的三元组和为 0。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 3 <= nums.length <= 3000
 * -10^5 <= nums[i] <= 10^5
 *
 *
 */

// @lcpr-template-start
package leetcode

import "slices"

// @lcpr-template-end
// @lc code=start
// 思路：先排序后计算
// left right mid
// right 为最大，left 初始为 0,mid 初始为 1
// 当整体大于 0 时，right 左移，直到小于等于 0
// 整体小于 0 时，mid 先右移，left 再右移
// 要符合最小加减数原则，每一次移动都会让当前的三个数之和变化最小
func threeSum(nums []int) (ans [][]int) {
	// 	slices.Sort(nums)
	// 	n := len(nums)
	// 	for i, x := range nums[:n-2] {
	// 		if i > 0 && x == nums[i-1] { // 跳过重复数字
	// 			continue
	// 		}
	// 		if x+nums[i+1]+nums[i+2] > 0 { // 优化一
	// 			break
	// 		}
	// 		if x+nums[n-2]+nums[n-1] < 0 { // 优化二
	// 			continue
	// 		}
	// 		j, k := i+1, n-1
	// 		for j < k {
	// 			s := x + nums[j] + nums[k]
	// 			if s > 0 {
	// 				k--
	// 			} else if s < 0 {
	// 				j++
	// 			} else { // 三数之和为 0
	// 				ans = append(ans, []int{x, nums[j], nums[k]})
	// 				for j++; j < k && nums[j] == nums[j-1]; j++ {
	// 				} // 跳过重复数字
	// 				for k--; k > j && nums[k] == nums[k+1]; k-- {
	// 				} // 跳过重复数字
	// 			}
	// 		}
	// 	}
	// 	return
	// }

	left := 0
	slices.Sort(nums)
	// 思路：确定一个移动另外两个
	for ; left < len(nums)-2; left++ {
		if left > 0 {
			if nums[left-1] == nums[left] {
				continue
			}
		}
		// -4 -1 -1 0 1 2
		right := len(nums) - 1
		mid := left + 1
		if nums[left]+nums[right]+nums[right-1] < 0 {
			continue
		}
		if nums[left]+nums[mid]+nums[mid+1] > 0 {
			break
		}
		for mid < right {
			sum := nums[mid] + nums[right]

			if sum == -nums[left] {
				ans = append(ans, []int{nums[left], nums[mid], nums[right]})
				for mid++; mid < right && nums[mid] == nums[mid-1]; mid++ {
				} // 跳过重复数字
				for right--; right > mid && nums[right] == nums[right+1]; right-- {
				} // 跳过重复数字
				continue
			}

			if sum > -nums[left] {
				right--
				continue
			}
			if sum < -nums[left] {
				mid++
				continue
			}

			//狗屎，没有将 right 重置导致逻辑错误
		}
	}
	return
}

// 	//-1 -1 0 1 2 -4
// }
// for right > mid && mid > left {
// 	sum := nums[right] + nums[mid] + nums[left]
// 	if sum == 0 {
// 		ans = append(ans, []int{nums[right], nums[mid], nums[left]})
// 		right--
// 		continue
// 	}
// 	if sum > 0 {
// 		right--
// 		continue
// 	}
// 	if sum < 0 {
// 		if mid-left > 1 {
// 			if nums[left+1]-nums[left] < nums[mid+1]-nums[mid] {
// 				left++
// 			}
// 		} else {
// 			mid++
// 		}
// 	}

// }

// @lc code=end

/*
// @lcpr case=start
// [-10,-5,-5,-4,-4,-3,-2,-2,0,0,1,2,2,2,2,5,5]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0]\n
// @lcpr case=end

*/
