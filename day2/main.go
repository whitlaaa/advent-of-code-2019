package main

import (
	"advent-of-code-2019/intcode"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	fmt.Println("Part 1: ", intcode.RunProgram(bytes, 12, 2))

	// Part 2
	originalResult := intcode.RunProgram(bytes, 0, 0)
	const targetResult int = 19690720
	nounWeight := intcode.RunProgram(bytes, 1, 0) - originalResult
	verbWeight := intcode.RunProgram(bytes, 0, 1) - originalResult
	noun := (targetResult - originalResult) / nounWeight
	verb := ((targetResult - originalResult) % noun) / verbWeight

	fmt.Printf("Part 2: noun = %d, verb = %d, answer = %d\n", noun, verb, (100*noun)+verb)
}
