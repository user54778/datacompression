package tunstalltree

import (
	"container/heap"
	"fmt"
	"math"
	"strings"
)

// IMPLEMENT ME: Implement an n-bit Tunstall code
// for a source that generates iid letters from an alphabet of size N.
// This should be done with a Heap.

// Step 1: Define the node structure of the tree
type Node struct {
	Sequence    string
	Probability float64
	Left        *Node
	Right       *Node
}

type TunstallTree struct {
	Root *Node
}

func NewLeafNode(sequence string, probability float64) *Node {
	return &Node{
		Sequence:    sequence,
		Probability: probability,
		Left:        nil,
		Right:       nil,
	}
}

// Step 2: We can use a Node slice to implement on the Heap interface.
// Step 3: Implement Len, Less, Swap, Push, and Pop to satisify it (make sure to make it a MAX-heap).
type MaxHeap []Node

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].Probability > h[j].Probability }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
	// Push and pop use pointer receivers since they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(Node))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func BuildTunstallCode(maxHeap *MaxHeap, n int, nodes []*Node) ([]*Node, []*Node) {
	var highests []*Node
	for maxHeap.Len() < (1 << n) {
		top := heap.Pop(maxHeap).(Node)
		highests = append(highests, &top)

		var leafNode *Node
		for _, n := range nodes {
			newProb := top.Probability * n.Probability
			newPrefix := top.Sequence + n.Sequence
			newProb = roundProbability(newProb)

			leafNode = &Node{
				Sequence:    newPrefix,
				Probability: newProb,
				Left:        &top,
				Right:       n,
			}

			nodes = append(nodes, leafNode)
			heap.Push(maxHeap, Node{newPrefix, newProb, nil, nil})
		}
	}

	return nodes, highests
}

func AssignTunstallCodes(nodes []*Node, codebook map[string]string, n int, highest []*Node) {
	code := 0
	var flag bool
	for _, v := range nodes {
		fmt.Println(v)
		// Skip the highest probability node (this is the one that was split)
		if v.Left == nil && v.Right == nil && v.Probability == highest[0].Probability {
			continue
		}
		if len(codebook) == (1 << n) {
			break
		}
		if highest[1].Probability == v.Probability && !flag {
			flag = true
			continue
		}

		binaryCode := fmt.Sprintf("%0*b", n, code)
		codebook[v.Sequence] = binaryCode

		code++
	}
}

func roundProbability(f float64) float64 {
	return math.Round(f*100) / 100
}

func removeNode(nodes []*Node, index int) []*Node {
	toRemove := len(nodes) - index
	fmt.Println("toRemove:", toRemove)
	ret := make([]*Node, 0)
	ret = append(ret, nodes[:index]...)
	return append(ret, nodes[index+1:]...)
}

// Step 4: Implement the Tunstall Coding using this heap.
// Step 5: Generate a codebook from the leaves of this.
/*
func BuildTunstallCode(maxHeap *MaxHeap, n int) map[string]string {
	// IMPLEMENT ME: Iterate from the len of the heap < 2^n
	// Pop the highest probability node, split it into child sequences
	// and push each sequence with the parent sequence concatenated with the child,
	// along with the product of the parent probability and the child.
	// Upon exiting this loop, build the tunstall code book by assigning highest occuring sequence with
	// the lowest bit value of the n bits, i.e., if it was AAA and 2-bit code, AAA -> 00
	parents := buildParents(maxHeap)
	for maxHeap.Len() < (1 << n) {
		top := heap.Pop(maxHeap).(Node)

		fmt.Println(top)
		var index int
		var newParent Node
		for i, v := range parents {
			newProb := top.Probability * v.Probability
			newPrefix := top.Sequence + v.Sequence

			// make a parent node
			if v == top {
				index = i
				fmt.Println(index)
				newParent = Node{Sequence: newPrefix, Probability: newProb}
				fmt.Println("New Parent:", newParent)
			}

			heap.Push(maxHeap, Node{Sequence: newPrefix, Probability: newProb})
		}
		parents = removeParent(parents, index)
		fmt.Println(parents)
		// Insert the parent at the index we removed. I.e., create a new slice with the elements up to parent, and append a node
		// (the parent) at index, and append the rest of the slice following it.
		parents = append(parents[:index], append([]Node{newParent}, parents[index:]...)...)
		fmt.Println(parents)
	}
	return nil
}

func buildParents(maxHeap *MaxHeap) []Node {
	var parents []Node
	for maxHeap.Len() > 0 {
		parents = append(parents, heap.Pop(maxHeap).(Node))
	}
	for _, v := range parents {
		heap.Push(maxHeap, v)
	}
	return parents
}

// removeParent is a helper to remove the parent at slice index i.
func removeParent(parent []Node, index int) []Node {
	return append(parent[:index], parent[index+1:]...)
}
*/

// Step 6: Create an encoder, which will simply match the longest sequence, and
// append the corresponding binary code for the matched sequence to the output.
func TunstallEncode(input string, codebook map[string]string) (string, error) {
	// TODO: implement
	encode := strings.Builder{}

	return encode.String(), nil
}

// Step 7: Create a decoder, which will reverse this process.
func TunstallDecode() string {
	// Simply read the bits from the encoded message, and match them to a corresponding sequence
	// in the codebook, and append that sequence to the output.
	// TODO: Implement
	return ""
}
