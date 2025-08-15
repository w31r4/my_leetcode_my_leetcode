/*
 * @lc app=leetcode.cn id=219 lang=golang
 * @lcpr version=30204
 *
 * [219] 存在重复元素 II
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 先遍历前 k+1 个数字，如果找到就返回 true
// 如果没有找到，那么将 [1:k+1] 个数字存入哈希表，建立一个指针指向
func containsNearbyDuplicate(nums []int, k int) bool {

	// longMap := make(map[int]int)
	// lengthOfNums := len(nums)
	// if lengthOfNums < k+1 {
	// 	//先处理数组长度小于 k+1 的情况
	// 	if lengthOfNums == 1 {
	// 		return false
	// 	}
	// 	for _, v := range nums {
	// 		longMap[v]++
	// 		if longMap[v] > 1 {
	// 			return true
	// 		}
	// 	}
	// 	return false
	// }
	// for _, v := range nums[:k+1] {
	// 	longMap[v]++
	// }
	// for i := 0; i < lengthOfNums; i++ {
	// 	if longMap[nums[i]] > 1 {
	// 		return true
	// 	} else {
	// 		longMap[nums[i]]--
	// 	}
	// 	boundary := i + k + 1
	// 	if boundary < lengthOfNums {
	// 		longMap[nums[boundary]]++
	// 	}
	// }
	// return false
	//方法太笨，没必要先建立 map 后查找，可以边建立边查找
	// longMap := map[int]int{}
	//该方法是维护 nums 中所有的值，如果有相等的且索引值差小于等于 k，那么返回 true
	// for i, val := range nums {
	// 	if j, ok := longMap[val]; ok && i-j <= k {
	// 		return true
	// 	}
	// 	longMap[val] = i
	// }
	// return false
	//这个方法和我的思路一样，但又略有不同，维护一个长度为 min(k+1，当前遍历到的索引长度 +1) 的滑动窗口，从尾端查询之前的窗口中是否存在相同元素，而不是从头开始，如果不存在，将当前元素加入，移除溢出窗口的元素，也就是窗口的第一个元素
	set := map[int]bool{}

	for i, x := range nums {
		if set[x] == true {
			return true
		}
		set[x] = true
		if i >= k {
			delete(set, nums[i-k])
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1,0,1,1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,1,2,3]\n2\n
// @lcpr case=end

*/
