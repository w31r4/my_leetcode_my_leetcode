/*
 * @lc app=leetcode.cn id=61 lang=golang
 * @lcpr version=30204
 *
 * [61] 旋转链表
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

// 似乎是简单题简单做
// 首先获取当前链表有多少个元素
// 用元素个数对 k 取模

func rotateRight(head *ListNode, k int) *ListNode {
	// 第一次遍历：目的是计算链表的总长度 length
	length := 0
	// 注意：这里的 curr 在计算完长度后，其状态（指向 nil）不会影响后续对 curr 的重新赋值和使用
	for curr := head; curr != nil; { // 从头节点开始遍历
		length++         // 每经过一个节点，长度加一
		curr = curr.Next // curr 指向下一个节点
	}

	// 处理特殊情况：如果链表为空、只有一个节点，或者 k 为 0 (无需旋转)，
	// 或者 k 是链表长度的整数倍 (旋转 k 次后等于没有旋转)，则直接返回原头节点。
	if length <= 1 || k == 0 || k%length == 0 {
		return head
	}

	// 计算实际需要从尾部移动到头部的节点数量。
	// 例如：length=5, k=2, last=2 (移动末尾 2 个)
	// 例如：length=5, k=7, last=7%5=2 (同样移动末尾 2 个)
	last := k % length
	// (用户原有注释) k 对 length 取模
	// (用户原有注释) 获取我们需要将倒数第几个开始的链表挪到前面去

	// 计算新链表的尾部是原链表的第几个节点（从 1 开始计数）。
	// 这个节点之后的部分将被移动到链表头部。
	// 例如：length=5, last=2 (移动末尾 2 个), 则 front=3。意味着原链表的前 3 个节点将构成旋转后链表的后半部分，
	// 第 3 个节点将成为新的尾节点。
	front := length - last
	// (用户原有注释) 获取我们需要将正数第几个开始的链表挪到前面去 (应理解为前 front 个节点构成第一段)

	// 创建哑节点，简化对链表头部的操作
	dummyHead := &ListNode{Next: head}
	// 保存原始头节点，因为旋转后，原链表的尾部会指向它
	oldHead := dummyHead.Next // oldHead 即为原始的 head

	// curr 指针重新用于定位分割点（即新的尾节点）
	curr := dummyHead // 从哑节点开始，方便计数和操作
	// 遍历 front 次，使得 curr 指向新的尾节点（即原链表的第 front 个节点）
	for i := 1; i <= front; i++ {
		curr = curr.Next
	}
	// 循环结束后，curr 指向的是旋转后链表的第一部分的尾节点 (new tail of the first part)

	// newHead 是旋转后链表的真正头节点，它是 curr 指向节点的下一个节点
	newHead := curr.Next
	// newTail 初始指向 newHead，它将作为迭代器找到旋转部分的真正尾部
	newTail := newHead
	// 断开链表：将第一部分的尾节点 curr 的 Next 指向 nil，形成独立的链表段
	curr.Next = nil

	// 遍历找到旋转到前面的那部分链表（从 newHead 开始）的尾节点
	for newTail.Next != nil {
		newTail = newTail.Next
	}
	// 循环结束后，newTail 指向原链表的最后一个节点，也就是旋转部分的尾节点

	// 将旋转部分的尾节点 (newTail) 连接到原始链表的头节点 (oldHead)
	newTail.Next = oldHead

	// 更新哑节点的 Next 指向新的头节点，构成完整的旋转后链表
	dummyHead.Next = newHead
	// 返回旋转后链表的头节点
	return dummyHead.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2]\n4\n
// @lcpr case=end

*/
