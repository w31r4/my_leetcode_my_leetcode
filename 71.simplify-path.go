/*
 * @lc app=leetcode.cn id=71 lang=golang
 * @lcpr version=30204
 *
 * [71] 简化路径
 */

// @lcpr-template-start
package leetcode

import "strings"

// @lcpr-template-end
// @lc code=start
func simplifyPath(path string) string {
	pathComponents := strings.FieldsFunc(path, func(r rune) bool { return r == '/' })
	stack := []string{}
	for _, adr := range pathComponents {
		switch adr {
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case ".":
			continue
		default:
			stack = append(stack, adr)
		}
	}
	if len(stack) == 0 {
		return "/"
	}
	var builder strings.Builder
	for _, val := range stack {
		builder.WriteString("/")
		builder.WriteString(val)
	}
	return builder.String()
}

// @lc code=end

/*
// @lcpr case=start
// "/home/"\n
// @lcpr case=end

// @lcpr case=start
// "/home//foo/"\n
// @lcpr case=end

// @lcpr case=start
// "/home/user/Documents/../Pictures"\n
// @lcpr case=end

// @lcpr case=start
// "/../"\n
// @lcpr case=end

// @lcpr case=start
// "/.../a/../b/c/../d/./"\n
// @lcpr case=end

*/
