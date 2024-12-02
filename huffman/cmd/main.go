package main

import (
	"fmt"

	hufftree "huffman.adpollak.net/tree"
)

func main() {
	/*
		probs := []float64{0.4, 0.2, 0.2, 0.1, 0.1}
		bits := []int{1, 2, 3, 4, 4}
		avgLen := hufftree.AverageLength(probs, bits)
		redun := hufftree.Redundancy(2.122, avgLen)
		fmt.Println("Average Length:", avgLen, "bits/symbol")
		fmt.Printf("Redundancy: %.3f bits/symbol\n", redun)
	*/

	// Using example 2.3.1
	symbols := []rune{'a', 'b', 'c', 'd', 'e'}
	probs := []float64{0.2, 0.4, 0.2, 0.1, 0.1}

	var nodes []*hufftree.Node
	for i, symbol := range symbols {
		nodes = append(nodes, hufftree.NewLeafNode(symbol, probs[i]))
	}

	huffManTree := hufftree.BuildHuffmanTree(nodes)

	codes := make(map[rune]string)

	huffManTree.AssignCodes(huffManTree.Root, "", codes)

	hufftree.PrintHuffmanTree(huffManTree.Root)

	for l, c := range codes {
		fmt.Printf("%c: %s\n", l, c)
	}

	// Extract each code length and append it to a slice
	var codeLengths []int
	for _, s := range symbols {
		if c, ok := codes[s]; ok {
			codeLengths = append(codeLengths, len(c))
		}
	}
	avgLen := hufftree.AverageLength(probs, codeLengths)
	redundancy := hufftree.Redundancy(2.122, avgLen)
	fmt.Printf("Average length of Huffman Code: %.2f\n", avgLen)
	fmt.Printf("Redundancy of Huffman Code: %.2f\n", redundancy)
}
