package main

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
