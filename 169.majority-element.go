/*
 * @lc app=leetcode.cn id=169 lang=golang
 * @lcpr version=30204
 *
 * [169] 多数元素
 *
 * https://leetcode.cn/problems/majority-element/description/
 *
 * algorithms
 * Easy (66.68%)
 * Likes:    2454
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 1.8M
 * Testcase Example:  '[3,2,3]'
 *
 * 给定一个大小为 n 的数组 nums，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
 *
 * 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [3,2,3]
 * 输出：3
 *
 * 示例 2：
 *
 * 输入：nums = [2,2,1,1,1,2,2]
 * 输出：2
 *
 *
 *
 * 提示：
 *
 *
 * n == nums.length
 * 1 <= n <= 5 * 10^4
 * -10^9 <= nums[i] <= 10^9
 *
 *
 *
 *
 * 进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func majorityElement(nums []int) int {
	// hashmap := make(map[int]int)
	// length := len(nums)
	// for _, val := range nums {
	// 	hashmap[val] += 1
	// 	if hashmap[val] > length/2 {
	// 		return val
	// 	}
	// }
	// return int(0)
	//哈希表法，On 时间，On 空间
	//还可以投票法，当有相同数字时给自己加票，当为政敌时减票，票数为 0 则换届选举
	//这道题实质是看那种数字在数组中最多
	flags := 0
	major := 0
	for _, val := range nums {

		if flags == 0 {
			major = val
		}
		if val == major {
			flags++
		} else if val != major {
			flags--
		}

	}
	return major
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [2,2,1,1,1,2,2]\n
// @lcpr case=end

*/
