/*
 * @lc app=leetcode.cn id=2 lang=golang
 * @lcpr version=30204
 *
 * [2] 两数相加
 *
 * https://leetcode.cn/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (45.18%)
 * Likes:    11171
 * Dislikes: 0
 * Total Accepted:    2.4M
 * Total Submissions: 5.2M
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
 *
 * 请你将两个数相加，并以相同形式返回一个表示和的链表。
 *
 * 你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 *
 *
 * 示例 1：
 *
 * 输入：l1 = [2,4,3], l2 = [5,6,4]
 * 输出：[7,0,8]
 * 解释：342 + 465 = 807.
 *
 *
 * 示例 2：
 *
 * 输入：l1 = [0], l2 = [0]
 * 输出：[0]
 *
 *
 * 示例 3：
 *
 * 输入：l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
 * 输出：[8,9,9,9,0,0,0,1]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 每个链表中的节点数在范围 [1, 100] 内
 * 0 <= Node.val <= 9
 * 题目数据保证列表表示的数字不含前导零
 *
 *
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
// func reverseList(head *ListNode) *ListNode {
// 	var prev *ListNode
// 	cur := head
// 	for cur != nil {
// 		Next := cur.Next
// 		cur.Next = prev
// 		prev = cur
// 		cur = Next
// 	}
// 	return prev
// }
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//惯例，先写暴力解法
	//不需要反转的一集
	//思路，将链表从前到后慢慢相加并把进位送到下一次循环中，当 l1 和 l2 都为 nil 且进位也不存在时终止循环，头指针不保存数组，只作为链表起始点的记录者，curr 指针指向当前结点，我们每一次运算得到的都是下一个结点的值，要用:&node{}或者 var prev *node 创建 nil 的新节点，new 会默认置 0
	dummy := &ListNode{}
	curr := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}
	return dummy.Next
	// for l1 != nil || l2 != nil {
	// 	if l1.Val+l2.Val >= 10 {
	// 		ans.Val += l1.Val + l2.Val - 10
	// 		Next := ans.Next
	// 		Next.Val = 1
	// 		ans = Next
	// 	} else {
	// 		ans.Val = l1.Val + l2.Val
	// 		Next := ans.Next
	// 		ans = Next
	// 	}
	// 	l1 = l1.Next
	// 	l2 = l2.Next
	// }
	// return reverseList(head)
}

// @lc code=end

/*
// @lcpr case=start
// [2,4,3]\n[5,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [0]\n[0]\n
// @lcpr case=end

// @lcpr case=start
// [9,9,9,9,9,9,9]\n[9,9,9,9]\n
// @lcpr case=end

*/
