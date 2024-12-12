package main

import (
	"fmt"
	"strings"

	"digram.adpollak.net/coder"
)

func main() {
	dictionary := make(map[string]string)
	dictionary["a"] = "000"
	dictionary["b"] = "001"
	dictionary["c"] = "010"
	dictionary["d"] = "011"
	dictionary["r"] = "100"
	dictionary["ab"] = "101"
	dictionary["ac"] = "110"
	dictionary["ad"] = "111"

	for k, v := range dictionary {
		fmt.Printf("k: %s, v: %s\n", k, v)
	}
	sequence := "abracadabra"
	fmt.Println("original sequence:", sequence)
	encoded := coder.Encode(dictionary, sequence)
	fmt.Println(encoded)
	decoded := coder.Decode(dictionary, encoded)
	fmt.Println(decoded)
	fmt.Println("Difference from Original and Decoded:", strings.Compare(sequence, decoded))
}
