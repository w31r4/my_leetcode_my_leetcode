/*
 * @lc app=leetcode.cn id=162 lang=golang
 * @lcpr version=30204
 *
 * [162] 寻找峰值
 */

// @lcpr-template-start
package leetcode

import "math"

// @lcpr-template-end
// @lc code=start
// 数字的排列除了目标状态有三种情况：
// 谷底型：两边都比中间高，查看两边是否山峰
// 斜坡型：一边高一边低，查看高点是否山峰，低点直接跳过到另一个邻接点
// 没有平坦
// 如果我们从头开始，那么一定是上坡型
// 且因为两边是负无穷，就算一直增长，在最后也是会有下降出现
// 然后我们从中间开始，观察中间和中间 +1 的增长情况
// 如果中间 +1 相较于中间增长了，那么右边一定存在峰顶
// 如果中间减一相当于中间增长了，那么左边一定存在峰顶
// 如果中间最高那再好不过
func findPeakElement(nums []int) int {
	start := 0
	length := len(nums)
	end := length - 1
	mid := start + (end-start)/2
	midNum := nums[mid]
	leftNum := 0
	if mid-1 == -1 {
		leftNum = math.MinInt
	} else {
		leftNum = nums[mid-1]
	}

	rightNum := 0
	if mid+1 == length {
		rightNum = math.MinInt
	} else {
		rightNum = nums[mid+1]
	}
	for midNum < rightNum || midNum < leftNum {
		if midNum < rightNum {
			start = mid + 1
			mid = start + (end-start)/2
			midNum = nums[mid]
			if mid-1 == -1 {
				leftNum = math.MinInt
			} else {
				leftNum = nums[mid-1]
			}
			if mid+1 == length {
				rightNum = math.MinInt
			} else {
				rightNum = nums[mid+1]
			}
		} else {
			end = mid - 1
			mid = start + (end-start)/2
			midNum = nums[mid]
			if mid-1 == -1 {
				leftNum = math.MinInt
			} else {
				leftNum = nums[mid-1]
			}
			if mid+1 == length {
				rightNum = math.MinInt
			} else {
				rightNum = nums[mid+1]
			}
		}
	}
	// isFound := false
	// var dfs func(int)
	// answer := 0
	// dfs = func(mid int) {
	// 	if isFound {
	// 		return
	// 	}
	// 	if mid >= length || mid < 0 {
	// 		return
	// 	}
	// 	midNum := nums[mid]
	// 	left := 0
	// 	if mid-1 >= 0 {
	// 		left = nums[mid-1]
	// 	} else {
	// 		left = math.MinInt
	// 	}
	// 	right := 0
	// 	if mid+1 < length {
	// 		right = nums[mid+1]
	// 	} else {
	// 		right = math.MinInt
	// 	}
	// 	if midNum > left && midNum > right {
	// 		isFound = true
	// 		answer = mid
	// 		return
	// 	}
	// 	if left > midNum && right > midNum {
	// 		dfs(mid + 1)
	// 		dfs(mid - 1)
	// 		return
	// 	}
	// 	if midNum > left {
	// 		dfs(mid - 2)
	// 		dfs(mid + 1)
	// 	} else {
	// 		dfs(mid - 1)
	// 		dfs(mid + 2)
	// 	}

	// }
	// dfs(mid)
	//愚蠢的解法
	return mid
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,1,3,5,6,4]\n
// @lcpr case=end

*/
