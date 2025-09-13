package huffman

import (
	"container/heap"
)

type huffmanNode struct {
	weight  int
	element rune
	isLeaf  bool
	left    *huffmanNode
	right   *huffmanNode
}

type queue []huffmanNode

func (h queue) Len() int           { return len(h) }
func (h queue) Less(i, j int) bool { return h[i].weight < h[j].weight }
func (h queue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *queue) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(huffmanNode))
}

func (h *queue) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var (
	internalPrefixTable map[rune]string
)

func buildTree(frequency map[rune]int) *huffmanNode {

	huffmanQueue := &queue{}
	heap.Init(huffmanQueue)

	for char, freq := range frequency {
		heap.Push(huffmanQueue, huffmanNode{freq, char, true, nil, nil})
	}

	for huffmanQueue.Len() > 1 {
		firstNode := heap.Pop(huffmanQueue).(huffmanNode)
		secondNode := heap.Pop(huffmanQueue).(huffmanNode)

		heap.Push(huffmanQueue, huffmanNode{firstNode.weight + secondNode.weight, 0, false, &firstNode, &secondNode})
	}

	if huffmanQueue.Len() == 1 {
		headNode := heap.Pop(huffmanQueue).(huffmanNode)
		return &headNode
	}

	return nil
}

func constructTable(headNode *huffmanNode) map[rune]string {
	internalPrefixTable = make(map[rune]string)
	traverseTree(headNode, "")
	return internalPrefixTable
}

func traverseTree(node *huffmanNode, prefix string) {
	if node == nil {
		return
	}

	if node.isLeaf {
		internalPrefixTable[node.element] = prefix
		return
	}
	traverseTree(node.left, prefix+"0")
	traverseTree(node.right, prefix+"1")
}
