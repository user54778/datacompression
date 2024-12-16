package main

import (
	"fmt"

	"lz78.adpollak.net/coding"
	"lz78.adpollak.net/trie"
)

func main() {
	trie := trie.NewLZ78Trie()

	sequence := "wabba wabba wabba wabba woo woo woo"
	/*
		for _, r := range sequence {
			fmt.Printf("Searching for: %s. The returned index is: %d\n", string(r), trie.Find(string(r)))
		}
	*/
	lz78 := coding.LZ78{
		Trie: trie,
	}
	output := lz78.LZ78Encode(sequence)
	fmt.Println(output)
	trie.PrintTrie()
}
