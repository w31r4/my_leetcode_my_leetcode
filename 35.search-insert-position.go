/*
 * @lc app=leetcode.cn id=35 lang=golang
 * @lcpr version=30204
 *
 * [35] 搜索插入位置
 */
package leetcode

// @lc code=start
func searchInsert(nums []int, target int) int {
	//递归要的空间大
	//循环要的空间小
	length := len(nums)
	left := 0
	right := length
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}

// @lc code=end

/*
// @lcpr case=start
// [1,3,5,6]\n5\n
// @lcpr case=end

// @lcpr case=start
// [1,3,5,6]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,3,5,6]\n7\n
// @lcpr case=end

*/
