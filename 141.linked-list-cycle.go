/*
 * @lc app=leetcode.cn id=141 lang=golang
 * @lcpr version=30204
 *
 * [141] 环形链表
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

// 怎么判断链表是否为环？
// 指针哈希表？
func hasCycle(head *ListNode) bool {
	//方法一：哈希表存储
	// if head == nil {
	// 	return false
	// }
	// listMap := make(map[*ListNode]bool)
	// for !listMap[head] {
	// 	//当 listMap 中没出现过当前链表中的元素时
	// 	listMap[head] = true
	// 	head = head.Next
	// 	if head == nil {
	// 		return false
	// 	}
	// }
	// return true
	//方法二：快慢指针
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,0,-4]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n0\n
// @lcpr case=end

// @lcpr case=start
// [1]\n-1\n
// @lcpr case=end

*/
