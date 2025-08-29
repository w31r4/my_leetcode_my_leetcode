/*
 * @lc app=leetcode.cn id=148 lang=golang
 * @lcpr version=30204
 *
 * [148] 排序链表
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.
**/

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//

	// 2. 使用快慢指针找到链表中点
	var prev *ListNode // 用来记录 slow 的前一个节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 3. 断开链表
	// 循环结束后，slow 是后半段的头，prev 是前半段的尾
	prev.Next = nil

	left := sortList(head)
	right := sortList(slow)
	return mergeList1(left, right)
}
func mergeList1(leftNode, rightNode *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for rightNode != nil && leftNode != nil {
		if leftNode.Val < rightNode.Val {
			tail.Next = leftNode
			leftNode = leftNode.Next
		} else {
			tail.Next = rightNode
			rightNode = rightNode.Next
		}
		tail = tail.Next
	}
	if rightNode == nil {
		tail.Next = leftNode
	} else {
		tail.Next = rightNode
	}
	return dummy.Next
}

// @lc code=end

/*
// @lcpr case=start
// [4,2,1,3]\n
// @lcpr case=end

// @lcpr case=start
// [-1,5,3,4,0]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
