/*
 * @lc app=leetcode.cn id=918 lang=golang
 * @lcpr version=30204
 *
 * [918] 环形子数组的最大和
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func maxSubarraySumCircular(nums []int) (answer int) {
	// length := len(nums)
	// newNums := make([]int, 0, length*2)
	// newNums = append(newNums, nums...)
	// newNums = append(newNums, nums...)
	// answer = nums[0]
	// for i := 0; i < length; i++ {
	// 	nums1 := newNums[i : i+length]
	// 	maxSubSum := 0
	// 	preSubSum := 0
	// 	for _, val := range nums1 {
	// 		maxSubSum = max(preSubSum+val, val)
	// 		preSubSum = maxSubSum
	// 		answer = max(answer, maxSubSum)
	// 	}
	// }
	//N 平方，时间超了
	maxSubSum := 0
	preSubSum := 0
	maxOrigin := nums[0]
	sum := 0
	for _, val := range nums {
		maxSubSum = max(preSubSum+val, val)
		preSubSum = maxSubSum
		maxOrigin = max(maxOrigin, maxSubSum)
		sum += val
	}
	minSubSum := 0
	preMinSum := 0
	maxRing := nums[0]
	//找到当前数组中最小的部分，将其移出环，这样剩下的就是最大的了
	for _, val := range nums {
		minSubSum = min(preMinSum+val, val)
		preMinSum = minSubSum
		maxRing = min(maxRing, minSubSum)
	}
	if maxRing == sum {
		return maxOrigin
	}
	maxRing = sum - maxRing
	answer = max(maxOrigin, maxRing)
	return
	//很妙的一个思路
	//将环分为了非跨和跨越式
	// 你问到了这个解法的最核心、最关键的地方！这是一个非常深刻的问题。

	// 你说得完全正确：从概念上讲，在一个真正的圆环里，任何一段子数组都是平等的，没有“跨越”和“不跨越”的区别。

	// 我们之所以要分两种情况，**完全是因为我们用来存储数据的工具——数组——是线性的，而不是环形的**。

	// 我们最强大的工具，**Kadane's 算法**，是一个处理**线性连续**子数组的大师。它只能处理 `[ . . . [一段连续的子数组] . . . ]` 这样的情况。

	// 现在，我们把环形数组拉直成一个线性数组，看看最大和子数组会是什么样子：

	// **情况一：最大和子数组是“普通”的，没有绕圈。**
	// 在线性数组里，它就是一个完整的、连续的块。
	// `[ . . . [最大子数组] . . . ]`
	// 对于这种情况，Kadane's 算法简直是为它量身定做的，可以直接找到答案。

	// **情况二：最大和子数组是“环形”的，绕了一圈。**
	// 在线性数组里，它其实是**两段分离的块**：一段在数组的末尾，另一段在数组的开头。
	// `[最大子数组的前半段] . . . . . [最大子数组的后半段]`
	// Kadane's 算法看到这个会“懵掉”，因为它一次只能处理一个连续的块，它没法把这两段加起来。

	// 这就是问题的关键：我们有一个很棒的工具，但它只能处理情况一，处理不了情况二。

	// 所以我们必须想个办法。这个办法就是你之前发现的那个聪明的“反向思维”：
	// 我们不直接去求那两段分离的块的和，而是反过来，去求**中间那段我们不想要的、连续的块**。

	// `[环形部分] [中间不想要的部分] [环形部分]`

	// 为了让“环形部分”的和最大，那么“中间不想要的部分”的和就必须**最小**。

	// 你看，我们把一个我们不会解的问题（求两段分离块的最大和），成功转化成了一个我们会解的问题（求一段连续块的最小和）。而求最小和，只需要稍微修改一下 Kadane's 算法就行了。

	// 总结一下：
	// *   **情况一**：最大和是普通连续块。我们用**标准 Kadane's** 解决。
	// *   **情况二**：最大和是环形块（在数组里是两头）。我们通过求**最小连续块**，再用**总和减去它**来间接解决。

	// 因为我们事先不知道答案会属于哪种情况，所以我们必须两种情况都计算出来，然后取它们中更大的那个作为最终答案。

	// 这个解释能让你理解我们为什么必须“兵分两路”来解决这个问题了吗？
}

// @lc code=end

/*
// @lcpr case=start
// [1,-2,3,-2]\n
// @lcpr case=end

// @lcpr case=start
// [5,-3,5]\n
// @lcpr case=end

// @lcpr case=start
// [3,-2,2,-3]\n
// @lcpr case=end

*/
