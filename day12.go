package main

import (
	"strings"
	"unicode"
)

func calcPaths(lines []string) (int, int) {
	navigation := make(map[string][]string)
	for _, v := range lines {
		both := strings.Split(v, "-")

		if both[1] == "start" {
			navigation[both[1]] = append(navigation[both[1]], both[0])
		} else if both[1] == "end" {
			navigation[both[0]] = append(navigation[both[0]], both[1])
		} else {
			navigation[both[0]] = append(navigation[both[0]], both[1])
			navigation[both[1]] = append(navigation[both[1]], both[0])
		}
	}

	var paths []string
	traverseThroughIndex(navigation, "start", "start", &paths, false)
	var pathsTwice []string
	traverseThroughIndex(navigation, "start", "start", &pathsTwice, true)
	return len(paths), len(pathsTwice)
}

func traverseThroughIndex(navigation map[string][]string, s string, currentPath string, paths *[]string, visitTwice bool) {
	for _, v := range navigation[s] {
		if v == "end" {
			currentPath := currentPath + "," + v
			*paths = append((*paths), currentPath)
			continue
		}
		if isAllLowerCase(v) && strings.Contains(currentPath, v) && hasVisitedTwice(&currentPath, visitTwice) {
			continue
		}
		currentPath := currentPath + "," + v
		traverseThroughIndex(navigation, v, currentPath, paths, visitTwice)
	}
}

func hasVisitedTwice(currentPath *string, visitTwice bool) bool {
	if !visitTwice {
		return true
	}
	isThere := make(map[string]int)
	for _, v := range strings.Split(*currentPath, ",") {
		isThere[v]++
		if isAllLowerCase(v) && isThere[v] > 1 {
			return true
		}
	}
	return false
}

func isAllLowerCase(s string) bool {
	hasUpper := false
	for _, r := range s {
		if unicode.IsUpper(r) {
			hasUpper = true
			break
		}
	}
	return !hasUpper
}
