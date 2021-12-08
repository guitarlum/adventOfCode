package main

import (
	"math"
	"sort"
	"strings"
)

func calcCrabFuel(lines []string, gausFuel bool) int {
	var crabs = extractCrabs(lines)
	var fuel int = 0
	if gausFuel {
		var average = calcAverage(crabs)

		for _, v := range crabs {
			var steps = int(math.Abs(float64(average - v)))
			fuel += (steps * (steps + 1) / 2)
		}
	} else {
		var median = calcMedian(crabs)
		for _, v := range crabs {
			fuel += int(math.Abs(float64(median - v)))
		}
	}

	return fuel
}

func calcMedian(crabs []int) int {
	sort.Ints(crabs)
	return crabs[len(crabs)/2]
}

func calcAverage(crabs []int) int {
	var sum int = 0
	for _, v := range crabs {
		sum += v
	}
	return sum / len(crabs)
}

func extractCrabs(lines []string) []int {
	crabStrings := strings.Split(lines[0], ",")
	var crabs []int
	for i := range crabStrings {
		crabs = append(crabs, strToInt(crabStrings[i]))
	}
	return crabs
}
