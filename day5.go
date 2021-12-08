package main

import (
	"strings"
)

type Vector struct {
	x1, y1 int
	x2, y2 int
}

func NewVector(x1, y1, x2, y2 int) *Vector {
	v := new(Vector)
	v.x1 = x1
	v.y1 = y1
	v.x2 = x2
	v.y2 = y2
	return v
}

func calcOverlaps(lines []string, withDiagonal bool) int {
	var vectors, maxX, maxY = extractVectors(lines, withDiagonal)
	myMap := buildMap(maxX, maxY)
	drawMap(myMap, vectors)
	var overlaps = countOverlaps(myMap)
	return overlaps
}

func buildMap(maxX int, maxY int) [][]int {
	var myMap = make([][]int, maxX)
	for i := range myMap {
		myMap[i] = make([]int, maxY)
	}
	return myMap
}

func drawMap(myMap [][]int, vectors []Vector) {
	for _, v := range vectors {
		myMap[v.x1-1][v.y1-1]++
		for {
			if v.x1 != v.x2 {
				if v.x1 > v.x2 {
					v.x1--
				} else {
					v.x1++
				}
			}
			if v.y1 != v.y2 {
				if v.y1 > v.y2 {
					v.y1--
				} else {
					v.y1++
				}
			}
			myMap[v.x1-1][v.y1-1]++
			if v.x1 == v.x2 && v.y1 == v.y2 {
				break
			}
		}
	}
}

func countOverlaps(myMap [][]int) int {
	var count int = 0
	for x := range myMap {
		for y := range myMap[x] {
			if myMap[x][y] > 1 {
				count++
			}
		}
	}
	return count
}

func extractVectors(lines []string, withDiagonal bool) ([]Vector, int, int) {
	var vectors []Vector
	var maxX int = 0
	var maxY int = 0
	for _, v := range lines {
		split := strings.Split(v, "->")
		from := strings.Split(split[0], ",")
		to := strings.Split(split[1], ",")
		var vec Vector = *NewVector(strToInt(from[0]),
			strToInt(from[1]),
			strToInt(to[0]),
			strToInt(to[1]))
		if withDiagonal || isHorizontalOrVertical(vec) {
			vectors = append(vectors, vec)
			maxX, maxY = assignMaxValue(vec, maxX, maxY)
		}
	}
	return vectors, maxX, maxY
}

func assignMaxValue(vec Vector, maxX int, maxY int) (int, int) {
	if vec.x1 > maxX {
		maxX = vec.x1
	} else if vec.x2 > maxX {
		maxX = vec.x2
	}
	if vec.y1 > maxY {
		maxY = vec.y1
	} else if vec.y2 > maxY {
		maxY = vec.y2
	}
	return maxX, maxY
}

func isHorizontalOrVertical(vec Vector) bool {
	return vec.x1 == vec.x2 || vec.y1 == vec.y2
}
