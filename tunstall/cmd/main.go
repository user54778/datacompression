package main

import (
	"container/heap"
	"fmt"
	"log"
	"strings"

	"tunstall.adpollak.net/tunstalltree"
)

func main() {
	a := tunstalltree.Node{Sequence: "a", Probability: 0.6}
	b := tunstalltree.Node{Sequence: "b", Probability: 0.3}
	c := tunstalltree.Node{Sequence: "c", Probability: 0.1}
	nodes := []*tunstalltree.Node{&a, &b, &c}
	h := &tunstalltree.MaxHeap{a, b, c}
	heap.Init(h)
	/*
		fmt.Printf("maximium :%#v\n", (*h)[0])
		for h.Len() > 0 {
			fmt.Printf("%#v\n", heap.Pop(h))
		}
	*/

	codebook := make(map[string]string)
	var highest []*tunstalltree.Node
	nodes, highest = tunstalltree.BuildTunstallCode(h, 3, nodes)
	tunstalltree.AssignTunstallCodes(nodes, codebook, 3, highest)

	input := "aaabaabaabaabaaa"
	fmt.Println(input)
	fmt.Println(codebook)
	encoded, err := tunstalltree.TunstallEncode(strings.ToLower(input), codebook)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(encoded)
}
