package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum1 := 0
	sum2 := 0
	for scanner.Scan() {
		mass, _ := strconv.Atoi(scanner.Text())
		sum1 += getModuleFuel(mass, false)
		sum2 += getModuleFuel(mass, true)
	}

	fmt.Println("Part 1: ", sum1)
	fmt.Println("Part 2: ", sum2)
}

func getModuleFuel(mass int, includeFuel bool) int {
	if !includeFuel {
		return calcFuel(mass)
	} else {
		totalFuel := calcFuel(mass)

		fuel := calcFuel(totalFuel)
		for fuel > 0 {
			totalFuel += fuel
			fuel = calcFuel(fuel)
		}

		return totalFuel
	}
}

func calcFuel(mass int) int {
	return (mass / 3) - 2
}
