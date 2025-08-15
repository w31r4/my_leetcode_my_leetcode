/*
 * @lc app=leetcode.cn id=55 lang=golang
 * @lcpr version=30204
 *
 * [55] 跳跃游戏
 *
 * https://leetcode.cn/problems/jump-game/description/
 *
 * algorithms
 * Medium (43.66%)
 * Likes:    3033
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 2.8M
 * Testcase Example:  '[2,3,1,1,4]'
 *
 * 给你一个非负整数数组 nums，你最初位于数组的 第一个下标。数组中的每个元素代表你在该位置可以跳跃的最大长度。
 *
 * 判断你是否能够到达最后一个下标，如果可以，返回 true；否则，返回 false。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [2,3,1,1,4]
 * 输出：true
 * 解释：可以先跳 1 步，从下标 0 到达下标 1, 然后再从下标 1 跳 3 步到达最后一个下标。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [3,2,1,0,4]
 * 输出：false
 * 解释：无论怎样，总会到达下标为 3 的位置。但该下标的最大跳跃长度是 0，所以永远不可能到达最后一个下标。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^4
 * 0 <= nums[i] <= 10^5
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func canJump(nums []int) bool {
	// locate := nums[0]
	// // contrl := 0
	// if nums[0] == 0 && len(nums) > 1 {
	// 	return false
	// }
	// for i, _ := range nums {
	// 	// for nums[locate] == 0 {
	// 	// 	locate--
	// 	// 	if nums[locate] > 0 {
	// 	// 		locate += nums[locate]
	// 	// 		if nums[locate]
	// 	// 	}
	// 	// }
	// 	if locate >= len(nums)-1 {
	// 		return true
	// 	}
	// 	if i == locate {
	// 		for nums[locate+1] > nums[locate] && nums[locate] != 0 {
	// 			locate++
	// 			if locate >= len(nums)-1 {
	// 				return true
	// 			}
	// 		}
	// 		locate = nums[locate] + locate
	// 		if locate >= len(nums)-1 {
	// 			return true
	// 		}
	// 		for (i > 0 && nums[i-1] > nums[i]) || (i > 0 && nums[i] == 0) {
	// 			i--
	// 		}
	// 		i += nums[i]
	// 		locate = max(i, locate)
	// 	}
	// 	if locate >= len(nums)-1 {
	// 		return true
	// 	}
	//思路：左右嗅探法哪边大去哪，但是有致命错误，如
	// 911 910 121 00，可能会跳过多个递增递减区间导致错过最大值
	// 思路二，扩容数组，在有效范围内找出最大可扩容值，扩容后减去已经查找过的范围，在新扩容出的数组中继续查找最大可扩容值，直到数组大小越界或者扩容后的数组大小保持不变
	// length := nums[0] + 1 //数组长度
	// board := 0            //底边界，每次遍历刷新
	// maxSize := 0
	// for length < len(nums) { //容量大小比较
	// 	for index, val := range nums[board:length] {
	// 		if board+index+val+1 > maxSize {
	// 			maxSize = board + index + val + 1
	// 		} //获取新容量大小
	// 	}
	// 	board = length
	// 	if length == maxSize {
	// 		return false
	// 	}
	// 	length = maxSize
	// 	//获取新最大容量大小，如果不变则返回 false
	// 	if length >= len(nums) {
	// 		return true
	// 	}
	//思路是对的，但是想复杂了，没必要一段一段扩容，直接全部遍历，当最大范围无法触及当前遍历到的范围时，返回 false，当最大范围可以触及当前遍历到的范围时，将新旧范围作比较，看看哪个更大，然后更新
	mx := 0
	for index, val := range nums {
		if index > mx {
			return false
		}
		mx = max(mx, index+val)

	}

	return true
}

// func maxInt(slice []int) int {
// 	if len(slice) == 0 {
// 		return 0
// 	}
// 	max := slice[0]
// 	for _, v := range slice {
// 		if v > max {
// 			max = v
// 		}
// 	}
// 	return max
// }

// @lc code=end

/*
// @lcpr case=start
// [2,3,1,1,4]\n
// @lcpr case=end

// @lcpr case=start
// [3,2,1,0,4]\n
// @lcpr case=end

*/
