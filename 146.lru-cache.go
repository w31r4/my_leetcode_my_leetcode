/*
 * @lc app=leetcode.cn id=146 lang=golang
 * @lcpr version=30204
 *
 * [146] LRU 缓存
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 哈希表和双向链表用于删除并拼接链表节点
// 哈希表用于 O(1) 复杂度查询
// 所有新加入或者被操作的链表节点都移动到最前面的节点
// 每次超出容量都移出最后一个节点
type DoubleListNode struct {
	Prev *DoubleListNode
	Next *DoubleListNode
	Val  int
	Key  int
}
type LRUCache struct {
	ListMap   map[int]*DoubleListNode
	Capacity  int
	DummyHead *DoubleListNode
	DummyTail *DoubleListNode
	NowCap    int
}

func Constructor4(capacity int) LRUCache {
	dummyHead := &DoubleListNode{}
	dummyTail := &DoubleListNode{}
	dummyHead.Next = dummyTail
	dummyTail.Prev = dummyHead
	ListMap := make(map[int]*DoubleListNode, capacity)
	return LRUCache{Capacity: capacity, NowCap: 0, DummyHead: dummyHead, ListMap: ListMap, DummyTail: dummyTail}
}

func (this *LRUCache) Get(key int) (ans int) {
	ans = -1
	if gotNode, ok := this.ListMap[key]; ok {
		ans = gotNode.Val
		gotNode.Next.Prev = gotNode.Prev
		gotNode.Prev.Next = gotNode.Next
		//跳过当前节点
		oldHead := this.DummyHead.Next
		this.DummyHead.Next = gotNode
		gotNode.Prev = this.DummyHead
		gotNode.Next = oldHead
		oldHead.Prev = gotNode
		//更新头节点

	}
	return ans
}

// put 需要更新 list，更新 map，更新 nowCap，更新 prev 和 next

func (this *LRUCache) Put(key int, value int) {

	if node, ok := this.ListMap[key]; ok {
		node.Val = value // 1. 更新现有节点的值

		// 2. 将这个 'node' 移动到链表头部
		// (从当前位置断开 'node')
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev

		// (在头部重新插入 'node')
		oldActualHead := this.DummyHead.Next
		this.DummyHead.Next = node
		node.Prev = this.DummyHead
		node.Next = oldActualHead
		oldActualHead.Prev = node
		// 注意：'this.ListMap[key]' 已经指向 'node'，此处无需更新 map 的指针。
		// 注意：'this.NowCap' 不变，因为是更新现有项。
	} else {
		newNode := &DoubleListNode{Val: value, Key: key}
		//当插入值不存在时
		//需要更新 nowCap，并判断是否超出容量，如果超出容量，获取链表尾节点的 key，将 map 中的 key 对应键值对删除，并将尾节点更新
		//同时将 nowCap--
		this.NowCap++
		if this.NowCap > this.Capacity {
			deleteTailNode := this.DummyTail.Prev
			deleteKey := deleteTailNode.Key
			deleteTailNode.Next.Prev = deleteTailNode.Prev
			deleteTailNode.Prev.Next = deleteTailNode.Next
			//跳过当前尾节点
			delete(this.ListMap, deleteKey)
			this.NowCap--
		}
		//当容量溢出时去除尾端

		if this.DummyHead.Next == this.DummyTail {
			//当链表中一个元素没有时
			//将 dummyHead 和 dummyTail 都指向当前节点
			this.DummyHead.Next = newNode
			newNode.Prev = this.DummyHead

			this.DummyTail.Prev = newNode
			newNode.Next = this.DummyTail
			this.ListMap[key] = newNode
		} else {
			//如果已有节点
			oldHead := this.DummyHead.Next
			this.DummyHead.Next = newNode
			newNode.Prev = this.DummyHead
			oldHead.Prev = newNode
			newNode.Next = oldHead
			this.ListMap[key] = newNode
		}
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
