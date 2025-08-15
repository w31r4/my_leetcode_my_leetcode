/*
 * @lc app=leetcode.cn id=80 lang=golang
 * @lcpr version=30204
 *
 * [80] 删除有序数组中的重复项 II
 *
 * https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/description/
 *
 * algorithms
 * Medium (63.30%)
 * Likes:    1248
 * Dislikes: 0
 * Total Accepted:    571.5K
 * Total Submissions: 902.3K
 * Testcase Example:  '[1,1,1,2,2,3]'
 *
 * 给你一个有序数组 nums，请你 原地 删除重复出现的元素，使得出现次数超过两次的元素只出现两次，返回删除后数组的新长度。
 *
 * 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
 *
 *
 *
 * 说明：
 *
 * 为什么返回数值是整数，但输出的答案是数组呢？
 *
 * 请注意，输入数组是以「引用」方式传递的，这意味着在函数里修改输入数组对于调用者是可见的。
 *
 * 你可以想象内部操作如下：
 *
 * // nums 是以“引用”方式传递的。也就是说，不对实参做任何拷贝
 * int len = removeDuplicates(nums);
 *
 * // 在函数里修改输入数组对于调用者是可见的。
 * // 根据你的函数返回的长度，它会打印出数组中 该长度范围内 的所有元素。
 * for (int i = 0; i < len; i++) {
 * print(nums[i]);
 * }
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [1,1,1,2,2,3]
 * 输出：5, nums = [1,1,2,2,3]
 * 解释：函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3。不需要考虑数组中超出新长度后面的元素。
 *
 *
 * 示例 2：
 *
 * 输入：nums = [0,0,1,1,1,1,2,3,3]
 * 输出：7, nums = [0,0,1,1,2,3,3]
 * 解释：函数应返回新长度 length = 7, 并且原数组的前七个元素被修改为 0, 0, 1, 1, 2, 3,
 * 3。不需要考虑数组中超出新长度后面的元素。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 3 * 10^4
 * -10^4 <= nums[i] <= 10^4
 * nums 已按升序排列
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func removeDuplicates(nums []int) int {
	// if len(nums) == 0 {
	// 	return 0
	// }
	// i := 0                           // 慢指针，标记有效数组的末尾
	// count := 1                       // 计数器，记录当前元素的重复次数
	// for j := 1; j < len(nums); j++ { // 快指针 j 遍历整个数组
	// 	if nums[j] == nums[i] { // 当前元素与慢指针元素相同
	// 		count++         // 增加重复计数
	// 		if count <= 2 { // 若重复不超过两次
	// 			i++               // 移动慢指针
	// 			nums[i] = nums[j] // 将当前元素复制到有效位置
	// 		}
	// 	} else { // 遇到新元素
	// 		i++               // 移动慢指针
	// 		nums[i] = nums[j] // 复制新元素到有效位置
	// 		count = 1         // 重置计数器为 1
	// 	}
	// }
	// return i + 1 // 有效数组长度为 i+1（i 是索引，长度需 +1）
	//双指针法
	// stackSize := 2
	// for i := 2; i < len(nums); i++ {
	// 	if nums[i] != nums[stackSize-2] {
	// 		nums[stackSize] = nums[i]
	// 		stackSize++
	// 	}
	// 	// 栈法，快指针指向大数组，栈指针维护创建小数组，栈头栈尾差两位，如果快指针等于栈尾，那么就不用赋值，快指针向下寻找不等值，如果快指针不等于栈尾则将栈头用快指针值覆盖，和快慢指针的数组是一样的，但是这里的基本单元为 2 位的栈，使得可以存下 2 位的相同数字
	// }
	// return min(stackSize, len(nums))
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,1,2,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,1,1,1,1,2,3,3]\n
// @lcpr case=end

*/
