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
	part1 := intcode.PrepareProgram(bytes, 12, 2)
	fmt.Println("Part 1: ", intcode.RunProgram(part1))

	// Part 2
	ogProgram := intcode.PrepareProgram(bytes, 0, 0)
	originalResult := intcode.RunProgram(ogProgram)
	const targetResult int = 19690720

	nounProgram := intcode.PrepareProgram(bytes, 1, 0)
	nounWeight := intcode.RunProgram(nounProgram) - originalResult
	verbProgram := intcode.PrepareProgram(bytes, 0, 1)
	verbWeight := intcode.RunProgram(verbProgram) - originalResult
	noun := (targetResult - originalResult) / nounWeight
	verb := ((targetResult - originalResult) % noun) / verbWeight

	fmt.Printf("Part 2: noun = %d, verb = %d, answer = %d\n", noun, verb, (100*noun)+verb)
}
