package unbounded

import "fmt"

const word uint8 = 0xff

func Encode(data string, model *Model) {
	// 1) init l, and h to word limit
	// 2) read in every r
	//  2a) get the prob of r
	//  2b) compute l and h
	//    3) infinitely loop until none of E1, E2 E3 holds.
	//      3a) E1 -> h < 1/2 range
	//        4) output bit 0 plus pending E3 bits
	//      3b) E2 -> l >= 1/2 range
	//        5) output bit 1 plus pending E3 bits
	//      3c) E3 -> l >= 1/4 && h < 3/4
	//        6) append pending bit
	//        7) shift l << 1; shift 0 into LSB, complement MSB of l
	//        8) shift h << 1; shift 1 into LSB, complement MSB of h.
	//      shift l << 1; set 0 into LSB
	//      shift h << 1; set 1 into LSB
	//    9) otherwise, break out of infinite loop
	// 10) output the tag
	// 13) append our final tag to the output and return that
	var l uint8 = 0
	h := word
	var pendingBits uint8
	output := make([]byte, 0)

	for _, r := range data {
		w := h - l + 1
		p := model.Symbols[string(r)]

		h = l + uint8((float64(w) * p.CurrentProbability))
		l = l + uint8((float64(w) * p.CumulativeProbability))

		for {
			if h < (word >> 1) {
				outputPendingBits(&output, &pendingBits, 0)
				l <<= 1
				h <<= 1
				h |= 1
				fmt.Println("E1")
				fmt.Printf("Updated l: %b, Updated h: %b\n", l, h)
				fmt.Printf("Updated output: %b\n", output)
			} else if l >= (word >> 1) {
				outputPendingBits(&output, &pendingBits, 1)
				l <<= 1
				h <<= 1
				h |= 1
				fmt.Println("E2")
				fmt.Printf("Updated l: %#x, Updated h: %#x\n", l, h)
				fmt.Printf("Updated output: %b\n", output)
			} else if l >= (word>>2) && h < ((word>>2)*3) {
				pendingBits++
				l <<= 1
				// l &= 0x7fffffff
				l &= 0x7f
				h <<= 1
				// h |= 0x80000001
				h |= 0x80
				fmt.Println("E3")
				fmt.Printf("Updated l: %#x, Updated h: %#x\n", l, h)
				fmt.Printf("Updated output: %b\n", output)
			} else {
				break
			}
		}
	}
	if l < (word >> 2) {
		outputPendingBits(&output, &pendingBits, 0)
	} else {
		outputPendingBits(&output, &pendingBits, 1)
	}

	fmt.Println("Final output", output)
}

func outputPendingBits(output *[]byte, pendingBits *uint8, inBit byte) {
	*output = append(*output, inBit)
	fmt.Println("Pending bits: ", *pendingBits)
	for i := *pendingBits; i > 0; i-- {
		*output = append(*output, inBit^1)
	}
	*pendingBits = 0
}
