/*
*    @author wangchunyan.wang
*    @date 5/12/21 11:32 PM
 */
package leetcode

import (
	"container/list"
	"fmt"
)

/*
我手中有一堆扑克牌， 但是观众不知道它的顺序。
- 第一步， 我从牌顶拿出一张牌， 放到桌子上。
- 第二步， 我从牌顶再拿一张牌， 放在手上牌的底部。
- 第三步， 重复第一步的操作， 直到我手中所有的牌都放到了桌子上。
最后， 观众可以看到桌子上牌的顺序是：13\12\11\10\9\8\7\6\5\4\3\2\1 请问， 我刚开始拿在手里的牌的顺序是什么？
*/

// 模拟题目动作
// 输入 1,12,2,8,3,11,4,9,5,13,6,10,7
// 输出 1,2,3,4,5,6,7,8,9,10,11,12,13
func getDstOrder(origin []int) {
	stack := list.New()
	queue := list.New()

	// 构造队列
	for _, num := range origin {
		queue.PushBack(num)
	}

	i := 0
	for {
		e := queue.Front()
		if e == nil {
			break
		}
		i++
		queue.Remove(e)
		if i%2 == 1 {
			// 出队 入栈
			stack.PushFront(e.Value)
		} else {
			// 出队 入队
			queue.PushBack(e.Value)
		}
	}

	for {
		e := stack.Front()
		if e == nil {
			break
		}
		fmt.Printf("stack e: %v\n", e.Value)
		stack.Remove(e)
	}
	fmt.Printf("\n\n")
}

// 题目逆操作
// 分析：如果这个操作倒着来
// 1. 从牌底部拿一张牌放到牌顶
// 2. 从桌子上拿一张牌放到牌顶
// 输入: 1,2,3,4,5,6,7,8,9,10,11,12,13
// 输入: 1,12,2,8,3,11,4,9,5,13,6,10,7,
func getOriginalOrder(now []int) []int {
	var original []int
	queue := list.New()

	for i := len(now) - 1; i >= 0; i-- {
		if queue.Front() != nil {
			// 出队尾
			e := queue.Back()
			queue.Remove(e)
			// 入队首
			queue.PushFront(e.Value)
		}
		// 出栈 入队首
		queue.PushFront(now[i])
	}

	// 遍历队列
	for {
		e := queue.Front()
		if e == nil {
			break
		}
		fmt.Printf("e is: %v\n", e.Value)
		queue.Remove(e)
	}

	return original
}
