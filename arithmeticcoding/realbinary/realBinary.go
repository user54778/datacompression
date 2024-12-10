package realbinary

// RealToBinary converts any real number into a binary number
// that falls under the interval [0, 1).
// This performs binary search on the binary digits of real by determining
// which side of the current midpoint the real number lies.
func RealToBinary(real float64, precision int) string {
	var binary string
	l := 0.0
	r := 1.0
	for i := 0; i < precision; i++ {
		mid := (l + r) / 2
		if real < mid {
			binary += "0"
			r = mid
		} else {
			binary += "1"
			l = mid
		}
	}
	return binary
}
