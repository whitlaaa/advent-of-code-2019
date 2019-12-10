package intcode

import (
	"log"
	"strconv"
	"strings"
)

const (
	Add = iota + 1
	Multiply
	Store
	Read
	Halt = 99
)

const (
	Position = iota
	Immediate
)

func PrepareProgram(bytes []byte, noun int, verb int) []string {
	program := strings.Split(string(bytes), ",")
	program[1] = strconv.Itoa(noun)
	program[2] = strconv.Itoa(verb)
	return program
}

func RunProgram(input []string) int {
	program := formatInput(input)

	for i := 0; i < len(program); {
		opCode := program[i]
		outputPos := program[i+3]

		var opResult int
		switch opCode {
		case Add:
			opResult = program[program[i+1]] + program[program[i+2]]
			program[outputPos] = opResult
			i += 4
			break
		case Multiply:
			opResult = program[program[i+1]] * program[program[i+2]]
			program[outputPos] = opResult
			i += 4
			break
		case Store:
			log.Println("Store operation")
		case Read:
			log.Println("Read operation")
		case Halt:
			return program[0]
		default:
			log.Panicln("Invalid op code: ", opCode)
		}
	}

	return program[0]
}

func formatInput(input []string) []int {
	var intProgram = []int{}
	for _, v := range input {
		value, _ := strconv.Atoi(v)
		intProgram = append(intProgram, value)
	}
	return intProgram
}
