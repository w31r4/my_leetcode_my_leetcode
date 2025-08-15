/*
 * @lc app=leetcode.cn id=25 lang=golang
 * @lcpr version=30204
 *
 * [25] K 个一组翻转链表
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

func reverseList(list *ListNode) (listEnd, listStart *ListNode) {
	var prev *ListNode
	listEnd = list
	for curr := list; curr != nil; {
		tmpNext := curr.Next
		curr.Next = prev
		prev = curr
		curr = tmpNext
	}
	listStart = prev
	return
}

// 思路：创建 dummyNode，记录每段循环中的第一个和最后一个节点
// 先进行 k 次循环，判断剩余链表是否需要反转，如果不需要将前一个反转后的链表尾指向当前链表头
//如果需要，先反转，后将前一个链表链表尾指向当前链表头
// 在这道题目中，我们需要使用到当前链表头，当前链表尾，反转后链表头，反转后链表尾，以及前一个链表尾
// 连接是将前一个链表尾指向当前反转或未反转链表头
// 链表头和反转后链表头都用 listStart 作为变量名

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	//哨兵链表头
	dummyHead := &ListNode{}
	dummyHead.Next = head
	//当前指向链表元素
	curr := head

	var listStart, listEnd, tmpEnd *ListNode
	//上一个反转后链表的结尾处
	var prevEnd *ListNode
	prevEnd = dummyHead

	for curr != nil {
		//开始循环，保存未反转前的链表头
		listStart = curr
		for i := 1; i <= k; i++ {
			if curr == nil {
				//当剩余链表长度小于 k 个时，将上一个反转后链表尾指向当前链表头
				prevEnd.Next = listStart
				//直接返回，break 只能脱离当前循环
				return dummyHead.Next
			}
			//当循环到第 k 个节点还没有结束循环时
			if i == k {
				//保存未反转前的链表尾
				listEnd = curr
			}
			//移动到下一个需要做判断的链表元素
			curr = curr.Next
		}
		//将当前链表截断，为反转做准备
		listEnd.Next = nil
		//将链表头送入反转函数
		tmpEnd, listStart = reverseList(listStart)
		//获得反转后的链表头和链表尾
		prevEnd.Next = listStart
		//前一个反转后的链表尾要指向反转后的链表头才行
		prevEnd = tmpEnd
		//更新前一个反转的链表尾
		// tmpEnd.Next = curr
		//当前反战后链表尾要指向下一段链表的当前开端，如果不反转的话可以直接返回

	}
	return dummyHead.Next
}

// func reverse(head *ListNode) *ListNode {
// 	var prev *ListNode
// 	curr := head
// 	for curr != nil {
// 		next := curr.Next
// 		curr.Next = prev
// 		prev = curr
// 		curr = next
// 	}
// 	return prev
// }

// func reverseKGroup(head *ListNode, k int) *ListNode {
// 	if head == nil || k == 1 {
// 		return head
// 	}

// 	dummy := &ListNode{Next: head}
// 	prevGroupEnd := dummy // 指向上一组反转后的尾节点，初始为 dummy

// 	for {
// 		// 1. 找到当前 K 个节点的组的头和尾
// 		groupStart := prevGroupEnd.Next
// 		groupEnd := prevGroupEnd
// 		// 尝试向前走 K 步，找到 groupEnd
// 		for i := 0; i < k; i++ {
// 			groupEnd = groupEnd.Next
// 			if groupEnd == nil { // 不足 K 个节点
// 				return dummy.Next // 剩余部分不反转，直接返回
// 			}
// 		}

// 		// 2. 记录下一组的开始节点
// 		nextGroupStart := groupEnd.Next

// 		// 3. 断开当前组与链表其余部分的连接，以便反转
// 		groupEnd.Next = nil // 非常重要！使得反转操作只在当前 K 个节点内进行

// 		// 4. 反转当前 K 个节点的组
// 		// groupStart 是当前组的头，反转后它将成为尾
// 		// 调用 reverse 函数反转从 groupStart 开始的子链表
// 		// 返回的新头是原来的 groupEnd
// 		reversedGroupHead := reverse(groupStart) // 反转后，原来的 groupStart 变成了尾，原来的 groupEnd 变成了头

// 		// 5. 连接：将反转后的组重新接回链表
// 		prevGroupEnd.Next = reversedGroupHead // 上一组的尾 指向 当前反转组的新头 (即原来的 groupEnd)
// 		groupStart.Next = nextGroupStart      // 当前反转组的新尾 (即原来的 groupStart) 指向 下一组的头

// 		// 6. 为下一轮迭代做准备
// 		prevGroupEnd = groupStart // prevGroupEnd 更新为当前反转组的尾部 (即原来的 groupStart)

// 		// 如果 nextGroupStart 为 nil，说明没有更多节点了
// 		if nextGroupStart == nil {
// 			break
// 		}
// 	}

// 	return dummy.Next
// }

// }

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n3\n
// @lcpr case=end

*/
