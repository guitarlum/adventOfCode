package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	direction string
	distance  int
}

func main() {

	http.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		var lines []int = readFileToInt("input/input1.txt")
		fmt.Fprintln(w, "1st Task: Find all increments in given input are: ", calcIncrements(lines))
		fmt.Fprintln(w, "2nd Task: Find all increments in slices of 3: ", calcIncrementsInSlices(lines))
	})

	http.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {
		var instructions = readFileToInstructions("input/input2.txt")
		var x, y = calcXY(instructions)
		fmt.Fprintln(w, "1st Task: Multiply horizontal and depth (x*y): ", x*y)
		x, y = calcXYAim(instructions)
		fmt.Fprintln(w, "2nd Task: Multiply horizontal and depth with aim (x*y): ", x*y)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}

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

func readFileToInt(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		lines = append(lines, i)
	}
	return lines
}

func readFileToInstructions(path string) []Instruction {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []Instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		var instruction Instruction
		instruction.direction = split[0]

		i, err := strconv.Atoi(split[1])
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		instruction.distance = i
		lines = append(lines, instruction)
	}
	return lines
}

func calcIncrements(lines []int) int {
	increments := 0
	for i, v := range lines {
		if i > 0 && v > lines[i-1] {
			increments++
		}
	}
	return increments
}

func calcIncrementsInSlices(lines []int) int {
	increments := 0
	for i, v := range lines {
		if i > 1 && i < (len(lines)-1) {
			if lines[i-2]+lines[i-1]+v < lines[i-1]+v+lines[i+1] {
				increments++
			}
		}
	}
	return increments
}
