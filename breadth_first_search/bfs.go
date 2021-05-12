package breadth_first_search

import (
	"fmt"

	"../queue"
)

type GraphNode struct {
	Name      string
	Value     int64
	Neighbors map[string]*GraphNode
}

func BFS(start, end *GraphNode) bool {
	// 节点队列
	queue := queue.Init()
	queue.Put(start)

	// 已检测节点
	checked := make(map[string]*GraphNode, 0)
	for queue.Size > 0 {
		ret := queue.Pop()
		if ret == nil {
			return false
		}
		tmpNode := ret.(*GraphNode)
		fmt.Printf("tmpnode:%s\n", tmpNode.Name)
		for _, nei := range tmpNode.Neighbors {
			if nei.Name == end.Name {
				fmt.Printf("found end:%s\n", end.Name)
				return true
			} else {
				_, ok := checked[nei.Name]
				if !ok {
					fmt.Printf("put:%s\n", nei.Name)
					queue.Put(nei)
				}
			}
		}
		checked[tmpNode.Name] = tmpNode
	}

	return false
}
