package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BingoNumber struct {
	Number int
	Drawn  bool
}

type Bingo struct {
	Card [5][5]BingoNumber
}

func (b Bingo) HasBingo() bool {
	for _, row := range b.Card {
		// Check each row
		rowBingo := true
		for _, columnVal := range row {
			if !columnVal.Drawn {
				rowBingo = false
				break
			}
		}
		if rowBingo {
			return true
		}
	}

	for colNum, _ := range b.Card[0] {
		// check each colunn
		colBingo := true
		for i := 0; i < 5; i++ {
			if !b.Card[i][colNum].Drawn {
				colBingo = false
				break
			}
		}
		if colBingo {
			return true
		}
	}

	return false
}

func (b Bingo) CalculatePoints() int {
	count := 0
	for _, row := range b.Card {
		for _, col := range row {
			if !col.Drawn {
				count += col.Number
			}
		}
	}
	return count
}

func (b *Bingo) MarkNum(number int) bool {
	for i, row := range b.Card {
		for j, col := range row {
			// b.Card[i][j].Drawn = true
			if col.Number == number {
				b.Card[i][j].Drawn = true
				// col.Drawn = true
				// b.Card[i][j] = col
				return true
			}
		}
	}
	return false
}

func main() {
	// Main for part 2
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	input := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	numberDraw := strings.Split(input[0], ",")

	boards := ParseBoards(input[2:])

	lastNum := 0
	for _, drawStr := range numberDraw {
		drawNum, _ := strconv.Atoi(drawStr)
		lastNum = drawNum
		for x, _ := range boards {
			boards[x].MarkNum(drawNum)
		}
		fmt.Printf("PreFilter %d\n", len(boards))
		if len(boards) == 1 {
			break
		}
		boards = FilterBingo(boards)
		fmt.Printf("Post %d\n", len(boards))

	}

	// fmt.Println(numberDraw)
	// fmt.Println(boards)
	fmt.Printf("Score: %d, lastNum %d, result %d\n",
		boards[0].CalculatePoints(), lastNum, lastNum*boards[0].CalculatePoints())
}

func FilterBingo(boards []Bingo) []Bingo {
	n := 0
	for _, board := range boards {
		if !board.HasBingo() {
			boards[n] = board
			n++
		}
	}
	boards = boards[:n]
	return boards
}

func main_part1() {
	// Main for part 1
	f, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	input := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	numberDraw := strings.Split(input[0], ",")
	// numberDraw := []string{"57"}

	boards := ParseBoards(input[2:])

	bingo := false
	lastNum := 0
	bingoIndex := 0
	for _, drawStr := range numberDraw {
		drawNum, _ := strconv.Atoi(drawStr)
		lastNum = drawNum
		for x, _ := range boards {
			if boards[x].MarkNum(drawNum) {
				if boards[x].HasBingo() {
					fmt.Println("BINGO")
					fmt.Println(boards[x])
					bingo = true
					bingoIndex = x
					break
				}
			}
		}
		if bingo {
			break
		}
	}

	// fmt.Println(numberDraw)
	// fmt.Println(boards)
	fmt.Printf("Score: %d, lastNum %d, result %d\n",
		boards[bingoIndex].CalculatePoints(), lastNum, lastNum*boards[bingoIndex].CalculatePoints())
}

func ParseBoards(lines []string) []Bingo {
	// assuming we start at the first board
	boards := []Bingo{}

	for i := 0; i < len(lines); i += 6 {
		nums := [5][5]BingoNumber{}
		for j, cardLine := range lines[i : i+5] {
			line := [5]BingoNumber{}
			lineNumbers := strings.Fields(cardLine)
			for k, number := range lineNumbers {
				num, _ := strconv.Atoi(number)
				line[k] = BingoNumber{Number: num, Drawn: false}
			}
			nums[j] = line
		}
		boards = append(boards, Bingo{Card: nums})
		// break
	}

	return boards
}
