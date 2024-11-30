package main

import (
	"fmt"

	hufftree "huffman.adpollak.net/tree"
)

func main() {
	probs := []float64{0.4, 0.2, 0.2, 0.1, 0.1}
	bits := []int{1, 2, 3, 4, 4}
	avgLen := hufftree.AverageLength(probs, bits)
	fmt.Println(avgLen, "bits/symbol")
}
