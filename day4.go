package main

import (
	"strconv"
	"strings"
)

type BingoBoard struct {
	winningNumber int
	board         [5][5]int
}

func calcBingoWinner(bingos []string, numbersInput []string) int {
	var numbers = convertNumbersInput(numbersInput)
	var bingoBoards = convertBingosToBoards(bingos)

	var bingoBoard = determineWinningBoard(bingoBoards, numbers)

	return calcScore(bingoBoard)
}

func calcBingoLoser(bingos []string, numbersInput []string) int {
	var numbers = convertNumbersInput(numbersInput)
	var bingoBoards = convertBingosToBoards(bingos)

	var bingoBoard = determineLoserBoard(bingoBoards, numbers)

	return calcScore(bingoBoard)
}

func determineLoserBoard(bbs []BingoBoard, numbers []int) BingoBoard {
	var amount int = len(bbs)
	for _, v := range numbers {
		for i := range bbs {
			if bbs[i].winningNumber == -1 {
				for row := 0; row < 5; row++ {
					for column := 0; column < 5; column++ {
						if bbs[i].board[row][column] == v {
							bbs[i].board[row][column] = -1
							if isWinnerBoard(bbs[i]) {
								bbs[i].winningNumber = v
								amount--
								if amount == 0 {
									return bbs[i]
								}
							}
						}
					}
				}
			}
		}
	}
	var noBoard BingoBoard
	return noBoard
}

func determineWinningBoard(bbs []BingoBoard, numbers []int) BingoBoard {
	for _, v := range numbers {
		for i := range bbs {
			for row := 0; row < 5; row++ {
				for column := 0; column < 5; column++ {
					if bbs[i].board[row][column] == v {
						bbs[i].board[row][column] = -1
						if isWinnerBoard(bbs[i]) {
							bbs[i].winningNumber = v
						}
					}
				}
			}
		}
	}

	var noBoard BingoBoard
	return noBoard
}

func isWinnerBoard(bb BingoBoard) bool {
	var winCount int = 0

	for row := 0; row < 5; row++ {
		for column := 0; column < 5; column++ {
			if bb.board[row][column] == -1 {
				winCount++
				if winCount == 5 {
					return true
				}
			} else {
				winCount = 0
				break
			}
		}
	}

	for column := 0; column < 5; column++ {
		for row := 0; row < 5; row++ {
			if bb.board[row][column] == -1 {
				winCount++
				if winCount == 5 {
					return true
				}
			} else {
				winCount = 0
				break
			}
		}
	}

	return false
}

func calcScore(bb BingoBoard) int {
	var sum = addNumbers(findUnmarkedNumbers(bb))

	return sum * bb.winningNumber
}

func addNumbers(i []int) int {
	var sum int = 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func findUnmarkedNumbers(board BingoBoard) []int {
	var unmarkedNumbers []int
	for _, row := range board.board {
		for _, value := range row {
			if value != -1 {
				unmarkedNumbers = append(unmarkedNumbers, value)
			}
		}
	}
	return unmarkedNumbers
}

func convertBingosToBoards(bingos []string) []BingoBoard {
	var bingoBoards []BingoBoard
	var currentBoard *BingoBoard = new(BingoBoard)
	currentBoard.winningNumber = -1
	for row, v := range bingos {
		if len(v) == 0 {
			bingoBoards = append(bingoBoards, *currentBoard)
			currentBoard = new(BingoBoard)
			currentBoard.winningNumber = -1
		}
		split := strings.Fields(v)
		for column, number := range split {
			intNumber, _ := strconv.Atoi(number)
			currentBoard.board[row%5][column] = intNumber
		}
	}
	return bingoBoards
}

func convertNumbersInput(numbersInput []string) []int {
	var numbers []int
	for _, v := range numbersInput {
		split := strings.Split(v, ",")
		for _, x := range split {
			i, _ := strconv.Atoi(x)
			numbers = append(numbers, i)
		}
	}
	return numbers
}
