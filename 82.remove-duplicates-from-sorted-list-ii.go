/*
 * @lc app=leetcode.cn id=82 lang=golang
 * @lcpr version=30204
 *
 * [82] 删除排序链表中的重复元素 II
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

// 这题需要一个计数器，记住当前数字出现了多少次，也需要一个记录节点记录当前数字的开头节点，大于一的话全部跳过，当当前节点只出现过一次时将其加入新链表中，并更新链表尾
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	newTail := dummyHead
	//用于记录上一个被保存的节点，是构建的新链表的尾链表
	isFirstTime := true
	//判断当前节点中的值是否只出现过一次
	for curr := head; curr != nil; {
		nowVal := curr.Val
		nowHead := curr
		curr = curr.Next
		for curr != nil && nowVal == curr.Val {
			isFirstTime = false
			curr = curr.Next
			//提前向后移动一位
		}
		if !isFirstTime {
			newTail.Next = curr
			//如果当前值出现过不止一次，那么将新链表的尾指针的 next 指向下一个 val 不同的链表节点
			isFirstTime = true
		} else {
			newTail = nowHead
			//如果仅出现过一次，此时尾指针 next 已经指向当前 nowHead，将尾指针赋值为当前 nowHead 完成尾插法
		}
	}
	return dummyHead.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,3,4,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [1,1,1,2,3]\n
// @lcpr case=end

*/
