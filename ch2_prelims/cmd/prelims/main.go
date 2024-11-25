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

	info, err := information.SelfInformation(probability, base)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The self-information associated with the probability %.2f is %.3f\n", probability, info)
}
