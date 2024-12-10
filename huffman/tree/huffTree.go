package hufftree

import (
	"fmt"
	"log"
	"sort"
)

// Node represents a node in a Huffman Tree. It assumes the probability is known.
type Node struct {
	letter      rune
	probability float64
	left, right *Node
}

// HuffmanTree is simply a wrapper around the root of a HuffmanTree.
type HuffmanTree struct {
	Root *Node
}

func NewHuffmanTree(root *Node) *HuffmanTree {
	return &HuffmanTree{
		Root: root,
	}
}

func NewLeafNode(letter rune, probability float64) *Node {
	return &Node{
		letter:      letter,
		probability: probability,
		left:        nil,
		right:       nil,
	}
}

// BuildHuffmanTree will build a Huffman tree.
func BuildHuffmanTree(nodes []*Node) *HuffmanTree {
	// Step 1: Sort the nodes by probability
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].probability < nodes[j].probability
	})

	// Step 2: Combine nodes until one node (the root) is left
	for len(nodes) > 1 {
		//  Step 2a: Take the two nodes with the lowest probability
		l := nodes[0]
		r := nodes[1]

		//  Step 2b: Create a new internal node that combines these
		// Internal node
		newNode := &Node{
			letter:      0, // only leaves have a rune (letter)
			probability: l.probability + r.probability,
			left:        l,
			right:       r,
		}
		// fmt.Println("Combining nodes into internal node:", newNode)

		// fmt.Printf("%#v\n", nodes)
		//  Step 2c: Remove these nodes and add the new one
		nodes = nodes[2:]

		//  Step 2d: Add this new node to our "sorter"
		nodes = append(nodes, newNode)
		// fmt.Println("Combining nodes")
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].probability < nodes[j].probability
		})

		for _, n := range nodes {
			fmt.Printf("%c %.3f\n", n.letter, n.probability)
		}
	}

	// Step 3: Return the root node
	return &HuffmanTree{
		Root: nodes[0],
	}
}

// PrintHuffmanTree performs an inorder traversal of the Huffman Tree.
func PrintHuffmanTree(node *Node) {
	if node == nil {
		return
	}
	PrintHuffmanTree(node.left)
	fmt.Printf("%#v\n", node)
	PrintHuffmanTree(node.right)
}

// Traverse the Huffman tree to assign binary codewords to each symbol.
func (tree *HuffmanTree) AssignCodes(node *Node, code string, codes map[rune]string) {
	if node == nil {
		return
	}

	// Assign the code to a leaf node.
	if node.left == nil && node.right == nil {
		if node.letter != 0 {
			codes[node.letter] = code
		}
		return
	}

	// Traverse the left and right subtrees
	tree.AssignCodes(node.left, code+"0", codes)
	tree.AssignCodes(node.right, code+"1", codes)
}

// AverageLength computes the average length of a binary code.
func AverageLength(probabilities []float64, bits []int) float64 {
	var length float64
	for i := 0; i < len(bits); i++ {
		length += probabilities[i] * float64(bits[i])
	}
	return length
}

// Redundancy is a measure of efficiency a given code.
func Redundancy(entropy float64, avgLen float64) float64 {
	if avgLen < entropy {
		log.Fatalf("average length can not be less than entropy")
	}
	return avgLen - entropy
}
