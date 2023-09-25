package main

import "fmt"

type Node struct {
	left  *Node
	right *Node
	data  int32
}

func (n *Node) insert(child *Node) *Node {
	if n != nil {
		if n.left == nil {
			n.left = child
			return n.left
		} else if n.right == nil {
			n.right = child
			return n.right
		}
	}
	return nil
}

func (n *Node) isFull() bool {
	return n.left != nil && n.right != nil
}

func Walk(t *Node, ch chan int32) {
	if t == nil {
		return
	}
	Walk(t.left, ch)
	ch <- t.data
	Walk(t.right, ch)
}

func Walker(t *Node) <-chan int32 {
	ch := make(chan int32)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}
func (n *Node) swap(depth int32, curDepth int32) {
	if n != nil && n.data != -1 {
		if curDepth%depth == 0 {
			n.left, n.right = n.right, n.left
		}
		n.left.swap(depth, curDepth+1)
		n.right.swap(depth, curDepth+1)
	}
}

func (n *Node) print() {
	if n.left != nil && n.right != nil {
		fmt.Printf("Node [%d] - Left [%d] / Right [%d]\n", n.data, n.left.data, n.right.data)
		n.left.print()
		n.right.print()
	}
}

func swapNodes(indexes [][]int32, queries []int32) [][]int32 {
	nodes := make(map[int32]*Node)
	root := &Node{left: nil, right: nil, data: 1}
	nodes[root.data] = root

	allNodes := []int32{}
	for rowIdx, row := range indexes {
		for colIdx, v := range row {
			allNodes = append(allNodes, v)
			nodes[v] = &Node{left: nil, right: nil, data: v}
			if rowIdx == 0 {
				if colIdx == 0 {
					root.left = nodes[v]
				} else {
					root.right = nodes[v]
				}
			}
		}
	}

	curInsertingIndex := 0
	for i := 2; i < len(allNodes); i++ {
		curNode := nodes[allNodes[curInsertingIndex]]
		if curNode.data == -1 {
			curInsertingIndex++
			i -= 1
			continue
		}
		curNode.insert(nodes[allNodes[i]])

		if curNode.isFull() {
			curInsertingIndex++
		}
	}

	result := [][]int32{}
	for _, v := range queries {
		root.swap(v, 1)
		agg := []int32{}
		for i := range Walker(root) {
			if i != -1 {
				agg = append(agg, i)
			}
		}
		result = append(result, agg)
	}

	return result
}

func main() {
	structure := [][]int32{
		{2, 3},
		{4, 5},
		{6, -1},
		{-1, 7},
		{8, 9},
		{10, 11},
		{12, 13},
		{-1, 14},
		{-1, -1},
		{15, -1},
		{16, 17},
		{-1, -1},
		{-1, -1},
		{-1, -1},
		{-1, -1},
		{-1, -1},
		{-1, -1},
	}

	queries := []int32{2, 3}

	swapNodes(structure, queries)
}
