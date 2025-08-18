/*
 * @lc app=leetcode.cn id=433 lang=golang
 * @lcpr version=30204
 *
 * [433] 最小基因变化
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// type Gene struct {
// 	Diff    int
// 	Content string
// }

// func newGeneFunc(diff int, content string) *Gene {
// 	return &Gene{Diff: diff, Content: content}
// }

func isOneDiff(gen1, gen2 string) bool {
	diffs := 0
	count := len(gen1)
	for i := 0; i < count; i++ {
		if gen1[i] != gen2[i] {
			diffs++
		}
	}
	if diffs == 1 {
		return true
	}
	return false
}

func minMutation(startGene string, endGene string, bank []string) int {
	// ✅ 用 BFS，从起始基因开始
	// ✅ 队列存储（基因序列，步数）
	// ✅ 每次取出基因时检查是否为目标
	// ✅ 用 visited 数组避免重复访问
	// ✅ 队列为空时返回 -1
	length := len(bank)
	visited := make([]bool, length)
	geneQueue := []string{}
	geneQueue = append(geneQueue, startGene)
	steps := 0
	for geneQueue != nil {
		temp := geneQueue
		geneQueue = nil
		for i := 0; i < len(temp); i++ {
			if temp[i] == endGene {
				return steps
			}
			for j := 0; j < length; j++ {
				if !visited[j] && isOneDiff(temp[i], bank[j]) {
					geneQueue = append(geneQueue, bank[j])
					visited[j] = true
				}
			}
		}
		steps++
	}
	return -1
}

// @lc code=end

/*
// @lcpr case=start
// "AACCGGTT"\n"AACCGGTA"\n["AACCGGTA"]\n
// @lcpr case=end

// @lcpr case=start
// "AACCGGTT"\n"AAACGGTA"\n["AACCGGTA","AACCGCTA","AAACGGTA"]\n
// @lcpr case=end

// @lcpr case=start
// "AAAAACCC"\n"AACCCCCC"\n["AAAACCCC","AAACCCCC","AACCCCCC"]\n
// @lcpr case=end

*/
