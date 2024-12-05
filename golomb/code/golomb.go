package code

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// IMPLEMENT ME: Implement a golomb code.
// This should implement the encoding procedure and decoding procedure for a golomb
// code for some integer m.
func EncodeGolomb(m int, n int) (string, error) {
	if m <= 0 || n <= 0 {
		return "", fmt.Errorf("m and n must be > 0")
	}
	// For the first 2^ceil(log_2m)-m values, we will use the floor(log_2m)-bit binary representation,
	// and use r + 2^(ceil(log_2m)) - m for the rest.
	ceilLog := math.Ceil(math.Log2(float64(m)))
	floorLog := math.Floor(math.Log2(float64(m)))
	fmt.Printf("ceilLog: %.2f, floorLog: %.2f\n", ceilLog, floorLog)

	representation := 8
	if representation < int(ceilLog) {
		representation *= 2
	}

	// Context -> upper represents the ceil of log_2m, which is used in both equation
	// for determining which log_2-bit binary representation to use.
	// upper := math.Exp2(ceilLog)
	firstValues := representation - m
	fmt.Println(ceilLog, math.Exp2(ceilLog))
	fmt.Println(representation, firstValues)

	// q will be represented by the unary code of q
	q := n / m // Go integer division tends to 0, but this is fine for us since n and m are both required > 0.
	r := n - q*m
	fmt.Printf("quotient: %d, remainder: %d\n", q, r)

	unaryQ := BuildUnary(q)
	fmt.Println("unary value of q:", unaryQ)

	var encodeRemainder string
	if r < firstValues {
		encodeRemainder = strconv.FormatInt(int64(r), 2)
		encodeRemainder = encodeRemainder + strings.Repeat("0", int(floorLog)-len(encodeRemainder))
	} else {
		encodeRemainder = strconv.FormatInt(int64(r), 2)
		encodeRemainder = encodeRemainder + strings.Repeat("0", int(ceilLog)-len(encodeRemainder))
	}
	return unaryQ + encodeRemainder, nil
}

func DecodeGolomb() int {
	// TODO: implement me
	return 0
}

func IsPowerTwo(x int) bool {
	// Essentially, take x, subtract by 1.
	// if x in bin is even, itll be a zero (and therefore divisible by two)
	// This won't work for 0, but doesn't matter since we don't call this on 0.
	return (x & (x - 1)) == 0
}

func BuildUnary(n int) string {
	var unaryCode string
	for i := 0; i < n; i++ {
		unaryCode += "1"
	}
	return unaryCode + "0"
}

func DecodeUnary(unaryCode string) int {
	var count int
	for _, r := range unaryCode {
		if string(r) == "1" {
			count++
		} else if string(r) == "0" {
			break
		}
	}
	return count
}

func BuildFastUnary(n int) int {
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

func DecodeFastUnary(unary int) int {
	// Now we want to count the amount of bits set.
	// We can do this by applying a bit mask on each LSB per iteration, and move our unary
	// value by 1 per iteration by using an arithmetic right shift by 1
	n := 0
	for unary > 0 {
		n += unary & 1
		unary >>= 1
	}
	return n
}
