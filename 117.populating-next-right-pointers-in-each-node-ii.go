/*
 * @lc app=leetcode.cn id=117 lang=golang
 * @lcpr version=30204
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
// type Node struct {
// 	Val   int
// 	Left  *Node
// 	Right *Node
// 	Next  *Node
// }

// 思路：递归吗？
// 怎么递归？
// 怎么样跨元素指向呢？
// 我们想像一个卡在一个接近无数层二叉树中间的结点该如何连接起来
// 这题的难点在于跨子树获取下一位要指向的节点
// 或许我们可以构建一个 next 网
// func connect(root *Node) *Node {
// 	if root == nil {
// 		return nil
// 	}
// 	connectToNext(root)
// 	return root
// }
// func connectToNext(node *Node) {
// 	if node == nil || (node.Left == nil && node.Right == nil) {
// 		//处理需要返回的情况
// 		return
// 	}

// 	if node.Left != nil {
// 		node.Left.Next = node.Right
// 	}
// 	//通过父节点已经建立的 Next 网询问父节点的 next 存不存在以及 next.Left 和 Right 存不存在，如果存在那么将当前结点靠右结点指向过去
// 	if node.Left == nil || node.Right == nil {
// 		nextNode := findNextChildNode(node.Next)
// 		if node.Left == nil {
// 			node.Right.Next = nextNode
// 		} else {
// 			node.Left.Next = nextNode
// 		}
// 	} else {
// 		node.Right.Next = findNextChildNode(node.Next)
// 	}
// 	connectToNext(node.Right)
// 	connectToNext(node.Left)
// 	//为什么先右后左？
// 	//因为我们需要从右开始找到 next，如果从左开始，那么会出现一个在右边的结点还没有建立指向右边的 next
// }

// func findNextChildNode(node *Node) *Node {
// 	for node != nil {

// 		if node.Left != nil {
// 			return node.Left

// 		}
// 		if node.Right != nil {
// 			return node.Right
// 		}
// 		node = node.Next
// 	}
// 	return nil
// }

//这种思路不能跨多个子树获取 next，只能跨一个子树
// func connectToNext(root, node1, node2 *Node) *Node {
// 	if node1 == nil && node2 == nil {
// 		return nil
// 	}
// 	if node1 == nil && node2 != nil {
// 		connectToNext(node2, node2.Left, node2.Right)
// 		return node2
// 	}
// 	if node1 != nil && node2 == nil {
// 		connectToNext(node1, node1.Left, node1.Right)
// 		return node1
// 	}

// 	node1.Next = node2

//		leftNode := connectToNext(node1, node1.Left, node1.Right)
//		rightNode := connectToNext(node2, node2.Left, node2.Right)
//		if leftNode != nil {
//			leftNode.Next = rightNode
//		}
//		if root.Next == nil {
//			return node1
//		}
//		return node2
//	}
//
// 法三:BFS 层序遍历
//
//	func connect(root *Node) *Node {
//		if root == nil {
//			return nil
//		}
//		queue := []*Node{root}
//		for queue != nil {
//			tmp := queue
//			queue = nil
//			for i, v := range tmp {
//				if i > 0 {
//					tmp[i-1].Next = v
//				}
//				if v.Left != nil {
//					queue = append(queue, v.Left)
//				}
//				if v.Right != nil {
//					queue = append(queue, v.Right)
//				}
//			}
//		}
//		return root
//	}
//
// 使用链表来记录下一级的状态而非队列，换一个保存数据的数据结构
func connect(root *Node) *Node {
	cur := root
	dummy := &Node{}
	var nxt *Node
	for cur != nil {
		dummy.Next = nil
		nxt = dummy
		for cur != nil {
			if cur.Left != nil {
				nxt.Next = cur.Left
				nxt = cur.Left
			}
			if cur.Right != nil {
				nxt.Next = cur.Right
				nxt = cur.Right
			}
			cur = cur.Next
		}
		cur = dummy.Next
	}
	return root
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,null,7]\n
// @lcpr case=end

// @lcpr case=start
// []\n
// @lcpr case=end

*/
