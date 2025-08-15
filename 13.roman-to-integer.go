/*
 * @lc app=leetcode.cn id=13 lang=golang
 * @lcpr version=30204
 *
 * [13] 罗马数字转整数
 *
 * https://leetcode.cn/problems/roman-to-integer/description/
 *
 * algorithms
 * Easy (64.00%)
 * Likes:    2924
 * Dislikes: 0
 * Total Accepted:    1.2M
 * Total Submissions: 1.8M
 * Testcase Example:  '"III"'
 *
 * 罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。
 *
 * 字符          数值
 * I             1
 * V             5
 * X             10
 * L             50
 * C             100
 * D             500
 * M             1000
 *
 * 例如，罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。27 写做  XXVII, 即为 XX + V
 * + II 。
 *
 * 通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数
 * 5 减小数 1 得到的数值 4。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：
 *
 *
 * I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
 * X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
 * C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。
 *
 *
 * 给定一个罗马数字，将其转换成整数。
 *
 *
 *
 * 示例 1:
 *
 * 输入: s = "III"
 * 输出：3
 *
 * 示例 2:
 *
 * 输入: s = "IV"
 * 输出：4
 *
 * 示例 3:
 *
 * 输入: s = "IX"
 * 输出：9
 *
 * 示例 4:
 *
 * 输入: s = "LVIII"
 * 输出：58
 * 解释：L = 50, V= 5, III = 3.
 *
 *
 * 示例 5:
 *
 * 输入: s = "MCMXCIV"
 * 输出：1994
 * 解释：M = 1000, CM = 900, XC = 90, IV = 4.
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 15
 * s 仅含字符 ('I', 'V', 'X', 'L', 'C', 'D', 'M')
 * 题目数据保证 s 是一个有效的罗马数字，且表示整数在范围 [1, 3999] 内
 * 题目所给测试用例皆符合罗马数字书写规则，不会出现跨位等情况。
 * IL 和 IM 这样的例子并不符合题目要求，49 应该写作 XLIX，999 应该写作 CMXCIX。
 * 关于罗马数字的详尽书写规则，可以参考 罗马数字 - 百度百科。
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
var ROMAN = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

//查找值用哈希表
//代码实现时，可以创建一个哈希表，把字符映射为对应的数值，这样无需写一堆 if-else。

func romanToInt(s string) int {
	// 遍历麻烦
	//
	//	for i := 0; i < len(s) && len(s) >= 2; i++ {
	//		switch {
	//		case i < len(s)-1 && s[i] == 'I' && s[i+1] == 'V':
	//			sum += 4
	//			i++
	//		case i < len(s)-1 && s[i] == 'I' && s[i+1] == 'X':
	//			sum += 9
	//			i++
	//		case i < len(s)-1 && s[i] == 'X' && s[i+1] == 'L':
	//			sum += 40
	//			i++
	//		case i < len(s)-1 && s[i] == 'X' && s[i+1] == 'C':
	//			sum += 90
	//			i++
	//		case i < len(s)-1 && s[i] == 'C' && s[i+1] == 'D':
	//			sum += 400
	//			i++
	//		case i < len(s)-1 && s[i] == 'C' && s[i+1] == 'M':
	//			sum += 900
	//			i++
	//		case
	//		}
	//	}

	ans := 0
	for i := 1; i < len(s); i++ { // 遍历相邻的罗马数字
		x, y := ROMAN[s[i-1]], ROMAN[s[i]]
		if x < y {
			ans -= x
			//将规则统一，如果前数小于后数，那么减去前数
			// iv==-1+4
		} else {
			ans += x
		}
	}
	return ans + ROMAN[s[len(s)-1]] // 加上最后一个
}

// @lc code=end

/*
// @lcpr case=start
// "III"\n
// @lcpr case=end

// @lcpr case=start
// "IV"\n
// @lcpr case=end

// @lcpr case=start
// "IX"\n
// @lcpr case=end

// @lcpr case=start
// "LVIII"\n
// @lcpr case=end

// @lcpr case=start
// "MCMXCIV"\n
// @lcpr case=end

*/
