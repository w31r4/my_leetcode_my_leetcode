/*
 * @lc app=leetcode.cn id=28 lang=golang
 * @lcpr version=30204
 *
 * [28] 找出字符串中第一个匹配项的下标
 *
 * https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/description/
 *
 * algorithms
 * Easy (44.69%)
 * Likes:    2398
 * Dislikes: 0
 * Total Accepted:    1.3M
 * Total Submissions: 2.8M
 * Testcase Example:  '"sadbutsad"\n"sad"'
 *
 * 给你两个字符串 haystack 和 needle，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0
 * 开始）。如果 needle 不是 haystack 的一部分，则返回  -1。
 *
 *
 *
 * 示例 1：
 *
 * 输入：haystack = "sadbutsad", needle = "sad"
 * 输出：0
 * 解释："sad" 在下标 0 和 6 处匹配。
 * 第一个匹配项的下标是 0，所以返回 0。
 *
 *
 * 示例 2：
 *
 * 输入：haystack = "leetcode", needle = "leeto"
 * 输出：-1
 * 解释："leeto" 没有在 "leetcode" 中出现，所以返回 -1。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= haystack.length, needle.length <= 10^4
 * haystack 和 needle 仅由小写英文字符组成
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func strStr(haystack string, needle string) int {
	//调库函数一行解决
	// return strings.Index(haystack, needle)
	//但是需要知道的还有很多
	//如 KMP 算法
	//正常匹配算法是从目标串和模式串第一位开始，如果遇到不匹配，则将模式串向后移动一位继续匹配，直到匹配为止，时间复杂度 O n*m
	//以上算法还是为了防止匹配好的字符串里的后缀有符合模式串前缀的存在被忽略，一个一个匹配不会有忽略
	//而 kmp 算法可以记录已经匹配好的目标串的后缀里有哪些能和当前模式串的前缀吻合的地方，直接跳过来，让每次匹配都为下一次匹配保存了一定的信息
	//KMP 算法需要求出前缀表，模式串中每个位置都有自己的前缀表的值
	//前缀表是什么：以 aabac 为例，它的前缀表数组为:0,1,0,1,0
	//也就是匹配 aacbcaabac 时，当匹配完 aac，发现 c 不匹配，之后，我们可以读取 c 前的前缀表，为 1，也就是从 c 向前，有一位已经匹配好了，因为前缀表就是记录模式串当前位置最多相同前后缀的情况
	//如模式串 aabaad，当匹配完 aabaa 后，如果出现不匹配的情况，不匹配的位置前两位已经匹配完了，也就是 aa，只要从 aa 后开始和匹配即可，如果还是不匹配，那么 aa 的前缀表值为 1，如果目标串为 aabaacaabaa，先从
	//aabaa d!=c aa b!=c a a!=c ,a!=c 跳过 c
	//那么如何求取前缀表呢？
	//这也很巧妙：初始化双指针 i,j,i 代表当前求取最大相同前后缀值的位置，j 代表最大相同前后缀的值
	//初始化 i=1,j=0
	//先比较 prefix[i] 和 prefix[j]
	//相等的话 j++,把 j 赋值给 prefix[i]
	//不想当的话回退到 prefix[j-1]
	//那么 prefix[j-1] = m 意味着什么呢？
	// 意味着在前 j-1 个字符中，前 m 个和后 m 个字符形成的前后缀是一样的
	//在匹配时，后缀一定等于后缀，当相同的后缀后面一位字符不匹配了，就而此时前缀又等于之前的后缀，我们就可以比较：相同的前缀 + 后一个值<->和前缀相同的当前后缀 + 后一个值

	j := 0
	prefix := make([]int, len(needle))
	//prefix 中 j 位置的值 n 代表从 j 截止前 n 位和后 n 位前后缀相同
	for i := 1; i < len(needle); i++ {
		// 不匹配时回退到 prefix[j-1]
		for j > 0 && needle[i] != needle[j] {
			j = prefix[j-1]
			//在重复串中 prefix[j-1]>0 时，说明出现了重复
			//prefix[n]==0 的部分为重复子串
			//那么 prefix[j-1] = m 意味着什么呢？
			// 意味着在前 j-1 个字符中，前 m 个和后 m 个字符形成的前后缀是一样的
			//在匹配时，此时匹配的前缀一定等于后缀，当相同的两个前后缀后面一位字符不匹配了，而此时 prefix[j-1]=m，也就是前缀的前 m 个字符前缀等于前缀的后 m 个字符后缀，而后缀的后 m 个字符后缀是等于前缀的后 m 个字符后缀相当于等于前缀的前 m 个字符前缀，我们就可以比较：相同的前缀 + 后一个值<->和前缀相同的当前后缀 + 后一个值
			// 比较 s[j] 是否等于 s[i]，若相等则找到了最长相等前后缀
			// 若仍然不等，则继续前移，直到前移到开头元素，说明没有相等的前后缀
		}
		if needle[i] == needle[j] {
			j++
		}
		prefix[i] = j
	}
	n := len(needle)
	if n == 0 {
		return 0
	}
	j = 0
	//j 是当前模式串和目标串的匹配字符个数
	for i := 0; i < len(haystack); i++ {
		for j > 0 && haystack[i] != needle[j] {
			j = prefix[j-1]
			/*
				当 j>0 且此时 haystack[i] 和模式串的 j 号字符不匹配时
				查询前缀表，将 j 重新赋值为 prefix[j-1]
				那么 prefix[j-1] = m 意味着什么呢？
				意味着在前 j-1 个字符中，前 m 个和后 m 个字符形成的前后缀是一样的
				我们将模式串重新从前 j-1 个字符的最大相同前缀开始重新匹配
				此时目标串的前 j-1 个字符和模式串的前 j-1 个字符是一致的，从抽象来看我们不需要管目标串
				只管模式串就行
				直到确实无法匹配，将 j 清零，并走向目标串的下一位
			*/
		}
		if haystack[i] == needle[j] {
			//如果当前目标串的字符值和模式串的字符值相同，那么目标串和模式串同时++
			j++
		}
		if j == n {
			//直到模式串全部匹配完毕为止
			return i - n + 1
		}
	}
	return -1
}

// @lc code=end

/*
// @lcpr case=start
// "sadbutsad"\n"sad"\n
// @lcpr case=end

// @lcpr case=start
// "leetcode"\n"leeto"\n
// @lcpr case=end

*/
