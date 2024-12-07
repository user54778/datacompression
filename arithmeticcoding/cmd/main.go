package main

import (
	"fmt"

	"arithcoding.adpollak.net/realbinary"
)

func main() {
	real := 0.8675309
	b := realbinary.RealToBinary(real, 10)
	fmt.Println(b)
}
