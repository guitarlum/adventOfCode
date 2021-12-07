package main

import (
	"fmt"
	"log"
	"net/http"
)

type Instruction struct {
	direction string
	distance  int
}

func main() {

	http.HandleFunc("/1", func(w http.ResponseWriter, r *http.Request) {
		var lines []int = readFileAsNumbers("input/input1.txt")
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

	http.HandleFunc("/3", func(w http.ResponseWriter, r *http.Request) {
		var lines = readFileAsStringArray("input/input3.txt")
		var gamma, epislon = calculateGammaEpisolon(lines)
		fmt.Fprintln(w, "1st Task: Multiply gamma and epsilon, power consumption: ", gamma*epislon)
		var oxygen, co02 = calculateOxygenCO02(lines)
		fmt.Fprintln(w, "1st Task: Multiply gamma and epsilon, oxygen * co02: ", oxygen*co02)
	})

	http.HandleFunc("/4", func(w http.ResponseWriter, r *http.Request) {
		var numbers = readFileAsStringArray("input/input4a.txt")
		var bingos = readFileAsStringArray("input/input4b.txt")
		var score = calcBingoWinner(bingos, numbers)
		fmt.Fprintln(w, "1st Task: score of winning bingo: ", score)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
