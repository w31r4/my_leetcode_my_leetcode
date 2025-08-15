/*
 * @lc app=leetcode.cn id=135 lang=golang
 * @lcpr version=30204
 *
 * [135] 分发糖果
 *
 * https://leetcode.cn/problems/candy/description/
 *
 * algorithms
 * Hard (48.45%)
 * Likes:    1643
 * Dislikes: 0
 * Total Accepted:    403.9K
 * Total Submissions: 834K
 * Testcase Example:  '[1,0,2]'
 *
 * n 个孩子站成一排。给你一个整数数组 ratings 表示每个孩子的评分。
 *
 * 你需要按照以下要求，给这些孩子分发糖果：
 *
 *
 * 每个孩子至少分配到 1 个糖果。
 * 相邻两个孩子评分更高的孩子会获得更多的糖果。
 *
 *
 * 请你给每个孩子分发糖果，计算并返回需要准备的 最少糖果数目。
 *
 *
 *
 * 示例 1：
 *
 * 输入：ratings = [1,0,2]
 * 输出：5
 * 解释：你可以分别给第一个、第二个、第三个孩子分发 2、1、2 颗糖果。
 *
 *
 * 示例 2：
 *
 * 输入：ratings = [1,2,2]
 * 输出：4
 * 解释：你可以分别给第一个、第二个、第三个孩子分发 1、2、1 颗糖果。
 * ⁠    第三个孩子只得到 1 颗糖果，这满足题面中的两个条件。
 *
 *
 *
 * 提示：
 *
 *
 * n == ratings.length
 * 1 <= n <= 2 * 10^4
 * 0 <= ratings[i] <= 2 * 10^4
 *
 *
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
func candy(ratings []int) int {
	//局部解得到全局解
	//思路：
	//递增区间糖果递增，不递增直接赋值为 1，单调递减糖果递增，不递减分情况，变大赋值为 2，不变赋值为 1
	//同时如果在部分区间极大值后的递减次数大于极大值之前的递增次数，需要将极大值的获得糖果数量修改为极小值的糖果数量 +1，具体操作是让糖果总量加上修改后 - 修改前的差值
	//前一位大于他，后一位等于他的情况也是极小值，需要重新判断
	//递增递减处理逻辑类似，因为 1234==4321 的糖果数，递减要符合 4321，可以直接赋值为 1234
	//最好不要写多个 if 判断，极其费时费力不讨好
	//人生第一道 hard,On 时间 O1 空间，圆满拿下
	candy := len(ratings)
	//糖果总量
	if candy == 1 {
		return 1
	}
	//排除特殊情况
	nowCandy := 0
	//当前值的糖果数
	peekNum := 5000000
	//一个足够大的值
	for i, rate := range ratings {
		if i == len(ratings)-1 {
			if nowCandy >= peekNum {
				candy += nowCandy - peekNum + 1
			}
			candy += nowCandy
			break
		}
		//当循环到最后一位时，先判断是否需要给极大值糖果增值，然后加上当前糖果量，返回
		if i >= 1 && ratings[i+1] > rate && rate < ratings[i-1] {
			if nowCandy >= peekNum {
				candy += nowCandy - peekNum + 1
			}
			candy += nowCandy
			nowCandy = 1
			continue
		}
		//到达极小值，判断是否需要给极大值糖果增值，加上当前糖果量，给下一个糖果赋值为 1(默认是 0)
		if ratings[i+1] > rate {
			candy += nowCandy
			nowCandy++
			continue
		}
		//递增区间，加上当前糖果，给下一位值糖果数量++
		if i >= 1 && ratings[i+1] < rate && rate > ratings[i-1] {
			peekNum = nowCandy
			candy += nowCandy
			nowCandy = 0
			continue
		}
		//到达极大值，将 peekNum 赋值为极大值的糖果量
		if ratings[i+1] < rate {
			candy += nowCandy
			nowCandy++
			continue
		}
		//递减区间，根据题意和递增区间处理情况相同，为什么呢？是因为我们始终寻求的是最小糖果量，如 123432->123421===123412 ,加上当前糖果，给下一位值糖果数量++
		if i >= 1 && ratings[i-1] > rate && rate == ratings[i+1] {
			if nowCandy >= peekNum {
				candy += nowCandy - peekNum + 1
			}
			peekNum = 5000000
			candy += nowCandy
			nowCandy = 0
			continue
		}
		//遇到前一位大后一位相同的极小值，它同时应该满足极小值的判断是否需要给极大值糖果量增值，同时需要满足对 peekNum 的重新无意义化和对下一位糖果赋值为 0
		if rate == ratings[i+1] {
			peekNum = 5000000
			candy += nowCandy
			nowCandy = 0
			continue
		}
		//不变区间，下一位赋值为 0,peekNum 无意义化

	}
	return candy
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,4,3,2,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,2]\n
// @lcpr case=end

*/
