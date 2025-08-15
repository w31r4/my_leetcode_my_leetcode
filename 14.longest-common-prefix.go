/*
 * @lc app=leetcode.cn id=14 lang=golang
 * @lcpr version=30204
 *
 * [14] 最长公共前缀
 *
 * https://leetcode.cn/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (44.71%)
 * Likes:    3319
 * Dislikes: 0
 * Total Accepted:    1.5M
 * Total Submissions: 3.3M
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * 编写一个函数来查找字符串数组中的最长公共前缀。
 *
 * 如果不存在公共前缀，返回空字符串 ""。
 *
 *
 *
 * 示例 1：
 *
 * 输入：strs = ["flower","flow","flight"]
 * 输出："fl"
 *
 *
 * 示例 2：
 *
 * 输入：strs = ["dog","racecar","car"]
 * 输出：""
 * 解释：输入不存在公共前缀。
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= strs.length <= 200
 * 0 <= strs[i].length <= 200
 * strs[i] 如果非空，则仅由小写英文字母组成
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func longestCommonPrefix(strs []string) string {
	// go 语言中的字符串是不可修改的，是一个整体，无法通过 strs[1][1] 的方式去直接访问字符串中某个字符
	//需要使用 []rune(strs) 类型强制转换才能访问单个字符
	//go 字符串库中包含 HasPrefix 函数，可以用以检查前后缀或是否存在某子串
	//本题思路：以最初字符串为初始字符串，将字符串数组中的字符串和它比较，看看是否存在该初始字符串前缀，如果不存在，那么初始字符串长度减一，直到存在为止，当初始字符串为空时，停止循环，直接返回""
	// exist := strs[0]
	// //以最初字符串为原点
	// for _, str := range strs {

	// 	for exist != "" && (!(strings.HasPrefix(str, exist))) {
	// 		exist = exist[:len(exist)-1]
	// 	}
	// 	if exist == "" {
	// 		break
	// 	}
	// }
	// return exist
	//法一为水平扫描，使用内置函数匹配，遇到不匹配的就--,直到匹配为止
	//法二为垂直扫描，从第一位开始，逐个匹配，遇到不匹配的就直接返回
	s0 := strs[0]
	for j := range s0 {
		//j 为 s0 索引
		for _, str := range strs {
			if j == len(str) || str[j] != s0[j] {
				return s0[:j]
				//左闭右开
			}
		}
	}
	return s0
}

// @lc code=end

/*
// @lcpr case=start
// ["flower","flow","flight"]\n
// @lcpr case=end

// @lcpr case=start
// ["dog","racecar","car"]\n
// @lcpr case=end

*/
