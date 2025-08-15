/*
 * @lc app=leetcode.cn id=6 lang=golang
 * @lcpr version=30204
 *
 * [6] Z 字形变换
 *
 * https://leetcode.cn/problems/zigzag-conversion/description/
 *
 * algorithms
 * Medium (53.69%)
 * Likes:    2488
 * Dislikes: 0
 * Total Accepted:    782.3K
 * Total Submissions: 1.5M
 * Testcase Example:  '"PAYPALISHIRING"\n3'
 *
 * 将一个给定字符串 s 根据给定的行数 numRows，以从上往下、从左到右进行 Z 字形排列。
 *
 * 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 * 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
 *
 * 请你实现这个将字符串进行指定行数变换的函数：
 *
 * string convert(string s, int numRows);
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "PAYPALISHIRING", numRows = 3
 * 输出："PAHNAPLSIIGYIR"
 *
 * 示例 2：
 *
 * 输入：s = "PAYPALISHIRING", numRows = 4
 * 输出："PINALSIGYAHRPI"
 * 解释：
 * P     I    N
 * A   L S  I G
 * Y A   H R
 * P     I
 *
 *
 * 示例 3：
 *
 * 输入：s = "A", numRows = 1
 * 输出："A"
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 1000
 * s 由英文字母（小写和大写）、',' 和 '.' 组成
 * 1 <= numRows <= 1000
 *
 *
 */

// @lcpr-template-start
package leetcode

import "strings"

// @lcpr-template-end
// @lc code=start
func convert(s string, numRows int) string {
	//思路如下：画出索引构成的数字 Z 形图，由图像观察发现，index%(2*numRows-2)==0 的在构成的 Z 形图第一行，index%(2*numRows-2)==1 和 index%(2*numRows-2)==(2*numRows-1) 的在第二行，如此类推，直到 index%(numRows-1)==0 的在最后一行
	//ok，初见杀
	if numRows == 1 {
		return s
	}
	loop_num := 2*numRows - 2
	length := len(s) - 1
	var builder strings.Builder
	for i := 0; i <= numRows-1; i++ {
		times := 1
		for n := i; n <= length; n += loop_num {
			builder.WriteByte(s[n])
			if n%(numRows-1) != 0 && times*loop_num-i <= length {
				builder.WriteByte(s[times*loop_num-i])
			}
			times++
		}
	}
	return builder.String()

}

// @lc code=end

/*
// @lcpr case=start
// "PAYPALISHIRING"\n3\n
// @lcpr case=end

// @lcpr case=start
// "PAYPALISHIRING"\n4\n
// @lcpr case=end

// @lcpr case=start
// "A"\n1\n
// @lcpr case=end

*/
