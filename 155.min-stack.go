/*
 * @lc app=leetcode.cn id=155 lang=golang
 * @lcpr version=30204
 *
 * [155] 最小栈
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
// 栈是从栈顶开始删除元素的，所以在没删除到当前最小值前不需要更新 MinVal，当删除到当前最小值时，再更新 MinVal，所以我们可以建立 minStack 存储出现过的最小值，当出现最小值时将其更新
type MinStack struct {
	Stack       []int
	MinValStack []int
}

func Constructor3() MinStack {
	return MinStack{
		Stack:       []int{},
		MinValStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	if len(this.MinValStack) == 0 || val <= this.MinValStack[len(this.MinValStack)-1] {
		this.MinValStack = append(this.MinValStack, val)
	}
}

func (this *MinStack) Pop() {
	if this.Stack[len(this.Stack)-1] == this.MinValStack[len(this.MinValStack)-1] {
		this.MinValStack = this.MinValStack[:len(this.MinValStack)-1]
	}
	this.Stack = this.Stack[:len(this.Stack)-1]

}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.MinValStack[len(this.MinValStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

/*
// @lcpr case=start
// ["MinStack","push","push","push","getMin","pop","top","getMin"][[],[-2],[0],[-3],[],[],[],[]]\n
// @lcpr case=end

*/
