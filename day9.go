package main

import (
	"strconv"
)

func calcLowPointRisk(lines []string) int {
	heightMap := buildHeightMap(lines)
	var sum int = 0
	for row, v := range heightMap {
		for column, cell := range v {
			isUpHigher := func(r int, c int, heightMap [][]int, cell int) bool {
				return row == 0 || heightMap[row-1][column] > cell
			}
			isDownHigher := func(r int, c int, heightMap [][]int, cell int) bool {
				return row == len(heightMap)-1 || heightMap[row+1][column] > cell
			}
			isLeftHigher := func(r int, c int, heightMap [][]int, cell int) bool {
				return column == 0 || heightMap[row][column-1] > cell
			}
			isRightHigher := func(r int, c int, heightMap [][]int, cell int) bool {
				return column == len(heightMap[row])-1 || heightMap[row][column+1] > cell
			}

			if isUpHigher(row, column, heightMap, cell) &&
				isDownHigher(row, column, heightMap, cell) &&
				isLeftHigher(row, column, heightMap, cell) &&
				isRightHigher(row, column, heightMap, cell) {
				sum += cell + 1
			}
		}
	}

	return sum
}

func buildHeightMap(lines []string) [][]int {
	var heightMap [][]int = make([][]int, len(lines))
	for row, v := range lines {
		for i := 0; i < len(v); i++ {
			var number, _ = strconv.Atoi(string(v[i]))
			heightMap[row] = append(heightMap[row], number)
		}
	}
	return heightMap
}
