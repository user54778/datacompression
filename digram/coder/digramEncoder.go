package coder

import (
	"strings"
)

// Encode is a digram encoder for a hardcoded dictionary, which is a simple hashmap.
func Encode(dictionary map[string]string, input string) string {
	var output strings.Builder

	for i := 0; i < len(input); i++ {
		var s string
		f := string(input[i])
		if i+1 < len(input) {
			s = string(input[i+1])
		}
		// Step 2: check if in dictionary
		// a) if yes, encode in the value and append to final output
		// b) if no, encode f, then, pull the next character
		// c) unless we're at end of string
		check := f + s
		if v, ok := dictionary[check]; ok {
			output.WriteString(v)
			i++ // Advance our input pointer to account for both runes.
		} else {
			if v, ok := dictionary[f]; ok {
				output.WriteString(v)
			}
		}
	}
	return output.String()
}
