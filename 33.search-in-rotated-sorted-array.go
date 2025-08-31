/*
 * @lc app=leetcode.cn id=33 lang=golang
 * @lcpr version=30204
 *
 * [33] 搜索旋转排序数组
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 4,5,6,1,2,3
// 7,1,2,3,4,5,7,1,2,3,4,5
// 二分算法的关键是确定要找的数字在哪边，且该数字固定比当前数字大或者小
// 可以确定一个虚拟数字，如果当前数字在起始数字和目标数字大小之间，那么左边界同时更新起始数字移动，如果不在，那么右边界移动
// 扩容法在扩容时时间复杂度是 O(N)
// 我们继续观察旋转数组，它的从头开始直到旋转点一直递增，从尾开始到旋转点一直递减
// 二分查找的关键是让闭的那边可以不断向前，哪边闭那边就可以跳过
// 通过比较 target 和 nums[0]、nums[end]，提前判断出 target 应该在旋转点的哪一侧（称之为 isLeft）。这就像在进入迷宫前，先看地图确定目标在哪个大区。
func search(nums []int, target int) int {
	length := len(nums)
	right := length
	left := 0
	mid := 0
	firstMidNum := 0
	isRight := false
	//我们先来预处理
	if target < nums[left] {
		//比开头小的，可能在右侧
		if target <= nums[right-1] {
			//如果确实小于等于结尾，那么确实在右侧
			firstMidNum = nums[right-1]
			//剩下的属于左侧的数字都比右侧最大数字要大
			//后续遇到比这个数大的数就左截断
			isRight = true
		} else {
			return -1
		}
	} else {
		//这里是 target 大于等于开头数字
		//那么一定在左侧
		//属于右侧的数字一定比左侧最小的数字要小
		firstMidNum = nums[left]
		//后续遇到比这个数小的数就右截断
	}
	for left < right {
		mid = left + (right-left)/2
		if isRight {
			if nums[mid] > firstMidNum {
				left = mid + 1
				continue
			} else {
				if target > nums[mid] {
					left = mid + 1
				} else {
					right = mid
				}
			}
		} else {
			if nums[mid] < firstMidNum {
				right = mid
				continue
			} else {
				if target > nums[mid] {
					left = mid + 1
				} else {
					right = mid
				}
			}
		}
	}
	if left == length {
		return -1
	}
	if nums[left] != target {
		left = -1
	}
	return left
}

// @lc code=end

/*
// @lcpr case=start
// [4,5,6,7,0,1,2]\n0\n
// @lcpr case=end

// @lcpr case=start
// [4,5,6,7,0,1,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

*/
