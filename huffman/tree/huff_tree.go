package hufftree

// Node represents a node in a Huffman Tree. It assumes the probability is known.
type Node struct {
	letter      rune
	probability float64
	left, right *Node
}

// HuffmanTree is simply a wrapper around the root of a HuffmanTree.
type HuffmanTree struct {
	root *Node
}

func NewHuffmanTree(root *Node) *HuffmanTree {
	return &HuffmanTree{
		root: root,
	}
}

func NewNode(letter rune, probability float64) *Node {
	return &Node{
		letter:      letter,
		probability: probability,
		left:        nil,
		right:       nil,
	}
}

// TODO: Implement a method for building a Huffman Tree.
// TODO: Implement sorting.
// TODO: Implement frequency count of letters from an alphabet.
// TODO: Implement average length of code.
func AverageLength(probabilities []float64, bits []int) float64 {
	var length float64
	for i := 0; i < len(bits); i++ {
		length += probabilities[i] * float64(bits[i])
	}
	return length
}

func Redundancy(entropy float64, avgLen float64) float64 {
	return entropy / avgLen
}
