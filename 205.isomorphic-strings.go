/*
 * @lc app=leetcode.cn id=205 lang=golang
 * @lcpr version=30204
 *
 * [205] 同构字符串
 *
 * https://leetcode.cn/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (49.95%)
 * Likes:    773
 * Dislikes: 0
 * Total Accepted:    341.7K
 * Total Submissions: 682.7K
 * Testcase Example:  '"egg"\n"add"'
 *
 * 给定两个字符串 s 和 t ，判断它们是否是同构的。
 *
 * 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
 *
 *
 * 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。
 *
 *
 *
 * 示例 1:
 *
 * 输入：s = "egg", t = "add"
 * 输出：true
 *
 *
 * 示例 2：
 *
 * 输入：s = "foo", t = "bar"
 * 输出：false
 *
 * 示例 3：
 *
 * 输入：s = "paper", t = "title"
 * 输出：true
 *
 *
 *
 * 提示：
 *
 *
 *
 *
 * 1 <= s.length <= 5 * 10^4
 * t.length == s.length
 * s 和 t 由任意有效的 ASCII 字符组成
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 构建规则是 target 中第一个字符标记为 0，后续遇到不同字符则标记++,建立次序哈希，记录该字符代表的标记
// 需要建立一个频率哈希表，记录字符频率到其对应的次序中
// 核心逻辑错误，这样的逻辑会让 abca，和 xxzy 对应上
// 应该建立双向表，当没有对应关系时建立对应关系，有对应关系时对比是否符合对应关系
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sToTMap := [128]int{}
	tToSMap := [128]int{}
	for i := 0; i < len(s); i++ {
		tPlusOne := t[i] + 1
		sPlusOne := s[i] + 1
		//防止 ascii 值等于 0 干扰判断
		isSAndTAlreadyMapped := sToTMap[s[i]] != 0 && tToSMap[t[i]] != 0
		isSAndTBothNotMapped := sToTMap[s[i]] == 0 && tToSMap[t[i]] == 0
		if isSAndTBothNotMapped {
			//当双方都没有建立关系时，双方同时建立关系
			sToTMap[s[i]] = int(tPlusOne)
			tToSMap[t[i]] = int(sPlusOne)
			//建立了双向关系
			continue
		} else if isSAndTAlreadyMapped {
			//当双方同时都建立了关系时进入判断
			if sToTMap[s[i]] != int(t[i]+1) {
				//当双方有一方建立的关系不是另一方时，返回 false
				return false
			}
			continue
		}
		//除此之外的情况都返回 false
		//也就是一个匹配了另一个没匹配
		return false
	}
	return true
}

//优雅解法：

// isIsomorphic 检查字串 s 和 t 是否同构。
// 同构定义：s 中的字元可以被替换得到 t，同时保持字元的顺序。
// 限制条件：
//  1. s 中所有相同的字元必须映射到 t 中相同的字元。
//  2. s 中不同的字元不能映射到 t 中相同的字元。
//     (即 s->t 和 t->s 的映射都必须是明确的，不能一对多或多对一混用)
// func isIsomorphic(s string, t string) bool {
// 	// 基本检查：长度不同，肯定不是同构
// 	if len(s) != len(t) {
// 		return false
// 	}

// 	// 使用阵列模拟哈希表，空间复杂度 O(C)，C 为字元集大小 (ASCII 为 128 或 256)
// 	// sToTMap: 记录 s 字元到 t 字元的映射关系 (s[i] -> t[i])
// 	sToTMap := [128]byte{} // 阵列元素预设为 0 (byte 的零值)
// 	// tIsMapped: 记录 t 中的字元是否已经被用作映射目标
// 	tIsMapped := [128]bool{} // 阵列元素预设为 false

// 	for i := 0; i < len(s); i++ {
// 		sChar := s[i]     // 当前 s 中的字元 (byte)
// 		tChar := t[i] + 1 // 当前 t 中的字元 (byte)
// 		//防止 0x00 干扰判断

		// 检查 sChar 是否已经有映射记录
		mappedTChar, sCharHasMapped := sToTMap[sChar], sToTMap[sChar] != 0 // Go 没有内建判断 key 是否存在的 array 语法，用值是否为 0 判断
		// 如果 sChar 已映射，但目标不是当前的 tChar，则违反规则 1
		if sCharHasMapped && mappedTChar != tChar {
			return false
		}

		// 检查 tChar 是否已经被其他 s 字元映射过了
		tCharIsAlreadyTarget := tIsMapped[tChar]
		// 如果 sChar 尚未映射，但 tChar 已经被别的 sChar 映射了，则违反规则 2
		if !sCharHasMapped && tCharIsAlreadyTarget {
			return false
		}

// 		// 如果上述检查都通过，且 sChar 之前未映射，则建立新的双向关联记录
// 		if !sCharHasMapped {
// 			sToTMap[sChar] = tChar  // 记录 sChar -> tChar 的映射
// 			tIsMapped[tChar] = true // 标记 tChar 已被用作映射目标
// 		}
// 		// 如果 sChar 之前已映射且映射正确（第一个 if 条件已处理），则无需操作，继续循环
// 	}

// 	// 如果循环正常结束，说明所有字元都满足同构条件
// 	return true
// }

// if len(s) != len(t) {

// 	return false
// }
// seqHash := [128]int{}
// freqHash := [129]int{}
// nowSeq := 1
// for _, val := range t {
// 	if seqHash[val] == 0 {
// 		//当前字母没有出现过
// 		seqHash[val] = nowSeq
// 		//记录当前字母是第几个出现的不同字母
// 		freqHash[seqHash[val]]++
// 		nowSeq++
// 		continue
// 		//给出现次数 +1
// 	}
// 	if seqHash[val] != 0 {
// 		//当前字母出现过
// 		freqHash[seqHash[val]]++
// 		//给出现次数 +1
// 	}
// }
// nowSeq = 1
// seqHash = [128]int{}
// for _, val := range s {
// 	if seqHash[val] == 0 {
// 		//当前字符没有出现过

// 		seqHash[val] = nowSeq
// 		//记录当前字母是第几个出现的不同字母
// 		freqHash[seqHash[val]]--
// 		if freqHash[seqHash[val]] < 0 {
// 			return false
// 		}
// 		nowSeq++
// 	}
// 	if seqHash[val] != 0 {
// 		//当前字母出现过
// 		freqHash[seqHash[val]]--
// 		if freqHash[seqHash[val]] < 0 {
// 			return false
// 		}
// 		//哈希表--
// 	}
// }
// return true

// @lc code=end

/*
// @lcpr case=start
// "egg"\n"add"\n
// @lcpr case=end

// @lcpr case=start
// "foo"\n"bar"\n
// @lcpr case=end

// @lcpr case=start
// "paper"\n"title"\n
// @lcpr case=end

*/
