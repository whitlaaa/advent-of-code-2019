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
	program := getProgram(strings.Split(string(bytes), ","))

	fmt.Println("Part 1: ", runProgram(program, 12, 2))
}

func runProgram(program []int, noun int, verb int) int {
	length := len(program)

	program[1] = noun
	program[2] = verb

	for i := 0; i < length; i = i + 4 {
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
			log.Fatalf("Invalid opcode: %d", opCode)
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
