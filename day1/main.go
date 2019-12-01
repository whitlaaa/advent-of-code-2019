package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("./input1.txt")
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		sum += calcModuleFuel(value)
	}

	log.Printf("Part 1: %d", sum)
}

func calcModuleFuel(value int) int {
	return (value / 3) - 2
}
