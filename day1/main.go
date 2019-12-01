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

	sum1 := 0
	sum2 := 0
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		sum1 += getModFuel1(mass)
		sum2 += getModFuel2(mass)
	}

	log.Printf("Part 1: %d", sum1)
	log.Printf("Part 2: %d", sum2)
}

func getModFuel1(mass int) int {
	return calcFuel(mass)
}

func getModFuel2(mass int) int {
	totalFuel := calcFuel(mass)

	fuel := calcFuel(totalFuel)
	for fuel > 0 {
		totalFuel += fuel
		fuel = calcFuel(fuel)
	}

	return totalFuel
}

func calcFuel(mass int) int {
	return (mass / 3) - 2
}
