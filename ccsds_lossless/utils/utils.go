package utils

import (
	"fmt"
	"log"
	"math"
)

// SelfInformation calculates the self-information from the probability
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

// Entropy calculates the entropy for a given sample space.
func Entropy(sampleSpace []int64, base float64) (float64, error) {
	total := len(sampleSpace)
	if total == 0 {
		return 0.0, fmt.Errorf("empty sample space")
	}

	// Count the occurrence of each value
	frequencies := make(map[int64]int64)
	for _, v := range sampleSpace {
		frequencies[v]++
	}

	// Accumulate information values for each value in the sample space,
	// and then return the avg information, the entropy.
	var entropy float64
	for k, v := range frequencies {
		prob := float64(v) / float64(total)
		fmt.Printf("prob for k %d is: %f\n", k, prob)

		selfInfo, err := SelfInformation(prob, base)
		if err != nil {
			return 0, nil
		}
		entropy += prob * selfInfo
	}
	return entropy, nil
}

// EntropyProbabilities computes entropy if already given a sequence of probabilities.
func EntropyProbabilities(probabilities []float64, base float64) (float64, error) {
	var entropy float64
	for _, prob := range probabilities {
		selfInfo, err := SelfInformation(prob, base)
		if err != nil {
			return 0, err
		}
		entropy += prob * selfInfo
	}
	return entropy, nil
}

// AverageLength computes the average length of a binary code.
func AverageLength(probabilities []float64, bits []int) float64 {
	var length float64
	for i := 0; i < len(bits); i++ {
		length += probabilities[i] * float64(bits[i])
	}
	return length
}

// Redundancy is a measure of efficiency a given code.
func Redundancy(entropy float64, avgLen float64) float64 {
	if avgLen < entropy {
		log.Fatalf("average length can not be less than entropy")
	}
	return avgLen - entropy
}