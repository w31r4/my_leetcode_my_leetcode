/*
 * @lc app=leetcode.cn id=21 lang=golang
 * @lcpr version=30204
 *
 * [21] 合并两个有序链表
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2 // 如果 list1 为空，无论 list2 是否为空，都返回 list2
	}
	if list2 == nil {
		return list1 // 如果 list2 为空 (此时 list1 必然不为空)，返回 list1
	}
	var head *ListNode
	// head := &ListNode{}

	// if list1.Val > list2.Val {
	// 	head = list2
	// 	list2 = list2.Next
	// } else {
	// 	head = list1
	// 	list1 = list1.Next
	// }
	tail := head
	for !(list1 == nil || list2 == nil) {
		if list1.Val > list2.Val {
			// 当 list2 中 val 更小时
			tail.Next = list2

			list2 = list2.Next
		} else {
			tail.Next = list1

			list1 = list1.Next
		}
		tail = tail.Next
	}
	if list1 == nil {
		tail.Next = list2
	} else if list2 == nil {
		tail.Next = list1
	}
	return head.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,4]\n[1,3,4]\n
// @lcpr case=end

// @lcpr case=start
// []\n[]\n
// @lcpr case=end

// @lcpr case=start
// []\n[0]\n
// @lcpr case=end

*/
