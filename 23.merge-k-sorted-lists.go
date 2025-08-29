/*
 * @lc app=leetcode.cn id=23 lang=golang
 * @lcpr version=30204
 *
 * [23] 合并 K 个升序链表
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
func mergeKLists(lists []*ListNode) *ListNode {
	// 这题最简单的思路就是单纯的的遍历
	//简单遍历是线性提升时间复杂度 N*k，一次只能排序一个
	//可以二分分治让查找更加均匀，减少排序次数
	//一次排序两个链表
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	length := len(lists)
	halfLength := length / 2
	leftLists := lists[:halfLength]
	rightLists := lists[halfLength:]
	return mergeList(mergeKLists(leftLists), mergeKLists(rightLists))
}
func mergeList(node1, node2 *ListNode) *ListNode {

	dummy := &ListNode{}
	tail := dummy
	for node2 != nil && node1 != nil {
		if node1.Val > node2.Val {
			tail.Next = node2
			node2 = node2.Next
		} else {
			tail.Next = node1
			node1 = node1.Next
		}
		tail = tail.Next
	}
	if node2 == nil {
		tail.Next = node1
	} else {
		tail.Next = node2
	}
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [[1,4,5],[1,3,4],[2,6]]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

// @lcpr case=start
// [[]]\n
// @lcpr case=end

*/
