package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calcXY(instructions []Instruction) (int, int) {
	var x, y int = 0, 0

	for _, c := range instructions {
		switch c.direction {
		case "up":
			y -= c.distance
		case "down":
			y += c.distance
		case "forward":
			x += c.distance
		}
	}

	return x, y
}

func calcXYAim(instructions []Instruction) (int, int) {
	var x, y, aim int = 0, 0, 0

	for _, c := range instructions {
		switch c.direction {
		case "up":
			aim -= c.distance
		case "down":
			aim += c.distance
		case "forward":
			x += c.distance
			y += aim * c.distance
		}
	}

	return x, y
}

func readFileToInstructions(path string) []Instruction {
	var lines []string = readFileAsStringArray(path)
	var instructions []Instruction

	for _, v := range lines {
		split := strings.Split(v, " ")
		var instruction Instruction
		instruction.direction = split[0]

		i, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		instruction.distance = i
		instructions = append(instructions, instruction)
	}

	return instructions
}
