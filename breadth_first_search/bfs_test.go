package breadth_first_search

import (
	"testing"
)

func TestBFS(t *testing.T) {
	nodeA := &GraphNode{
		Name:      "A",
		Neighbors: make(map[string]*GraphNode),
	}
	nodeB := &GraphNode{
		Name:      "B",
		Neighbors: make(map[string]*GraphNode),
	}
	nodeC := &GraphNode{
		Name:      "C",
		Neighbors: make(map[string]*GraphNode),
	}
	nodeD := &GraphNode{
		Name:      "D",
		Neighbors: make(map[string]*GraphNode),
	}
	nodeF := &GraphNode{
		Name:      "F",
		Neighbors: make(map[string]*GraphNode),
	}
	nodeG := &GraphNode{
		Name:      "G",
		Neighbors: make(map[string]*GraphNode),
	}
	nodeH := &GraphNode{
		Name:      "H",
		Neighbors: make(map[string]*GraphNode),
	}
	nodeM := &GraphNode{
		Name:      "M",
		Neighbors: make(map[string]*GraphNode),
	}
	nodeN := &GraphNode{
		Name:      "N",
		Neighbors: make(map[string]*GraphNode),
	}

	nodeA.Neighbors[nodeB.Name] = nodeB
	nodeA.Neighbors[nodeC.Name] = nodeC
	nodeB.Neighbors[nodeA.Name] = nodeA
	nodeB.Neighbors[nodeF.Name] = nodeF
	nodeC.Neighbors[nodeA.Name] = nodeA
	nodeC.Neighbors[nodeD.Name] = nodeD
	nodeC.Neighbors[nodeF.Name] = nodeF
	nodeD.Neighbors[nodeC.Name] = nodeC
	nodeD.Neighbors[nodeH.Name] = nodeH
	nodeF.Neighbors[nodeB.Name] = nodeB
	nodeF.Neighbors[nodeC.Name] = nodeC
	nodeF.Neighbors[nodeG.Name] = nodeG
	nodeG.Neighbors[nodeF.Name] = nodeF
	nodeG.Neighbors[nodeN.Name] = nodeN
	nodeH.Neighbors[nodeD.Name] = nodeD
	nodeH.Neighbors[nodeM.Name] = nodeM
	nodeM.Neighbors[nodeH.Name] = nodeH
	nodeM.Neighbors[nodeN.Name] = nodeN
	nodeN.Neighbors[nodeG.Name] = nodeG
	nodeN.Neighbors[nodeM.Name] = nodeM

	/*nodes := make([]*GraphNode, 0)
	nodes = append(nodes, nodeA)
	nodes = append(nodes, nodeB)
	nodes = append(nodes, nodeC)
	nodes = append(nodes, nodeD)
	nodes = append(nodes, nodeF)
	nodes = append(nodes, nodeG)
	nodes = append(nodes, nodeH)
	nodes = append(nodes, nodeM)
	nodes = append(nodes, nodeN)*/

	flag := BFS(nodeA, nodeN)
	t.Logf("bfs return:%v\n", flag)
}
