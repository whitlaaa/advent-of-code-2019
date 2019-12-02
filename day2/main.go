package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Part 1
	fmt.Println("Part 1: ", runProgram(bytes, 12, 2))

	// Part 2
	originalResult := runProgram(bytes, 0, 0)
	const targetResult int = 19690720
	nounWeight := runProgram(bytes, 1, 0) - originalResult
	verbWeight := runProgram(bytes, 0, 1) - originalResult
	noun := (targetResult - originalResult) / nounWeight
	verb := ((targetResult - originalResult) % noun) / verbWeight

	fmt.Printf("Part 2: noun = %d, verb = %d, answer = %d\n", noun, verb, (100*noun)+verb)
}

func runProgram(bytes []byte, noun int, verb int) int {
	program := getProgram(strings.Split(string(bytes), ","))
	length := len(program)

	program[1] = noun
	program[2] = verb

	for i := 0; i < length; i += 4 {
		opCode := program[i]
		input1 := program[program[i+1]]
		input2 := program[program[i+2]]
		outputPos := program[i+3]

		var opResult int
		if opCode == 1 {
			opResult = input1 + input2
		} else if opCode == 2 {
			opResult = input1 * input2
		} else if opCode == 99 {
			break
		} else {
			log.Fatalln("Invalid opcode: ", opCode)
		}

		program[outputPos] = opResult
	}

	return program[0]
}

func getProgram(input []string) []int {
	var intProgram = []int{}
	for _, v := range input {
		value, _ := strconv.Atoi(v)
		intProgram = append(intProgram, value)
	}
	return intProgram
}
