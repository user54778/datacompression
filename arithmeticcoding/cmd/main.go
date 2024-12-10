package main

import (
	"fmt"
	"log"

	"arithcoding.adpollak.net/impractical"
)

func main() {
	s := []rune{'W', 'X', 'Y', 'Z'}
	t, err := impractical.UnscaledEncodeArithmetic(s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
	symbols, err := impractical.UnscaledDecodeArithmetic(t)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(symbols)
}
