package information

import (
	"fmt"
	"math"
)

// information calculates the self-information from the probability
// of some event A.
func SelfInformation(probability float64, base float64) (float64, error) {
	if base <= 0 {
		base = 2
	}

	if probability <= 0 || probability > 1 {
		return 0, fmt.Errorf("invalid probability: %.2f", probability)
	}

	// NOTE: math.Log returns ln(x)
	return -math.Log(probability) / math.Log(base), nil
}
