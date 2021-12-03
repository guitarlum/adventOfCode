package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readFileAsNumbers(path string) []int {
	var lines []string = readFileAsStringArray(path)
	var numbers []int
	for _, v := range lines {
		i, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		numbers = append(numbers, i)
	}

	return numbers
}

func readFileAsStringArray(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
