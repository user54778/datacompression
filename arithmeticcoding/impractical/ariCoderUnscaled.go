package impractical

import (
	"fmt"
	"log"
	"math"
	"strings"
)

type ProbabilityPair struct {
	Probability float64
	Cumulative  float64
}

func modelProbability(c rune) (ProbabilityPair, error) {
	if c >= 'A' && c <= 'Z' {
		return ProbabilityPair{Probability: float64(c-'A') * .01, Cumulative: float64(c-'A')*0.01 + 0.01}, nil
	}
	return ProbabilityPair{}, fmt.Errorf("character is out of range")
}

func getSymbol(s float64) (rune, error) {
	if s >= 0.0 && s < 0.26 {
		return 'A' + rune(s*100), nil
	}
	return 0, fmt.Errorf("symbol not in range")
}

func roundProbability(f float64) float64 {
	return math.Round(f*100) / 100
}

// unscaledEncodeArithmetic builds an unscaled Arithmetic code. Note that this
// will be subject to underflow and is only built for learning purposes.
func UnscaledEncodeArithmetic(symbols []rune) (float64, error) {
	low := 0.0
	high := 1.0
	for _, s := range symbols {
		probPair, err := modelProbability(s)
		if err != nil {
			log.Fatal(err)
		}

		w := high - low
		high = low + w*probPair.Cumulative
		low = low + w*probPair.Probability
	}
	// We now have a final message [low, high), and our tag is the center of this range.
	fmt.Printf("Final interval: [%f, %f)\n", low, high)
	tag := low + (high-low)/2
	return tag, nil
}

// UnscaledDecodeArithmetic decodes an arithmetic code, assuming its
// tag is given as a float value.
func UnscaledDecodeArithmetic(message float64) (string, error) {
	low := 0.0
	high := 1.0
	var symbols strings.Builder
	for {
		w := high - low
		c, err := getSymbol((message - low) / w)
		if err != nil {
			return "", err
		}

		symbols.WriteString(string(c))

		if c == 'Z' {
			return symbols.String(), nil
		}

		probPair, err := modelProbability(c)
		if err != nil {
			return "", err
		}

		high = low + w*probPair.Cumulative
		low = low + w*probPair.Probability
	}
}
