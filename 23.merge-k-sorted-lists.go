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
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	dyingListHead := lists[0]
	mergedListHead := lists[1]
	dummy := &ListNode{}
	tail := dummy
	for mergedListHead != nil && dyingListHead != nil {
		if dyingListHead.Val > mergedListHead.Val {
			tail.Next = mergedListHead
			mergedListHead = mergedListHead.Next
		} else {
			tail.Next = dyingListHead
			dyingListHead = dyingListHead.Next
		}
		tail = tail.Next
	}
	if mergedListHead == nil {
		tail.Next = dyingListHead
	} else {
		tail.Next = mergedListHead
	}
	lists[1] = dummy.Next
	return mergeKLists(lists[1:])
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
