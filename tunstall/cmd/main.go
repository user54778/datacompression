package main

import (
	"container/heap"
	"fmt"

	"tunstall.adpollak.net/tunstalltree"
)

func main() {
	a := tunstalltree.Node{Sequence: "a", Probability: 0.6}
	b := tunstalltree.Node{Sequence: "b", Probability: 0.3}
	c := tunstalltree.Node{Sequence: "c", Probability: 0.1}
	h := &tunstalltree.MaxHeap{a, b, c}
	heap.Init(h)
	fmt.Printf("maximium :%#v\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%#v\n", heap.Pop(h))
	}
}
