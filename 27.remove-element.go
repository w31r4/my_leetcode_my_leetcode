/*
 * @lc app=leetcode.cn id=27 lang=golang
 * @lcpr version=30204
 *
 * [27] 移除元素
 *
 * https://leetcode.cn/problems/remove-element/description/
 *
 * algorithms
 * Easy (60.31%)
 * Likes:    2480
 * Dislikes: 0
 * Total Accepted:    1.9M
 * Total Submissions: 3.1M
 * Testcase Example:  '[3,2,2,3]\n3'
 *
 * 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素。元素的顺序可能发生改变。然后返回 nums 中与 val
 * 不同的元素的数量。
 *
 * 假设 nums 中不等于 val 的元素数量为 k，要通过此题，您需要执行以下操作：
 *
 *
 * 更改 nums 数组，使 nums 的前 k 个元素包含不等于 val 的元素。nums 的其余元素和 nums 的大小并不重要。
 * 返回 k。
 *
 *
 * 用户评测：
 *
 * 评测机将使用以下代码测试您的解决方案：
 *
 * int[] nums = [...]; // 输入数组
 * int val = ...; // 要移除的值
 * int[] expectedNums = [...]; // 长度正确的预期答案。
 * ⁠                           // 它以不等于 val 的值排序。
 *
 * int k = removeElement(nums, val); // 调用你的实现
 *
 * assert k == expectedNums.length;
 * sort(nums, 0, k); // 排序 nums 的前 k 个元素
 * for (int i = 0; i < actualLength; i++) {
 * ⁠   assert nums[i] == expectedNums[i];
 * }
 *
 * 如果所有的断言都通过，你的解决方案将会 通过。
 *
 *
 *
 * 示例 1：
 *
 * 输入：nums = [3,2,2,3], val = 3
 * 输出：2, nums = [2,2,_,_]
 * 解释：你的函数函数应该返回 k = 2, 并且 nums 中的前两个元素均为 2。
 * 你在返回的 k 个元素之外留下了什么并不重要（因此它们并不计入评测）。
 *
 * 示例 2：
 *
 * 输入：nums = [0,1,2,2,3,0,4,2], val = 2
 * 输出：5, nums = [0,1,4,0,3,_,_,_]
 * 解释：你的函数应该返回 k = 5，并且 nums 中的前五个元素为 0,0,1,3,4。
 * 注意这五个元素可以任意顺序返回。
 * 你在返回的 k 个元素之外留下了什么并不重要（因此它们并不计入评测）。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 0 <= nums.length <= 100
 * 0 <= nums[i] <= 50
 * 0 <= val <= 100
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func removeElement(nums []int, val int) int {
	// length := len(nums) - 1
	// i := 0
	// for i <= length {
	// 	if nums[i] == val {
	// 		nums[i] = nums[length]
	// 		length--
	// 	} else {
	// 		i++
	// 	}
	// }

	// return length + 1
	//我的解法，双指针相向，老是搞不清楚判定条件在哪，最好可以画个图看看，这样理解:i 和 length 一定会有一个相遇处，且在这个一定是 nums[i]==val 的地方，此时 length==i，一定会触发 length--,
	//让 length 回到它忠实的索引上，我们重点要搞清楚的是当 i==length 时在何处发生了什么，这样就能写好答案
	//当然，这题是我想复杂了，不用在乎 k 位后是什么
	k := 0
	for _, num := range nums {
		if num != val {
			//当找到 num==val 时，保持 k 不变，等待下一个 num!=val，将其传入 nums[k] 中，然后将 k 进位
			//简单来讲就是有几个 val 慢指针就停几次，在快指针非 val 时将慢指针所在的 val 值重新赋值为 val，避免数组出现 val
			//快指针无重复地遍历了所有变量，所以慢指针不会重复，重复的会被后面的人覆盖
			//覆盖中的慢指针所指向的可以看作一个无值空数组
			nums[k] = num
			k++
		}
	}
	return k
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,1,0,0,1,1,5,7,9]\n0\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,2,3,0,4,2]\n2\n
// @lcpr case=end

*/
