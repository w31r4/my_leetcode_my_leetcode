/*
 * @lc app=leetcode.cn id=909 lang=golang
 * @lcpr version=30204
 *
 * [909] 蛇梯棋
 */

// @lcpr-template-start
package leetcode

// @lcpr-template-end
// @lc code=start
//快使用无敌的 BFS 想想办法

func snakesAndLadders(board [][]int) int {
	ezSlice := make([]int, len(board)*len(board))
	theEnd := len(board)*len(board) - 1
	isFromHeadToTail := true
	indexNumber := 0
	for i := len(board) - 1; i >= 0; i-- {
		if isFromHeadToTail {
			for j := 0; j < len(board); j++ {
				ezSlice[indexNumber] = board[i][j]
				indexNumber++
			}
			isFromHeadToTail = !isFromHeadToTail
		} else {
			for j := len(board) - 1; j >= 0; j-- {
				ezSlice[indexNumber] = board[i][j]
				indexNumber++
			}
			isFromHeadToTail = !isFromHeadToTail
		}
	}
	bfsQueue := []int{0}
	step := 0
	visited := map[int]bool{}
	//大水漫灌!!!
	//分为远点和变数点，远点求最远，变数点爬一爬
	for len(bfsQueue) != 0 {
		curr := bfsQueue
		bfsQueue = nil
		for _, local := range curr {
			//将当前队列中的所有值取出来遍历范围内有没有特殊值
			//如果有，加入队列中，如果没有，仅加入最大值
			maxNormal := 0
			for i := 1; i <= 6; i++ {

				if local+i == theEnd || ezSlice[local+i]-1 == theEnd {
					//两种可能抵达终点，一种是当前 index 到达终点，另一种是当前数值到达终点，无论哪一种都是达到终点
					return step + 1
				}
				if ezSlice[local+i] != -1 && visited[ezSlice[local+i]-1] == false {
					//用 visited 防止走回头路
					bfsQueue = append(bfsQueue, ezSlice[local+i]-1)
					visited[ezSlice[local+i]-1] = true
					//将一步以内的出现的新的变数点加入 queue 中
				}
				if ezSlice[local+i] == -1 && visited[local+i] == false {
					maxNormal = local + i
					visited[maxNormal] = true
					//记录一步以内的新的远点
					//远点是平常的，只要经过就要被标记上，因为其中没有多余的信息
				}
			}
			if maxNormal != 0 {
				bfsQueue = append(bfsQueue, maxNormal)
			}
			//没有环真对了
		}
		step++
		//一步以内的所有事都做完了
	}
	return -1
}

// @lc code=end

/*
// @lcpr case=start
// [[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,-1,-1,-1,-1,-1],[-1,35,-1,-1,13,-1],[-1,-1,-1,-1,-1,-1],[-1,15,-1,-1,-1,-1]]\n
// @lcpr case=end

// @lcpr case=start
// [[-1,-1],[-1,3]]\n
// @lcpr case=end

*/
