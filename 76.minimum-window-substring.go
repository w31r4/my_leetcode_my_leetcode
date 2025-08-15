/*
 * @lc app=leetcode.cn id=76 lang=golang
 * @lcpr version=30204
 *
 * [76] 最小覆盖子串
 *
 * https://leetcode.cn/problems/minimum-window-substring/description/
 *
 * algorithms
 * Hard (47.22%)
 * Likes:    3234
 * Dislikes: 0
 * Total Accepted:    782.2K
 * Total Submissions: 1.6M
 * Testcase Example:  '"ADOBECODEBANC"\n"ABC"'
 *
 * 给你一个字符串 s、一个字符串 t。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 ""
 * 。
 *
 *
 *
 * 注意：
 *
 *
 * 对于 t 中重复字符，我们寻找的子字符串中该字符数量必须不少于 t 中该字符数量。
 * 如果 s 中存在这样的子串，我们保证它是唯一的答案。
 *
 *
 *
 *
 * 示例 1：
 *
 * 输入：s = "ADOBECODEBANC", t = "ABC"
 * 输出："BANC"
 * 解释：最小覆盖子串 "BANC" 包含来自字符串 t 的 'A'、'B' 和 'C'。
 *
 *
 * 示例 2：
 *
 * 输入：s = "a", t = "a"
 * 输出："a"
 * 解释：整个字符串 s 是最小覆盖子串。
 *
 *
 * 示例 3:
 *
 * 输入：s = "a", t = "aa"
 * 输出：""
 * 解释：t 中两个字符 'a' 均应包含在 s 的子串中，
 * 因此没有符合条件的子字符串，返回空字符串。
 *
 *
 *
 * 提示：
 *
 *
 * ^m == s.length
 * ^n == t.length
 * 1 <= m, n <= 10^5
 * s 和 t 由英文字母组成
 *
 *
 *
 * 进阶：你能设计一个在 o(m+n) 时间内解决此问题的算法吗？
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 思路：先找到一个覆盖子串，然后左指针不断向右移动直到遇到目标子串中的元素，记录长度后移去该元素，并不断移动到下一个目标字串元素，右指针开始向右遍历，直到遇到该元素或者子串中的元素，如果遇到先该元素记录长度，将左指针继续不断向右移动直到遇到目标子串元素，进入循环，如果遇到其他元素，左指针不断位移直到遇到该元素，期间记录损失元素，当损失元素全部为 0 时，记录长度
//和标准答案就输一手 map 转用 [128]int{}
//map 性能太低了
// func minWindow(s string, t string) string {
// 	hashMap := make(map[byte]int)
// 	sLen := len(s)
// 	left, right := 0, 0
// 	for i := range t {
// 		hashMap[t[i]]++
// 	}
// 	tLen := len(t)
// 	foundByte := 0
// 	nowMap := make(map[byte]int)
// 	minLeft, minRight := 0, 100001
// 	for ; left <= right && right < sLen; right++ {
// 		if hashCount, ok := hashMap[s[right]]; ok {

// 			if nowMap[s[right]] > hashCount {
// 				continue
// 			}
// 			//当某字符出现次数超越预期时，跳过计算
// 			foundByte++
// 			// if hashCount < nowMap[s[right]] {
// 			//以下是对于 t 中重复字符，我们寻找的子字符串中该字符数量必须等于 t 中该字符数量的解法
// 			//又没好好读题
// 			// for left <= right && s[left] != s[right] {
// 			// 	if hashCount, ok := nowMap[s[left]]; ok {
// 			// 		if hashCount > 0 {
// 			// 			nowMap[s[left]]--
// 			// 			foundByte--
// 			// 			if nowMap[s[left]] == 0 {
// 			// 				delete(nowMap, s[left])
// 			// 			}
// 			// 		}
// 			// 	}
// 			// 	left++
// 			// }
// 			// if s[left] == s[right] {
// 			// 	nowMap[s[left]]--
// 			// 	if nowMap[s[left]] == 0 {
// 			// 		delete(nowMap, s[left])
// 			// 	}
// 			// 	foundByte--
// 			// 	left++
// 			// }

// 		}
// 		for foundByte == tLen {
// 			//跳过所有非目标 left 指针和过多次数字符，直到遇到次数恰好的字符
// 			if hashCount, ok := nowMap[s[left]]; ok {
// 				if hashCount == hashMap[s[left]] {
// 					nowMap[s[left]]--
// 					if nowMap[s[left]] == 0 {
// 						delete(nowMap, s[left])
// 					}
// 					left++
// 					foundByte--
// 					if right-left+1 < minRight-minLeft {
// 						minLeft = left - 1
// 						minRight = right
// 					}
// 				} else {
// 					nowMap[s[left]]--
// 					left++
// 				}
// 			} else {
// 				left++
// 			}

// 		}
// 	}
// 	if minRight == 100001 {
// 		return ""
// 	}

//		return s[minLeft : minRight+1]
//	}
func minWindow(s string, t string) string {
	// 使用大小为 128 的 int 数组替代 map[byte]int
	// 假设输入字符串只包含 ASCII 值在 0-127 范围内的字符
	var requiredCounts [128]int // 存储 t 中每个字符需要的数量 (替代 hashMap)
	var currentCounts [128]int  // 存储当前滑动窗口中每个字符的数量 (替代 nowMap)

	sLen := len(s)
	tLen := len(t)

	// 处理特殊情况
	if tLen == 0 || sLen == 0 || sLen < tLen {
		return ""
	}

	// 填充 requiredCounts 数组
	for i := range t {
		char := t[i]
		// 检查字符是否在数组索引范围内
		if char < 128 {
			requiredCounts[char]++
		} else {
			// 如果 t 包含非 ASCII 字符 (>= 128)，则这个 128 大小的数组无法完全处理
			// 此时应考虑使用 map 或 [256]int{}
			// 按照题目要求，我们忽略大于等于 128 的字符在 requiredCounts 中
			// 或者更严格地返回错误，这里选择忽略，但实际应用需谨慎
			// fmt.Printf("Warning: Character '%c' (byte value %d) in t is >= 128, cannot be fully tracked with [128]int{}.\n", char, char)
		}
	}

	left, right := 0, 0
	// foundCount 记录当前窗口中，满足 requiredCounts 数量的 t 中字符实例的总数
	// 当 foundCount 等于 tLen 时，表示窗口包含了 t 中所有字符（包括重复）
	foundCount := 0

	// 记录找到的最小窗口的起始位置和长度
	minLen := sLen + 1 // 初始化为大于任何可能窗口长度的值
	minStart := 0

	// 滑动窗口 [left ... right] (right 是包含的边界)
	for right < sLen {
		charRight := s[right] // 当前进入窗口的字符

		// 检查 charRight 是否是我们关心的字符 (即在 t 中出现过，且在数组索引范围内)
		if charRight < 128 && requiredCounts[charRight] > 0 {
			// 更新当前窗口中该字符的计数
			currentCounts[charRight]++

			// 如果当前字符的计数还没有超过 t 中所需的数量，foundCount 增加
			// 这样 foundCount 最终会等于 tLen 当且仅当窗口包含了所有必需的字符实例
			if currentCounts[charRight] <= requiredCounts[charRight] {
				foundCount++
			}
		}
		// 如果 charRight >= 128 或者 requiredCounts[charRight] == 0，它不是我们需要追踪的字符，直接忽略

		// 检查当前窗口是否包含了 t 的所有字符（包括重复）
		for foundCount == tLen {
			// 窗口 [left ... right] 是一个有效的覆盖 t 的子串

			// 1. 更新最小窗口记录（如果当前窗口更小）
			currentLen := right - left + 1
			if currentLen < minLen {
				minLen = currentLen
				minStart = left // minStart 记录的是有效窗口的起始位置
			}

			// 2. 尝试缩小窗口，移动 left 指针
			charLeft := s[left] // 当前离开窗口的字符

			// 检查 charLeft 是否是我们关心的字符 (在数组索引范围内)
			if charLeft < 128 {
				// 更新当前窗口中该字符的计数（它要离开了）
				currentCounts[charLeft]--

				// 如果在移除该字符后，当前窗口中该字符的数量小于 t 中所需的数量
				// 说明我们刚刚移除的是一个必需的字符实例，需要减少 foundCount
				// 注意这里的判断是 currentCounts[charLeft] < requiredCounts[charLeft]，
				// 因为 currentCounts 已经减 1 了。
				if requiredCounts[charLeft] > 0 && currentCounts[charLeft] < requiredCounts[charLeft] {
					foundCount--
				}
			}
			// 如果 charLeft >= 128，它不是我们需要追踪的字符，移除它不影响 foundCount

			// 3. 移动 left 指针，缩小窗口
			left++

			// 内层循环会继续执行，直到 foundCount != tLen（窗口不再有效）
		}

		// 移动 right 指针，扩展窗口
		right++
	}

	// 外层循环结束后，检查是否找到了有效的窗口
	if minLen > sLen { // 如果 minLen 还是初始值，说明没有找到覆盖 t 的子串
		return ""
	}

	// 返回找到的最小窗口子串
	return s[minStart : minStart+minLen]
}

// @lc code=end

/*
// @lcpr case=start
// "ADOBECODEBANC"\n"ABC"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"a"\n
// @lcpr case=end

// @lcpr case=start
// "a"\n"aa"\n
// @lcpr case=end

*/
