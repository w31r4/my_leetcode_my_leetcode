/*
 * @lc app=leetcode.cn id=380 lang=golang
 * @lcpr version=30204
 *
 * [380] O(1) 时间插入、删除和获取随机元素
 *
 * https://leetcode.cn/problems/insert-delete-getrandom-o1/description/
 *
 * algorithms
 * Medium (52.55%)
 * Likes:    937
 * Dislikes: 0
 * Total Accepted:    220.9K
 * Total Submissions: 420.3K
 * Testcase Example:  '["RandomizedSet","insert","remove","insert","getRandom","remove","insert","getRandom"]\n' +
  '[[],[1],[2],[2],[],[1],[2],[]]'
 *
 * 实现 RandomizedSet 类：
 *
 *
 *
 *
 * RandomizedSet() 初始化 RandomizedSet 对象
 * bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true；否则，返回 false。
 * bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true；否则，返回 false。
 * int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
 *
 *
 * 你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。
 *
 *
 *
 * 示例：
 *
 * 输入
 * ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove",
 * "insert", "getRandom"]
 * [[], [1], [2], [2], [], [1], [2], []]
 * 输出
 * [null, true, false, true, 2, true, false, 2]
 *
 * 解释
 * RandomizedSet randomizedSet = new RandomizedSet();
 * randomizedSet.insert(1); // 向集合中插入 1。返回 true 表示 1 被成功地插入。
 * randomizedSet.remove(2); // 返回 false，表示集合中不存在 2。
 * randomizedSet.insert(2); // 向集合中插入 2。返回 true。集合现在包含 [1,2] 。
 * randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2。
 * randomizedSet.remove(1); // 从集合中移除 1，返回 true。集合现在包含 [2] 。
 * randomizedSet.insert(2); // 2 已在集合中，所以返回 false。
 * randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2。
 *
 *
 *
 *
 * 提示：
 *
 *
 * -2^31 <= val <= 2^31 - 1
 * 最多调用 insert、remove 和 getRandom 函数 2 * 10^5 次
 * 在调用 getRandom 方法时，数据结构中 至少存在一个 元素。
 *
 *
 *
 *
*/

// @lcpr-template-start
package leetcode

import "math/rand"

// @lcpr-template-end
// @lc code=start
// 定义结构体
type RandomizedSet struct {
	numSet map[int]int
	num    []int
}

// 创建构建函数，防止用户直接创建 nil map 和数组初始化 numSet，防止后续操作时出现 nil map panic。
//提供一种标准化的构造方式，类似于面向对象语言的构造函数（如 Java/C++ 的 new RandomizedSet()）。

func Constructor2() RandomizedSet {
	return RandomizedSet{
		numSet: make(map[int]int),
		num:    make([]int, 0),
	}
	//创建哈希表和可变长度数组
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.numSet[val]; ok {
		return false
	} else {
		this.num = append(this.num, val)
		//向变长数组中增加元素
		this.numSet[val] = len(this.num) - 1
		//哈希表的键值对分别存储数组的值和数组的键
		return true
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.numSet[val]; ok {
		index := this.numSet[val]
		//获取数组中要删除的元素的键
		lastVal := this.num[len(this.num)-1]
		//获取数组中最后一个值
		this.num[index] = lastVal
		//将要删除的元素替换成最后一个元素
		this.numSet[lastVal] = index
		//将哈希表中记录的数组的最后一个值的索引修改为替换后的索引
		this.num = this.num[:len(this.num)-1]
		//移除最后一个元素，将数组缩容
		delete(this.numSet, val)
		//移除哈希表中的对应元素
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.num[rand.Intn(len(this.num))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
