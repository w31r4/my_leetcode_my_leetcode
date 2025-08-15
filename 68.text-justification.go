/*
 * @lc app=leetcode.cn id=68 lang=golang
 * @lcpr version=30204
 *
 * [68] 文本左右对齐
 *
 * https://leetcode.cn/problems/text-justification/description/
 *
 * algorithms
 * Hard (56.02%)
 * Likes:    465
 * Dislikes: 0
 * Total Accepted:    104.5K
 * Total Submissions: 186.1K
 * Testcase Example:  '["This", "is", "an", "example", "of", "text", "justification."]\n16'
 *
 * 给定一个单词数组 words 和一个长度 maxWidth ，重新排版单词，使其成为每行恰好有 maxWidth 个字符，且左右两端对齐的文本。
 *
 * 你应该使用“贪心算法”来放置给定的单词；也就是说，尽可能多地往每行中放置单词。必要时可用空格 ' ' 填充，使得每行恰好有 maxWidth
 * 个字符。
 *
 * 要求尽可能均匀分配单词间的空格数量。如果某一行单词间的空格不能均匀分配，则左侧放置的空格数要多于右侧的空格数。
 *
 * 文本的最后一行应为左对齐，且单词之间不插入额外的空格。
 *
 * 注意：
 *
 *
 * 单词是指由非空格字符组成的字符序列。
 * 每个单词的长度大于 0，小于等于 maxWidth。
 * 输入单词数组 words 至少包含一个单词。
 *
 *
 *
 *
 * 示例 1:
 *
 * 输入：words = ["This", "is", "an", "example", "of", "text", "justification."],
 * maxWidth = 16
 * 输出：
 * [
 * "This    is    an",
 * "example  of text",
 * "justification.  "
 * ]
 *
 *
 * 示例 2:
 *
 * 输入:words = ["What","must","be","acknowledgment","shall","be"], maxWidth = 16
 * 输出：
 * [
 * "What   must   be",
 * "acknowledgment  ",
 * "shall be        "
 * ]
 * 解释：注意最后一行的格式应为 "shall be    " 而不是 "shall     be",
 * 因为最后一行应为左对齐，而不是左右两端对齐。
 * ⁠    第二行同样为左对齐，这是因为这行只包含一个单词。
 *
 *
 * 示例 3:
 *
 * 输入:words =
 * ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we","do"]，maxWidth
 * = 20
 * 输出：
 * [
 * "Science  is  what we",
 * ⁠ "understand      well",
 * "enough to explain to",
 * "a  computer.  Art is",
 * "everything  else  we",
 * "do                  "
 * ]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= words.length <= 300
 * 1 <= words[i].length <= 20
 * words[i] 由小写英文字母和符号组成
 * 1 <= maxWidth <= 100
 * words[i].length <= maxWidth
 *
 *
 */

// @lcpr-template-start
package leetcode

import (
	"strings"
)

// @lcpr-template-end
// @lc code=start
// 思路：双指针，先合并分散的字符串，i 指向当前的字符，j 指向最近的空格，当 i<maxwidth 时，i 不断向后，当 n=maxwidth 时，i=j,j 是当前字符串长度，j<=width,width-j=和 width 的差值为剩余需要不全的空格值
// 似乎不需要先合并，可以使用 len() 函数求出单个单词长度，每经过一个单词都要先看长度是否超过 maxwidth，没超过就将单词加入临时字符串并自动加上一个空格的长度，
// 如果超过了，将没超过的部
// 如何补齐剩余空格？
// 没有思路
// 使用 auto fill 函数
func fullJustify(words []string, maxWidth int) []string {

	temp := make([]string, 0)
	//临时数组
	ans := make([]string, 0)
	//答案
	width := 0

	for i := 0; i < len(words); i++ {
		width += len(words[i]) + 1
		//加上当前单词长度
		//加入一个空格的长度
		if width <= maxWidth+1 {
			//如果加入当前单词不越界
			//当加入当前单词长度超过最大值时，处理字符串
			//加入一个空格的长度
			temp = append(temp, words[i])
		} else {
			//如果加入当前单词越界
			ans = append(ans, autoFill(temp, maxWidth))
			//autoFill 会在中间插入空格
			temp = []string{}
			//清空临时数组
			temp = append(temp, words[i])
			//加入当前单词到空数组
			width = len(words[i]) + 1
			//重置 width
			//将当前单词重新长度 +1 赋值给 width

		}
		if i == len(words)-1 {
			ans = append(ans, leftFill(temp, maxWidth))
			//最后一位使用 leftFill
		}
	}
	return ans
}

func autoFill(words []string, maxWidth int) string {
	//传入最大值和目标字符串
	width := 0
	space_num := max(len(words)-1, 1)
	//需要的填入空格的次数
	var builder strings.Builder
	for i, _ := range words {
		width += len(words[i])
	}
	//获取单词占位
	space_to_fill := maxWidth - width
	//得到剩余的空格位置
	base_space := space_to_fill / space_num
	//基础单次填入空格值
	left_space := space_to_fill % space_num
	//剩余空格值
	gap_num := make([]int, space_num)

	for i := 0; i < len(gap_num); i++ {
		gap_num[i] = base_space
		if i < left_space {
			gap_num[i]++
		}
	}

	for i, val := range words {
		builder.WriteString(val)
		if i <= space_num-1 {
			for n := 0; n < gap_num[i]; n++ {
				builder.WriteRune(' ')
			}
		}
	}
	return builder.String()
}
func leftFill(words []string, maxWidth int) string {
	width := 0
	var builder strings.Builder
	for i, val := range words {
		width += len(words[i])
		builder.WriteString(val)
		if i != len(words)-1 {
			builder.WriteRune(' ')
		}
	}
	space_left := maxWidth - width - len(words) + 1
	//空格剩余=最大长度 - 单词长度 - 单词间的空格
	for i := 0; i < space_left; i++ {
		builder.WriteRune(' ')
	}
	return builder.String()
}

// @lc code=end

/*
// @lcpr case=start
// ["This", "is", "an", "example", "of", "text", "justification."]\n16\n
// @lcpr case=end

// @lcpr case=start
// ["What","must","be","acknowledgment","shall","be"]\n16\n
// @lcpr case=end

// @lcpr case=start
// ["Science","is","what","we","understand","well","enough","to","explain","to","a","computer.","Art","is","everything","else","we"\n20\n
// @lcpr case=end

*/
