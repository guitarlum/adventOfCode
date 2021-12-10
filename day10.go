package main

import "sort"

func calcErrorAndCompletionScore(lines []string) (int, int) {
	syntaxPoint := map[byte]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	completionPoint := map[byte]int{
		')': 1,
		']': 2,
		'}': 3,
		'>': 4,
	}
	closerMap := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	errorSum, incompleteLines := extractErrorSumAndIncompleteLines(lines, closerMap, syntaxPoint)
	completionSum := calcCompletionSum(incompleteLines, completionPoint, closerMap)
	sort.Ints(completionSum)

	return errorSum, completionSum[len(completionSum)/2]
}

func calcCompletionSum(incompleteLines []string, completionPoint map[byte]int, closerMap map[byte]byte) []int {
	var completionSum []int
	for _, v := range incompleteLines {
		v = ReverseString(v)
		var sum int = 0
		for i := 0; i < len(v); i++ {
			sum *= 5
			sum += completionPoint[closerMap[v[i]]]
		}
		completionSum = append(completionSum, sum)
	}
	return completionSum
}

func extractErrorSumAndIncompleteLines(lines []string, closerMap map[byte]byte, syntaxPoint map[byte]int) (int, []string) {
	var errorSum int
	var incompleteLines []string

	for _, v := range lines {
		var closings []byte
		var keepLine bool = true
		for i := 0; i < len(v); i++ {
			if v[i] == '(' ||
				v[i] == '[' ||
				v[i] == '{' ||
				v[i] == '<' {
				closings = append(closings, v[i])
			} else if v[i] == closerMap[closings[len(closings)-1]] {
				closings = RemoveLastElement(closings)
			} else {
				errorSum += syntaxPoint[v[i]]
				keepLine = false
				break
			}
		}
		if keepLine {
			incompleteLines = append(incompleteLines, string(closings))
		}
	}
	return errorSum, incompleteLines
}
