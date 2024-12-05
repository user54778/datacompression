package main

import (
	"fmt"

	"golomb.adpollak.net/code"
)

func main() {
	/*
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
	*/
	/*
		t := math.Log2(100)
		fmt.Println(t)
		fmt.Println(math.Ceil(math.Log2(10)))
	*/
	fmt.Println(code.EncodeGolomb(5, 2))
	fmt.Println(code.EncodeGolomb(5, 3))
	fmt.Println(code.EncodeGolomb(5, 7))
	fmt.Println(code.EncodeGolomb(5, 4))
	fmt.Println(code.EncodeGolomb(5, 12))
	fmt.Println(code.EncodeGolomb(5, 8))
	fmt.Println(code.EncodeGolomb(5, 21))
}
