package trie

import (
	"fmt"
	"strings"
)

// Why a Trie? Well, inputs are coded as a double. These inputs are
// encoded such that entries in the dictionary get longer (generally by one)
// which perfectly matches that of a Trie, which builds substrings.

// Node is the backing Node type for a LZ78Trie. It contains the unique
// index for the explicit dictionary, along with a hashmap
// containing children prefixes to other nodes.
// The entry is implicit in the hashmap.
type Node struct {
	Index    int
	Children map[string]*Node
}

type LZ78Trie struct {
	Root      *Node
	NextIndex int // The next index in the dictionary.
}

// NewLZ78Trie creates a new LZ78Trie. The root node is really just a sentinel
// node, it doesn't store any characters of the tree and is only used to
// create the structure.
func NewLZ78Trie() *LZ78Trie {
	return &LZ78Trie{
		Root: &Node{
			Index:    0,
			Children: make(map[string]*Node),
		},
		NextIndex: 1, // We start indexing at 1
	}
}

// Insert adds a key/value pair in the form of a substring into the LZ78Trie, and
// returns its unique index in the dictionary.
func (t *LZ78Trie) Insert(substring string) int {
	curr := t.Root // start at root node
	for _, r := range substring {
		c := string(r)
		// Check if this letter exists; add it to the trie if not.
		if _, ok := curr.Children[c]; !ok {
			curr.Children[c] = &Node{
				Index:    0,
				Children: make(map[string]*Node),
			}
		}
		curr = curr.Children[c] // move to the next node.
	}

	// Encoded with index value 0; set to the NextIndex.
	if curr.Index == 0 {
		curr.Index = t.NextIndex
		t.NextIndex++ // advance for next substring
	}

	return curr.Index
}

// Find searches for a certain search string key passed in as a parameter in the
// LZ78Trie.
func (t *LZ78Trie) Find(substring string) int {
	curr := t.Root
	for _, r := range substring {
		c := string(r)
		if _, ok := curr.Children[c]; !ok {
			return 0
		}
		curr = curr.Children[c]
	}
	return curr.Index
}

func (t *LZ78Trie) PrintTrie() {
	t.printNode(t.Root, 0)
}

func (t *LZ78Trie) printNode(node *Node, depth int) {
	indent := strings.Repeat("   ", depth)
	if depth > 0 {
		fmt.Printf("%sNode Index: %d, Children: %d\n", indent, node.Index, len(node.Children))
	}

	for s, cN := range node.Children {
		fmt.Printf("%s└─ Char: %s\n", indent, s)
		t.printNode(cN, depth+1)
	}
}
