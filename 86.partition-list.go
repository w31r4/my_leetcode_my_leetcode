/*
 * @lc app=leetcode.cn id=86 lang=golang
 * @lcpr version=30204
 *
 * [86] 分隔链表
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 思路：维护两个指针，一个指向较大元素，一个指向较小元素，遍历完后拼接
// partition 函数根据给定值 x 将链表分隔成两部分：
// 所有小于 x 的节点排在前面，所有大于或等于 x 的节点排在后面。
// 每个部分中节点的原始相对顺序应被保留。
// 思路：维护两个指针，一个指向较大元素，一个指向较小元素，遍历完后拼接
func partition(head *ListNode, x int) *ListNode {
	// 处理边界情况：如果链表为空，直接返回。
	if head == nil {
		return head
	}
	//并不需要处理边界情况
	// dummyHead 用于构建“小于 x”的节点链表 (smaller list)。
	// 它的 Next 指向原链表的头，简化对 smaller list 头部的操作。
	dummyHead := &ListNode{Next: head}
	// smallerCurr 是 smaller list 的当前尾指针。
	smallerCurr := dummyHead

	// isFirstTime 标志用于处理“大于等于 x”的节点链表 (bigger list) 的第一个节点。
	isFirstTime := true
	// biggerHead 是 bigger list 的头指针。
	var biggerHead *ListNode
	// biggerCurr 是 bigger list 的当前尾指针。
	var biggerCurr *ListNode

	// 遍历原始链表
	for curr := head; curr != nil; curr = curr.Next {
		if curr.Val >= x {
			// 当前节点的值大于或等于 x，将其加入 bigger list
			if isFirstTime {
				// 如果是 bigger list 的第一个节点
				biggerCurr = curr   // bigger list 的尾指针指向当前节点
				biggerHead = curr   // bigger list 的头指针也指向当前节点
				isFirstTime = false // 不再是第一次了
			} else {
				// 如果 bigger list 已有节点，则将当前节点追加到其尾部
				biggerCurr.Next = curr // 上一个大值节点的 Next 指向当前节点
				biggerCurr = curr      // 更新 bigger list 的尾指针为当前节点
			}
		} else {
			// 当前节点的值小于 x，将其加入 smaller list
			smallerCurr.Next = curr // smaller list 的尾指针的 Next 指向当前节点
			smallerCurr = curr      // 更新 smaller list 的尾指针为当前节点
		}
	}

	// 遍历结束后，需要正确设置两个子链表的尾部。
	// bigger list 的尾部需要指向 nil，以防形成环或连接到不期望的节点。
	if biggerCurr != nil { // 检查 bigger list 是否为空
		biggerCurr.Next = nil
	}

	// 将 smaller list 的尾部连接到 bigger list 的头部。
	smallerCurr.Next = biggerHead

	// 返回 dummyHead 的下一个节点，即分隔后链表的真正头节点。
	return dummyHead.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,4,3,2,5,2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [2,1]\n2\n
// @lcpr case=end

*/
