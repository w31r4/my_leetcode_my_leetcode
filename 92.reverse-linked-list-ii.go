/*
 * @lc app=leetcode.cn id=92 lang=golang
 * @lcpr version=30204
 *
 * [92] 反转链表 II
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
/**
 * Definition for singly-linked list.

 */
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	i := 0
	var rightNode, leftNode, lastNode, beforeLeft *ListNode
	isFirstTime := true

	for curr := head; curr != nil; {
		if !isFirstTime {
			//当非首次循环时，将 curr 向后移动一位
			lastNode = curr
			curr = curr.Next
		}
		i++
		if i == left {

			beforeLeft = lastNode
			leftNode = curr
		}
		for i > left && i < right {
			tmp := curr.Next
			curr.Next = lastNode
			lastNode = curr
			curr = tmp
			i++
		}
		if i == right {
			rightNode = curr
			afterRight := curr.Next
			curr.Next = lastNode
			if beforeLeft == nil {
				head = rightNode
			} else {
				beforeLeft.Next = rightNode
			}
			leftNode.Next = afterRight
			break
		}
		isFirstTime = false
	}
	return head
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n4\n
// @lcpr case=end

// @lcpr case=start
// [5]\n1\n1\n
// @lcpr case=end

*/
