package main

import (
	"fmt"

	"golomb.adpollak.net/code"
)

func main() {
	var n1 int64 = 1000
	s1 := code.BuildUnary(n1)
	fmt.Println(s1)
	v := code.DecodeUnary(s1)
	fmt.Println(v)

	var n int64 = 63
	s := code.BuildFastUnary(n)
	fmt.Printf("%b\n", s)
	orig := code.DecodeFastUnary(s)
	fmt.Println(orig)
}
