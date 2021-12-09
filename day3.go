package main

import (
	"fmt"
	"strconv"
)

func calculateGammaEpisolon(lines []string) (int, int) {
	var common []int = make([]int, len(lines[0]), len(lines[0]))
	for _, v := range lines {
		var maxIndex = len(v)
		for i := 0; i < maxIndex; i++ {
			if v[i] == '1' {
				common[i]++
			} else {
				common[i]--
			}
		}
	}

	var gamma, epsilon int = 0, 0

	for i, j := 0, len(common)-1; i < j; i, j = i+1, j-1 {
		common[i], common[j] = common[j], common[i]
	}

	for i := 0; i < len(common); i++ {
		if common[i] > 0 {
			gamma = setBit(gamma, uint(i))
		} else {
			epsilon = setBit(epsilon, uint(i))
		}
	}

	return gamma, epsilon
}

func calculateOxygenCO02(lines []string) (int, int) {
	appendCriteriaOxygen := func(common int, current byte) bool {
		return (common >= 0 && current == '1') || (common < 0 && current == '0')
	}
	appendCriteriaC02 := func(common int, current byte) bool {
		return (common < 0 && current == '1') || (common >= 0 && current == '0')
	}
	oxygen := reduce(lines, appendCriteriaOxygen)
	co02 := reduce(lines, appendCriteriaC02)

	return oxygen, co02
}

func reduce(lines []string, appendCriteria func(common int, current byte) bool) int {
	maxIndex := len(lines[0])
	reduceList := lines

	for i := 0; i < maxIndex; i++ {
		mostCommon := findMostCommon(reduceList, i)
		var subList []string
		for removeIndex, v := range reduceList {
			if appendCriteria(mostCommon, v[i]) {
				subList = append(subList, reduceList[removeIndex])
			}
		}
		reduceList = subList
		if (i == maxIndex-1) && (len(reduceList) > 1) {
			i = 0
		} else if len(reduceList) == 1 {
			break
		}
	}

	return binaryToInt(reduceList)
}

func findMostCommon(oxyCopy []string, i int) int {
	var mostCommon int = 0

	for _, v := range oxyCopy {
		if v[i] == '0' {
			mostCommon--
		} else {
			mostCommon++
		}
	}
	return mostCommon
}

func binaryToInt(oxyCopy []string) int {
	var oxygen int64 = 0
	oxygen, err := strconv.ParseInt(oxyCopy[0], 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return int(oxygen)
}
