/*
 * @lc app=leetcode.cn id=19 lang=golang
 * @lcpr version=30204
 *
 * [19] 删除链表的倒数第 N 个结点
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

// 链表没有索引，需要知道现在是第几个链表需要在遍历循环中单独创建索引变量
// 如果需要知道是倒数第几个就更需要知道一共有多少元素
// 如果是需要一次扫描实现那么就
// 需要建立哈希表存储遍历过的链表元素
// 否则无法回退至目标链表元素

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// if head.Next == nil {
	// 	return nil
	// }
	// //当链表中只有一个元素时，直接返回 nil
	// listIndexMap := make(map[int]*ListNode)
	// //建立哈希表记录已经遍历过的链表元素
	// dummyHead := &ListNode{Next: head}
	// //建立哨兵头节点
	// curr := head
	// i := 1
	// for curr != nil {
	// 	listIndexMap[i] = curr
	// 	i++
	// 	curr = curr.Next
	// 	//全部遍历
	// }
	// //声明要跳过的节点前一个节点
	// var prevNode *ListNode
	// if i-n-1 == 0 {
	// 	prevNode = dummyHead
	// } else {
	// 	prevNode = listIndexMap[i-n-1]
	// }
	// prevNode.Next = listIndexMap[i-n].Next
	// return dummyHead.Next
	//法一哈希表
	//法二：这题的关键是要提前知道什么时候到达链表尾，从而将链表的规定倒数第 n 个元素跳过
	//我们可以创建一个先遣节点，让先遣节点先向前走 n 步，再派出主力节点，让主力节点和先遣节点一起向前走，这样当先遣节点到达链表尾时主力节点也就到达了倒数第 n 个节点
	//延迟指针法
	dummyNode := &ListNode{Next: head}
	frontNode := dummyNode
	for i := 1; i <= n; i++ {
		frontNode = frontNode.Next
	}
	mainNode := dummyNode
	for {
		if frontNode.Next == nil {
			mainNode.Next = mainNode.Next.Next
			break
			//跳过目标值
		}
		mainNode = mainNode.Next

		frontNode = frontNode.Next
	}
	return dummyNode.Next
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n1\n
// @lcpr case=end

*/
