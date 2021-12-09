package main

import (
	"strconv"
	"strings"
)

type Pair struct {
	input  []string
	output []string
}

func calcDigit1478(lines []string) int {
	var outputValues = extractOutputValues(lines)
	var digitCount = 0
	for i := range outputValues {
		for _, v := range outputValues[i] {
			if len(v) == 2 || len(v) == 3 || len(v) == 4 || len(v) == 7 {
				digitCount++
			}
		}
	}
	return digitCount
}

/*
 0000
1    2
1    2
 3333
4    5
4    5
 6666

76543210

00111011 = 0 = 59
00100100 = 1 = 36
01011101 = 2 = 93
01101101 = 3 = 109
00101110 = 4 = 46
01101011 = 5 = 107
01111011 = 6 = 123
00100101 = 7 = 37
01111111 = 8 = 127
01101111 = 9 = 111
*/

func calcSumOfOutput(lines []string) int {
	var pairs = extractInputOutputValues(lines)

	var digitMap map[byte]int = make(map[byte]int)
	digitMap[59] = 0
	digitMap[36] = 1
	digitMap[93] = 2
	digitMap[109] = 3
	digitMap[46] = 4
	digitMap[107] = 5
	digitMap[123] = 6
	digitMap[37] = 7
	digitMap[127] = 8
	digitMap[111] = 9
	var sumOfAll int = 0
	for _, p := range pairs {
		var lineCodeMap map[byte]byte = make(map[byte]byte)

		var zeroBit byte = extractZeroBit(p)
		lineCodeMap[0] = zeroBit
		var secondBit, fifthBit = extract2nd5thBit(p)
		lineCodeMap[2] = secondBit
		lineCodeMap[5] = fifthBit
		var firstBit, fourthBit = extract4th1stBit(p, lineCodeMap)
		lineCodeMap[1] = firstBit
		lineCodeMap[4] = fourthBit
		var thirdBit = extractThirdBit(p, lineCodeMap)
		lineCodeMap[3] = thirdBit
		var sixthBit = extractSixthBit(p, lineCodeMap)
		lineCodeMap[6] = sixthBit

		var codeLineMap map[byte]byte = make(map[byte]byte)
		for k, v := range lineCodeMap {
			codeLineMap[v] = k
		}
		var outputNumber string = ""
		for _, v := range p.output {
			var output byte = 0
			for i := 0; i < len(v); i++ {
				output = setBitInByte(output, uint(codeLineMap[v[i]]))
			}
			outputNumber += strconv.Itoa(digitMap[output])
		}
		sumOfAll += strToInt(outputNumber)
	}

	return sumOfAll
}

func extractSixthBit(p Pair, lineCodeMap map[byte]byte) byte {
	eight := findEight(p)
	var known string
	for _, v := range lineCodeMap {
		known += string(v)
	}
	var sixthBit byte
	for i := 0; i < len(eight); i++ {
		if !strings.Contains(known, string(eight[i])) {
			sixthBit = eight[i]
		}
	}
	return sixthBit
}

func findEight(p Pair) string {
	var eight string
	for _, v := range p.input {
		if len(v) == 7 {
			eight = v
			break
		}
	}
	return eight
}

func extractThirdBit(p Pair, lineCodeMap map[byte]byte) byte {
	eight := findEight(p)
	var zero string
	for _, v := range p.input {
		if len(v) == 6 && strings.Contains(v, string(lineCodeMap[2])) && strings.Contains(v, string(lineCodeMap[4])) {
			zero = v
		}
	}

	var thirdBit byte

	for i := 0; i < len(eight); i++ {
		if !strings.Contains(zero, string(eight[i])) {
			thirdBit = eight[i]
		}
	}

	return thirdBit
}

func extract4th1stBit(p Pair, lineCodeMap map[byte]byte) (byte, byte) {
	eight := findEight(p)

	var five string
	for _, v := range p.input {
		if len(v) == 5 && !strings.Contains(v, string(lineCodeMap[2])) {
			five = v
		}
	}

	var forthBit byte

	var fiveSearch string = five + string(lineCodeMap[2])
	for i := 0; i < len(eight); i++ {
		if !strings.Contains(fiveSearch, string(eight[i])) {
			forthBit = eight[i]
		}
	}

	var two string
	for _, v := range p.input {
		if len(v) == 5 && !strings.Contains(v, string(lineCodeMap[5])) {
			two = v
		}
	}
	var firstBit byte

	var secondSearch string = two + string(lineCodeMap[5])
	for i := 0; i < len(eight); i++ {
		if !strings.Contains(secondSearch, string(eight[i])) {
			firstBit = eight[i]
		}
	}

	return firstBit, forthBit
}

func extract2nd5thBit(p Pair) (byte, byte) {
	var one string
	var six string

	for _, v := range p.input {
		if len(v) == 2 {
			one = v
			break
		}
	}

	for _, v := range p.input {
		if len(v) == 6 &&
			((strings.Contains(v, string(one[0])) || strings.Contains(v, string(one[1]))) && !(strings.Contains(v, string(one[0])) && strings.Contains(v, string(one[1])))) {
			six = v
		}
	}

	var secondBit byte
	var fifthBit byte
	if strings.Contains(six, string(one[0])) {
		fifthBit = one[0]
		secondBit = one[1]
	} else {
		fifthBit = one[1]
		secondBit = one[0]
	}

	return secondBit, fifthBit
}

func extractZeroBit(p Pair) byte {
	var seven string
	var one string
	for _, v := range p.input {
		if len(v) == 3 {
			seven = v
		}
		if len(v) == 2 {
			one = v
		}
	}
	var zeroBit byte
	for i := 0; i < len(seven); i++ {
		if !strings.Contains(one, string(seven[i])) {
			zeroBit = seven[i]
		}
	}
	return zeroBit
}

func extractOutputValues(lines []string) [][]string {
	var outValues [][]string = make([][]string, len(lines))
	for i, v := range lines {
		split := strings.Split(v, "|")
		outputValues := strings.Split(split[1], " ")
		for _, value := range outputValues {
			if value != "" {
				outValues[i] = append(outValues[i], value)
			}
		}
	}
	return outValues
}

func extractInputOutputValues(lines []string) []Pair {
	var pairs []Pair
	for _, v := range lines {
		inputOutput := strings.Split(v, "|")
		var pair Pair
		input := strings.Split(inputOutput[0], " ")
		output := strings.Split(inputOutput[1], " ")
		for _, i := range input {
			if i != "" {
				pair.input = append(pair.input, i)
			}
		}
		for _, o := range output {
			if o != "" {
				pair.output = append(pair.output, o)
			}
		}
		pairs = append(pairs, pair)
	}
	return pairs
}
