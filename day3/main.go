package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var origin point = point{0, 0}

func main() {
	wirePoints1, wirePoints2 := getWirePoints()

	// var intersections []point
	// for _, p1 := range wirePoints1 {
	// 	for _, p2 := range wirePoints2 {

	// 	}
	// }

	log.Println(wirePoints1, wirePoints2)
}

func getWirePoints() (wp1 []point, wp2 []point) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		steps1 := strings.Split(scanner.Text(), ",")

		// wp1 = append(wp1, origin)
		for i, v := range steps1 {
			direction := v[:1]
			magnitude, _ := strconv.Atoi(v[1:])

			if i == 0 {
				wp1 = append(wp1, getSegmentEndPoint(origin, direction, magnitude))
			} else {
				wp1 = append(wp1, getSegmentEndPoint(wp1[i], direction, magnitude))
			}
		}
	}

	if scanner.Scan() {
		steps2 := strings.Split(scanner.Text(), ",")

		// wp2 = append(wp2, origin)
		for i, v := range steps2 {
			direction := v[:1]
			magnitude, _ := strconv.Atoi(v[1:])

			if i == 0 {
				wp2 = append(wp2, getSegmentEndPoint(origin, direction, magnitude))
			} else {
				wp2 = append(wp2, getSegmentEndPoint(wp2[i], direction, magnitude))
			}
		}
	}

	return
}

func getSegmentEndPoint(start point, direction string, magnitude int) point {
	switch direction {
	case "U":
		return point{start.x, start.y + magnitude}
	case "D":
		return point{start.x, start.y - magnitude}
	case "R":
		return point{start.x + magnitude, start.y}
	case "L":
		return point{start.x - magnitude, start.y}
	default:
		return point{0, 0}
	}
}

type point struct {
	x int
	y int
}
