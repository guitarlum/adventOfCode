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
		var winnerScore = calcBingoWinner(bingos, numbers)
		fmt.Fprintln(w, "1st Task: score of winning bingo: ", winnerScore)
		var lastWinnerScore = calcBingoLoser(bingos, numbers)
		fmt.Fprintln(w, "2nd Task: score of losing bingo: ", lastWinnerScore)

	})

	http.HandleFunc("/5", func(w http.ResponseWriter, r *http.Request) {
		var lines = readFileAsStringArray("input/input5.txt")
		var without = calcOverlaps(lines, false)
		var with = calcOverlaps(lines, true)

		fmt.Fprintln(w, "1st Task: overlaps without diagonals: ", without)
		fmt.Fprintln(w, "2nd Task: overlaps with diagonals: ", with)

	})

	http.HandleFunc("/6", func(w http.ResponseWriter, r *http.Request) {
		var lines = readFileAsStringArray("input/input6.txt")

		fmt.Fprintln(w, "1st Task: after 80 days: ", calcLanternFish(lines, 80))
		fmt.Fprintln(w, "2nd Task: after 256 days: ", calcLanternFish(lines, 256))
	})

	http.HandleFunc("/7", func(w http.ResponseWriter, r *http.Request) {
		var lines = readFileAsStringArray("input/input7.txt")

		fmt.Fprintln(w, "1st Task: crabs fuel for normal consumption: ", calcCrabFuel(lines, false))
		fmt.Fprintln(w, "2nd Task: crabs fuel for gauss consumption: ", calcCrabFuel(lines, true))
	})

	http.HandleFunc("/8", func(w http.ResponseWriter, r *http.Request) {
		var lines = readFileAsStringArray("input/input8.txt")
		fmt.Fprintln(w, "1st Task: 1,4,7,8 output digit count: ", calcDigit1478(lines))
		fmt.Fprintln(w, "2nd Task: Sum of all outputs: crabs fuel for gauss consumption: ", calcSumOfOutput(lines))
	})

	http.HandleFunc("/9", func(w http.ResponseWriter, r *http.Request) {
		var lines = readFileAsStringArray("input/input9.txt")
		fmt.Fprintln(w, "1st Task: Calc Low point Risk: ", calcLowPointRisk(lines))
		fmt.Fprintln(w, "2nd Task: Multiply 3 largest basins: ", calcBasinSize(lines))
	})

	http.HandleFunc("/10", func(w http.ResponseWriter, r *http.Request) {
		var lines = readFileAsStringArray("input/input10.txt")
		error, complete := calcErrorAndCompletionScore(lines)
		fmt.Fprintln(w, "1st Task: Calc Error Score: ", error)
		fmt.Fprintln(w, "2nd Task: Calc Completion Score: ", complete)
	})

	http.HandleFunc("/11", func(w http.ResponseWriter, r *http.Request) {
		var lines = readFileAsStringArray("input/input11.txt")
		flashes, steps := calcFlashes(lines)
		fmt.Fprintln(w, "1st Task: Calc Flashes: ", flashes)
		fmt.Fprintln(w, "2nd Task: Calc Sync Steps: ", steps)
	})

	http.HandleFunc("/12", func(w http.ResponseWriter, r *http.Request) {
		var lines = readFileAsStringArray("input/input12.txt")
		paths, pathsTwice := calcPaths(lines)
		fmt.Fprintln(w, "1st Task: Paths visit once: ", paths)
		fmt.Fprintln(w, "2nd Task: Paths visit twices: ", pathsTwice)
	})

	http.HandleFunc("/13", func(w http.ResponseWriter, r *http.Request) {

		var points = readFileAsStringArray("input/input13a.txt")
		var instructions = readFileAsStringArray("input/input13b.txt")

		pointsAfterFirst, image := calcPoints(points, instructions)

		fmt.Fprintln(w, "1st Task: Points after first fold: ", pointsAfterFirst)
		fmt.Fprintln(w, "2nd Task: Code: ")
		for y := 0; y <= 6; y++ {
			for x := 0; x <= 40; x++ {
				if image[Point{x, y}] {
					fmt.Fprint(w, "#")
				} else {
					fmt.Fprint(w, ".")
				}
			}
			fmt.Fprintln(w)
		}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
