/*
 * @lc app=leetcode.cn id=34 lang=golang
 * @lcpr version=30204
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func searchRange(nums []int, target int) []int {
	// 首先找到位置，然后左右开找边界
	//左闭右开
	answer := []int{-1, -1}
	length := len(nums)
	right := length
	left := 0
	isFound := false
	mid := 0
	for left < right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			isFound = true
			break
		}
		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if !isFound {
		return answer
	}
	midNum := mid
	left = midNum
	right = length
	for left < right {
		mid = left + (right-left)/2
		if target == nums[mid] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	answer[1] = left - 1
	right = midNum - 1
	left = -1
	for left < right {
		mid = left + (right-left)/2 + 1
		//加一确保重心后置，让 (1+2)/2=2
		if target == nums[mid] {
			right = mid - 1
		} else {
			left = mid
		}
	}
	answer[0] = left + 1
	return answer
}

// @lc code=end

/*
// @lcpr case=start
// [5,7,7,8,8,10]\n8\n
// @lcpr case=end

// @lcpr case=start
// [5,7,7,8,8,10]\n6\n
// @lcpr case=end

// @lcpr case=start
// []\n0\n
// @lcpr case=end

*/
