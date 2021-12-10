package main

import (
	"sort"
	"strconv"
)

func calcBasinSize(lines []string) int {
	heightMap := buildHeightMap(lines)
	basins := make(map[int]int)

	for i := -1; true; i-- {
		var newBasinMark = markNewBasin(heightMap, i)
		if newBasinMark == 0 {
			break
		}
		basins[newBasinMark]++
		markAllAdjacents(heightMap, newBasinMark, basins)
	}

	basinValues := sortDescBySum(basins)

	var sum = 1
	for i := 0; i < 3; i++ {
		sum *= basinValues[i]
	}
	return sum
}

func sortDescBySum(basins map[int]int) []int {
	var basinValues []int
	for _, v := range basins {
		basinValues = append(basinValues, v)
	}

	sort.Ints(basinValues)
	for i, j := 0, len(basinValues)-1; i < j; i, j = i+1, j-1 {
		basinValues[i], basinValues[j] = basinValues[j], basinValues[i]
	}
	return basinValues
}

func markAllAdjacents(heightMap [][]int, newBasinMark int, basins map[int]int) {
	isUpBasin := func(r int, c int, heightMap [][]int, newBasinMark int) bool {
		return r != 0 && heightMap[r-1][c] == newBasinMark
	}
	isDownBasin := func(r int, c int, heightMap [][]int, newBasinMark int) bool {
		return r != len(heightMap)-1 && heightMap[r+1][c] == newBasinMark
	}
	isLeftBasin := func(r int, c int, heightMap [][]int, newBasinMark int) bool {
		return c != 0 && heightMap[r][c-1] == newBasinMark
	}
	isRightBasin := func(r int, c int, heightMap [][]int, newBasinMark int) bool {
		return c != len(heightMap[r])-1 && heightMap[r][c+1] == newBasinMark
	}

	for true {
		notFound := true
		for r := 0; r < len(heightMap); r++ {
			for c := 0; c < len(heightMap[r]); c++ {
				if heightMap[r][c] != 9 && heightMap[r][c] != newBasinMark {
					if isUpBasin(r, c, heightMap, newBasinMark) ||
						isDownBasin(r, c, heightMap, newBasinMark) ||
						isLeftBasin(r, c, heightMap, newBasinMark) ||
						isRightBasin(r, c, heightMap, newBasinMark) {
						heightMap[r][c] = newBasinMark
						basins[newBasinMark]++
						notFound = false
					}
				}
			}
		}
		if notFound {
			break
		}
	}
}

func markNewBasin(heightMap [][]int, i int) int {
	for r := 0; r < len(heightMap); r++ {
		for c := 0; c < len(heightMap[r]); c++ {
			if heightMap[r][c] >= 0 && heightMap[r][c] != 9 {
				heightMap[r][c] = i
				return heightMap[r][c]
			}
		}
	}
	return 0
}

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
