package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	var matches []int
	for value := 134792; value <= 675810; value++ {
		var hasDupe bool
		var isIncreasing bool = true
		chars := strconv.Itoa(value)
		for k := range chars {
			if k == len(chars)-1 {
				continue
			}

			digit1, _ := strconv.Atoi(string(chars[k]))
			digit2, _ := strconv.Atoi(string(chars[k+1]))

			gte, dupe := compareDigits(digit1, digit2)
			if !gte {
				isIncreasing = false
				break
			}
			if dupe {
				sub := string(chars[k]) + string(chars[k]) + string(chars[k])

				// Part 2
				// If the found duplicate is part of a larger matching group, not a valid duplicate match
				if !strings.Contains(chars, sub) && !hasDupe {
					hasDupe = true
				}
			}
		}
		if hasDupe && isIncreasing {
			matches = append(matches, value)
		}
	}

	log.Println(len(matches))
}

func compareDigits(digit1 int, digit2 int) (isGte bool, isDupe bool) {
	if digit1 == digit2 {
		return true, true
	}

	if digit2 >= digit1 {
		return true, false
	}

	return false, false
}
