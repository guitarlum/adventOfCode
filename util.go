package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func strToInt(s string) int {
	intNumber, _ := strconv.Atoi(strings.TrimSpace(s))
	return intNumber
}

func strToUint8(s string) uint8 {
	intNumber, _ := strconv.Atoi(strings.TrimSpace(s))
	return uint8(intNumber)
}

func setBit(n int, pos uint) int {
	n |= (1 << pos)
	return n
}

func setBitInByte(n byte, pos uint) byte {
	n |= (1 << pos)
	return n
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func RemoveLastElement(array []byte) []byte {
	if len(array) > 0 {
		array = array[:len(array)-1]
	}
	return array
}
