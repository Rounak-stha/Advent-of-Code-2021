package main

import (
	"AOC/file_paths"
	"AOC/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cardAndScore struct{ 
	score int64 
	card [][]string 
}

type Board struct {
	cells  [][]string
	called [][]bool
}

func (b *Board) getScore(mutiplier int64) int64 {
	var score int64
	for i, row := range b.cells {
		for j, cell := range row {
			if b.called[i][j] != true {
				num, err := strconv.Atoi(cell)
				utils.HandleError(err)
				score += int64(num)
			}
		}
	}
	return score * mutiplier
}
func (b *Board) checkBingo(row, col int) bool {
	rows := len(b.called)
	columns := len(b.called[0])
	rowBingo := true
	columnBingo := true
	for i := 0; i < rows; i++ {
		if b.called[i][col] == false {
			columnBingo = false
			break
		}
	}
	for i := 0; i < columns; i++ {
		if b.called[row][i] == false {
			rowBingo = false
			break
		}
	}

	return rowBingo || columnBingo
}

func (b *Board) checkNumAndBingo(num string) bool {
	bingo := false
	for i, row := range b.cells {
		for j, cell := range row {
			if cell == num {
				b.called[i][j] = true
				bingo = b.checkBingo(i, j)
			}
		}
	}
	return bingo
}


func createBoard(rawData string) Board {
	lines := strings.Split(rawData, "\r\n")
	if len(lines[len(lines) - 1]) == 0 {
		lines = lines[:len(lines) - 1]
	}
	card := make([][]string, len(lines), len(lines))
	cardCalled := make([][]bool, len(lines), len(lines))
	for i, line := range lines {
		nums := strings.Fields(line)
		card[i] = nums
		cardCalled[i] = make([]bool, len(nums), len(nums))
	} 
	return Board{ card, cardCalled}
}

func parsedData(data []byte) ([]string, []Board) {
	stringifiedData := string(data)
	parsedData := strings.Split(stringifiedData, "\r\n\r\n")
	var boards []Board
	gameInput := strings.Split(parsedData[0], ",")
	for _, card := range parsedData[1:] {
		board := createBoard(card)
		boards = append(boards, board)
	}

	return gameInput, boards
}

func main() {
	data, err := os.ReadFile(file_paths.Day_4)
	utils.HandleError(err)
	gameInput, boards := parsedData(data)
	// fmt.Println(boards)
	gameResult := playAndGetScore(gameInput, boards)
	fmt.Println("The First Bingo Card to win is ", gameResult[0].card, " with an score of ", gameResult[0].score)
	fmt.Println("The Last Bingo Card to win is ", gameResult[1].card, " with an score of ", gameResult[1].score)
}

func playAndGetScore(gameInput []string, boards []Board) []cardAndScore {
	// var scores []int64
	firstAndLast := make([]cardAndScore, 0, 2)
	// var cardsAndScores []cardAndScore
	var bingo bool
	for _, val := range gameInput {
		bingoedBoardIndicies := []int{}
		for i, board := range boards {
			bingo = board.checkNumAndBingo(val)
			if bingo {
				val, err := strconv.Atoi(val)
				utils.HandleError(err)
				score := board.getScore(int64(val))
				// scores = append(scores, score)
				// cardsAndScores = append(cardsAndScores,  cardAndScore { score, board.cells })
				bingoedBoardIndicies = append(bingoedBoardIndicies, i)
				if len(firstAndLast) == 0 {
					firstAndLast = append(firstAndLast, cardAndScore { score, board.cells })
				}
				if len(boards) == 1 {
					firstAndLast = append(firstAndLast, cardAndScore { score, board.cells })
				}
			}
		}

		for i := len(bingoedBoardIndicies)-1; i >= 0; i-- {
			boards[bingoedBoardIndicies[i]] = boards[len(boards) - 1]
			boards = boards[:len(boards) - 1]
		}
	}

	return firstAndLast
} 