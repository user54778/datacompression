package unbounded

import (
	"fmt"
	"strings"
)

// Decode does not decode an arithmetic encoded byte slice.
func Decode(encoded []byte, model *Model, originalLength int) string {
	var l uint8
	h := word
	var value uint8

	byteIndex := 0
	var bitMask byte = 0x80

	for i := 0; i < 8; i++ {
		value <<= 1
		value |= getBit(encoded, &byteIndex, &bitMask)
	}

	var decoded strings.Builder

	for decoded.Len() < originalLength {
		w := uint16(h-l) + 1
		fmt.Println(w)

		var match string
		for s, p := range model.Symbols {
			high := l + uint8(float64(w)*p.CurrentProbability) - 1
			low := l + uint8(float64(w)*p.CumulativeProbability)
			fmt.Printf("Searching for symbol p: %#v. Range is %d %d.\n", p, low, high)

			if value >= low && value <= high {
				match = s
				break
			}
		}
		decoded.WriteString(match)

		// This is about what will stay the same in a real implementation; the question is really how to get
		// the above stuff implemented correctly?
		p := model.Symbols[match]
		h = l + uint8((float64(w) * p.CurrentProbability)) - 1
		l = l + uint8((float64(w) * p.CumulativeProbability))

		for {
			if l >= 0x80 || h < 0x80 {
				l <<= 1
				h <<= 1
				h |= 1
				value <<= 1
				value |= getBit(encoded, &byteIndex, &bitMask)
			} else if l >= 0x40 && h <= 0xC0 {
				l <<= 1
				l &= 0x7f
				h <<= 1
				h |= 0x81

				value -= 0x40
				value <<= 1
				value |= getBit(encoded, &byteIndex, &bitMask)
			} else {
				break
			}
		}
	}
	return decoded.String()
}

// FIXME: AI-slop. Not really sure how to actually read in this pseudo-infinite stream of bits.
// Also, not sure what this actually does, but either the encoder is wrong or this is wrong.
// I will look to remake this later or do something to fix up this coder, but for now I'm somewhat satisifed
// with the current state of this and have no real desire to mess around with the bit twiddling needed to implement.
func getBit(encoded []byte, byteIndex *int, bitMask *byte) uint8 {
	var bit uint8
	if *byteIndex < len(encoded) {
		if encoded[*byteIndex]&*bitMask != 0 {
			bit = 1
		}

		*bitMask >>= 1
		if *bitMask == 0 {
			*byteIndex++
			*bitMask = 0x80
		}
	}
	return bit
}
