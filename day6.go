package main

import (
	"strings"
)

func calcLanternFish(lines []string, days int) uint64 {
	var fishMap map[int]uint64 = make(map[int]uint64)
	extractFish(fishMap, lines)

	for day := 0; day < days; day++ {
		var spawnNew uint64 = fishMap[0]
		fishMap[0] = 0

		for i := 0; i < 8; i++ {
			fishMap[i] = fishMap[i+1]
		}

		fishMap[6] += spawnNew
		fishMap[8] = spawnNew
	}

	return sumOfAllFish(fishMap)
}

func sumOfAllFish(fishMap map[int]uint64) uint64 {
	var sum uint64 = 0
	for _, value := range fishMap {
		sum += value
	}
	return sum
}

func extractFish(fishMap map[int]uint64, lines []string) {
	split := strings.Split(lines[0], ",")
	for _, v := range split {
		fishMap[strToInt(v)]++
	}
}
