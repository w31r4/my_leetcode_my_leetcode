/*
 * @lc app=leetcode.cn id=128 lang=golang
 * @lcpr version=30204
 *
 * [128] 最长连续序列
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 维护一个哈希表 matchMap，但是哈希表不能按 key 大小遍历遍历
// 我们可以维护一个数组 keys，存储哈希表中出现过的 key，将其排序后用于顺序索引
// 不要用 sort 函数
// 可以增加时间复杂度
// 比如遍历哈希表的同时为每一个表中元素找下一位直到没有为止
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	// matchMap := map[int]bool{}
	// keys := []int{}
	// for _, val := range nums {
	// 	if matchMap[val] != true {
	// 		keys = append(keys, val)
	// 		matchMap[val] = true
	// 	}
	// }
	// sort.Ints(keys)
	// //这个函数可以用，哈希表顺序查找不可或缺
	// ans := 1
	// count := 1
	// current := keys[0]
	// for _, val := range keys[1:] {
	// 	if val == current+1 {
	// 		count++
	// 		current = val
	// 	} else {
	// 		current = val
	// 		count = 1
	// 	}
	// 	ans = max(ans, count)
	// }
	// return ans
	matchMap := map[int]bool{}
	for _, v := range nums {

		matchMap[v] = true

	}
	ans := 0
	for num := range matchMap {
		//哈希表随机遍历
		//因为我们要找最长的序列长度，所以当当前元素前一位存在时我们放弃查找以这个元素为开头的最长连续序列
		if matchMap[num-1] {
			continue
		}
		nextNum := num + 1
		for matchMap[nextNum] {
			nextNum++
		}
		ans = max(ans, nextNum-num)
	}
	return ans

}

// @lc code=end

/*
// @lcpr case=start
// [100,4,200,1,3,2]\n
// @lcpr case=end

// @lcpr case=start
// [0,3,7,2,5,8,4,6,0,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,0,1,2]\n
// @lcpr case=end

*/
