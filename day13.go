package main

import (
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

type Fold struct {
	xy string
	v  int
}

func calcPoints(points []string, instructions []string) (int, map[Point]bool) {
	folds := extractFolds(instructions)
	image := fillImage(points)

	sum := 0
	firstInstruction := true

	for _, f := range folds {

		switch f.xy {
		case "x":
			for p := range image {
				if image[p] && p.X > f.v {
					image[p] = false
					newX := f.v - (p.X - f.v)
					if newX >= 0 {
						image[Point{newX, p.Y}] = true
					}
				}
			}
		case "y":
			for p := range image {
				if image[p] && p.Y > f.v {
					image[p] = false
					newY := f.v - (p.Y - f.v)
					if newY >= 0 {
						image[Point{p.X, newY}] = true
					}
				}
			}
		}
		if firstInstruction {
			sum = 0
			for p := range image {
				if image[p] {
					sum++
				}
			}
			firstInstruction = false
		}
	}

	return sum, image
}

func fillImage(points []string) map[Point]bool {
	image := make(map[Point]bool)
	for _, v := range points {
		split := strings.Split(v, ",")
		x, _ := strconv.Atoi(split[0])
		y, _ := strconv.Atoi(split[1])
		image[Point{x, y}] = true
	}
	return image
}

func extractFolds(instructions []string) []Fold {
	folds := make([]Fold, len(instructions))
	for i, v := range instructions {
		split := strings.Split(v, "=")
		number, _ := strconv.Atoi(string(split[1]))
		folds[i] = Fold{split[0], number}
	}
	return folds
}
