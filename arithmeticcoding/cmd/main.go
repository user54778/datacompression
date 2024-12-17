package main

import (
	"fmt"
	"log"

	"arithcoding.adpollak.net/unbounded"
	"arithcoding.adpollak.net/utils"
)

func main() {
	/*
		s := []rune{'W', 'X', 'Y', 'Z'}
		t, err := impractical.UnscaledEncodeArithmetic(s)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(t)
		symbols, err := impractical.UnscaledDecodeArithmetic(t)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(symbols)
	*/

	/*
		f, err := os.Open("test.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		content, err := io.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}
	*/
	data, err := utils.CleanInput("num.txt")
	if err != nil {
		log.Fatal(err)
	}
	model := unbounded.NewModel(0)
	model.ComputeCount(data)

	model.ComputeCumulative()
	// fmt.Println(model.Symbols)

	small := "1321"

	out := unbounded.Encode(small, model)
	fmt.Println(out)
	decode := unbounded.Decode(out, model, len(small))
	fmt.Println(decode)
}
