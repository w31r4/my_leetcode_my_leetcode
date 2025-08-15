/*
 * @lc app=leetcode.cn id=228 lang=golang
 * @lcpr version=30204
 *
 * [228] 汇总区间
 */

// @lcpr-template-start
package leetcode

import (
	"strconv"
)

// @lcpr-template-end
// @lc code=start
func summaryRanges(nums []int) (ans []string) {
	// if len(nums) == 0 {
	// 	return
	// }
	// if len(nums) == 1 {
	// 	ans = append(ans, strconv.Itoa(nums[0]))
	// 	return
	// }
	// var builder strings.Builder
	// builder.WriteString(strconv.Itoa(nums[0]))
	// isFirstTime := true
	// for i, num := range nums[1:] {
	// 	if num-1 == nums[i] {
	// 		//当前数字减一等于前一位数字
	// 		if isFirstTime {
	// 			builder.WriteString("->")
	// 			//如果是第一次连续区间，打印上"->""
	// 		}
	// 		isFirstTime = false
	// 		if i+2 == len(nums) {
	// 			//如果已经是最后一位了依旧连续
	// 			builder.WriteString(strconv.Itoa(num))
	// 			ans = append(ans, builder.String())
	// 		}
	// 		continue
	// 	}
	// 	//当此时不连续的情况下，如果区间内仅有一个数字
	// 	if isFirstTime {
	// 		ans = append(ans, builder.String())
	// 		builder.Reset()
	// 		builder.WriteString(strconv.Itoa(num))
	// 		///将 builder 清空将当前数字写入 builder 中
	// 		if i+2 == len(nums) {
	// 			//如果已经是最后一位
	// 			ans = append(ans, builder.String())
	// 		}
	// 		continue
	// 	}

	// 	//此时已经不连续，且区间不止一个数字
	// 	isFirstTime = true
	// 	builder.WriteString(strconv.Itoa(nums[i]))
	// 	ans = append(ans, builder.String())
	// 	builder.Reset()
	// 	builder.WriteString(strconv.Itoa(num))
	// 	if i+2 == len(nums) {
	// 		//如果已经是最后一位
	// 		ans = append(ans, builder.String())
	// 	}
	// }
	// return
	//上述思路略显复杂，建议使用下面的思路
	//维护一个循环即可
	// 使用一个主循环遍历数组（例如，索引为 i）。
	// 对于每个元素 nums[i]，将其视为一个潜在区间的开始。
	// 使用一个内部逻辑（可以是内部循环或指针 j）来查找这个连续区间的结束点（即 nums[j+1] == nums[j]+1）。
	// 一旦找到区间的结束点 nums[j]，根据 i 和 j 的关系来格式化字符串：
	// 如果 i == j，则为单个数字。
	// 如果 i < j，则为 nums[i]->nums[j]。
	// 将格式化后的字符串添加到结果列表中。
	// 将主循环的索引 i 更新为 j+1，从下一个新数字开始查找。
	//总结：少用 range 多用 i
	if len(nums) == 0 {
		return []string{}
	}

	n := len(nums)
	for i := 0; i < n; {
		start := nums[i] // 当前区间的开始值
		j := i           // 区间结束的索引

		// 向前查找连续区间的结束
		// 只要下一个元素存在 (j+1 < n) 并且是连续的 (nums[j+1] == nums[j]+1)
		for j+1 < n && nums[j+1] == nums[j]+1 {
			j++
		}

		// 形成区间字符串
		if i == j {
			// 区间只有一个数字
			ans = append(ans, strconv.Itoa(start))
		} else {
			// 区间有多个数字
			// 可以使用 strings.Builder 进一步优化这里的字符串拼接，但为了简洁性此处直接拼接
			ans = append(ans, strconv.Itoa(start)+"->"+strconv.Itoa(nums[j]))
		}

		// 移动 i 到下一个区间的开始位置
		i = j + 1
	}
	return ans
}

// @lc code=end

/*
// @lcpr case=start
// [0,1,2,4,5,7]\n
// @lcpr case=end

// @lcpr case=start
// [0,2,3,4,6,8,9]\n
// @lcpr case=end

*/
