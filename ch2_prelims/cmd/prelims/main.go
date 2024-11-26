package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	"losslessprelims.adpollak.net/information"
)

func main() {
	var probability float64
	var baseInput string
	flag.Float64Var(&probability, "probability", 0.0, "supply the probability of event a")
	flag.StringVar(&baseInput, "base", "2", "supply the base of logarithm")

	flag.Parse()

	// Parse the given base
	var base float64
	if strings.ToLower(baseInput) == "e" {
		base = math.E
	} else {
		var err error
		base, err = strconv.ParseFloat(baseInput, 64)
		if err != nil {
			log.Fatalf("invalid base: %v", err)
		}
	}

	/*
		info, err := information.SelfInformation(probability, base)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The self-information associated with the probability %.2f is %.3f\n", probability, info)
	*/
	/*
		probs := []float64{1.0 / 16.0, 2.0 / 16.0}
		entropy, err := information.Entropy(probs, base)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Entropy: %.3f\n", entropy)
	*/
	/*
		sampleSpace := []int64{1, 2, 3, 2, 3, 4, 5, 4, 5, 6, 7, 8, 9, 8, 9, 10}
		entropy, err := information.Entropy(sampleSpace, base)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%.3f\n", entropy)

		residualSpace := []int64{1, 1, 1, -1, 1, 1, 1, -1, 1, 1, 1, 1, 1, -1, 1, 1}
		entropyCorrelation, err := information.Entropy(residualSpace, base)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%.2f\n", entropyCorrelation)
	*/

	threeA := []int64{1, 2, 3, 4}
	threeAEntropy, err := information.Entropy(threeA, base)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f\n", threeAEntropy)
	threeB := []float64{1.0 / 2, 1.0 / 4, 1.0 / 8, 1.0 / 8}
	threeBEntropy, err := information.EntropyProbabilities(threeB, base)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f\n", threeBEntropy)

	threeC := []float64{0.505, 1.0 / 4, 1.0 / 8, 0.12}
	threeCEntropy, err := information.EntropyProbabilities(threeC, base)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%.2f\n", threeCEntropy)
}
