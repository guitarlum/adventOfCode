package main

import (
	"strconv"
)

func calcFlashes(lines []string) (int, int) {
	octos := extractOctos(lines)

	var flashes int = 0
	var steps int = 0
	var allGood bool = false

	for ; !allGood; steps++ {
		for column := range octos {
			for row := range octos[column] {
				increaseOcto(octos, column, row)
			}
		}

		allGood = true
		for column := range octos {
			for row := range octos[column] {
				if octos[column][row] == -1 {
					octos[column][row] = 0
					if steps < 100 {
						flashes++
					}
				} else {
					allGood = false
				}
			}
		}
	}

	return flashes, steps
}

func increaseOcto(octos [][]int, row, column int) {

	if row < 0 || row == len(octos) || column < 0 || column == len(octos[row]) {
		return
	}

	switch octos[column][row] {
	case -1:
		return
	case 9:
		octos[column][row] = -1
		increaseOcto(octos, row-1, column-1)
		increaseOcto(octos, row-1, column)
		increaseOcto(octos, row-1, column+1)
		increaseOcto(octos, row, column-1)
		increaseOcto(octos, row, column+1)
		increaseOcto(octos, row+1, column)
		increaseOcto(octos, row+1, column+1)
		increaseOcto(octos, row+1, column-1)
	default:
		octos[column][row]++
	}
}

func extractOctos(lines []string) [][]int {
	octos := make([][]int, len(lines))

	for i, v := range lines {
		for j := 0; j < len(v); j++ {
			octo, _ := strconv.Atoi(string(v[j]))
			octos[i] = append(octos[i], octo)
		}
	}
	return octos
}
