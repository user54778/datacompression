package code

// IMPLEMENT ME: Implement a golomb code.
// This should implement the encoding procedure and decoding procedure for a golomb
// code for some integer m.

func BuildUnary(n int64) string {
	var unaryCode string
	var i int64
	for i = 0; i < n; i++ {
		unaryCode += "1"
	}
	return unaryCode + "0"
}

func DecodeUnary(unaryCode string) int64 {
	var count int64
	for _, r := range unaryCode {
		if string(r) == "1" {
			count++
		} else if string(r) == "0" {
			break
		}
	}
	return count
}

func BuildFastUnary(n int64) int64 {
	// This works because we want to set 2^n bits, minus 1.
	// We have n bits we want to set with a 0 appended.
	// Example: n = 3. 3 = 0011.
	// 1 << 3 = 8, = 1000. We can then get the rest of the bits set:
	// 1 << 3 - 1 = 7 = 1110 == unary form of 3.
	if n > 63 {
		return 0
	} else {
		return (1 << n) - 1
	}
}

func DecodeFastUnary(unary int64) int64 {
	// Now we want to count the amount of bits set.
	// We can do this by applying a bit mask on each LSB per iteration, and move our unary
	// value by 1 per iteration by using an arithmetic right shift by 1
	var n int64 = 0
	for unary > 0 {
		n += unary & 1
		unary >>= 1
	}
	return n
}
