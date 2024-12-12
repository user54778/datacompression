package coder

import (
	"fmt"
	"strings"
)

// Decode is a digram decoder using a hardcoded dictionary; it takes in input
// and decodes an encoded digram sequence
func Decode(dictionary map[string]string, input string) string {
	var output strings.Builder

	// NOTE: Calling string() around the concatentation results in
	// Go thinking I am asking for a numeric sum of byte values
	/*
		for i := len(input) - 1; i >= 0; i-- {
			var bit string
			if i-2 > 0 {
				bit = string(input[i]) + string(input[i-1]) + string(input[i-2])
				output.WriteString(bit)
				i -= 2
			} else {
				fmt.Println("rem i:", input[i])
			}
		}
	*/
	for i := 0; i < len(input); i++ {
		var bit string
		if i+2 < len(input) {
			bit = string(input[i]) + string(input[i+1]) + string(input[i+2])
			fmt.Println("bit: ", bit)
			if k, ok := valueExistInMap(dictionary, bit); ok {
				output.WriteString(k)
			}
			i += 2
		}
	}
	return output.String()
}

func valueExistInMap(dictionary map[string]string, target string) (string, bool) {
	for k, v := range dictionary {
		if v == target {
			return k, true
		}
	}
	return "", false
}
