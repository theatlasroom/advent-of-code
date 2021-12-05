package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/theatlasroom/advent-of-code/go/utils"
)

/**
--- Day 4: Giant Squid ---

You're already almost 1.5km (almost a mile) below the surface of the ocean, already so deep that you can't see any sunlight. What you can see, however, is a giant squid that has attached itself to the outside of your submarine.

Maybe it wants to play bingo?

Bingo is played on a set of boards each consisting of a 5x5 grid of numbers. Numbers are chosen at random, and the chosen number is marked on all boards on which it appears. (Numbers may not appear on all boards.) If all numbers in any row or any column of a board are marked, that board wins. (Diagonals don't count.)

The submarine has a bingo subsystem to help passengers (currently, you and the giant squid) pass the time. It automatically generates a random order in which to draw numbers and a random set of boards (your puzzle input). For example:

7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7

After the first five numbers are drawn (7, 4, 9, 5, and 11), there are no winners, but the boards are marked as follows (shown here adjacent to each other to save space):

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7

After the next six numbers are drawn (17, 23, 2, 0, 14, and 21), there are still no winners:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7

Finally, 24 is drawn:

22 13 17 11  0         3 15  0  2 22        14 21 17 24  4
 8  2 23  4 24         9 18 13 17  5        10 16 15  9 19
21  9 14 16  7        19  8  7 25 23        18  8 23 26 20
 6 10  3 18  5        20 11 10 24  4        22 11 13  6  5
 1 12 20 15 19        14 21 16 12  6         2  0 12  3  7

At this point, the third board wins because it has at least one complete row or column of marked numbers (in this case, the entire top row is marked: 14 21 17 24 4).

The score of the winning board can now be calculated. Start by finding the sum of all unmarked numbers on that board; in this case, the sum is 188. Then, multiply that sum by the number that was just called when the board won, 24, to get the final score, 188 * 24 = 4512.

To guarantee victory against the giant squid, figure out which board will win first. What will your final score be if you choose that board?
*/

type board = []int
type bingo struct {
	Numbers []int
	Boards  []board
}

func hasWon(values map[int]bool, arr []int) bool {
	won := true
	for _, v := range arr {
		_, ok := values[v]
		if !ok {
			won = false
			break
		}
	}
	return won
}

func checkRows(values map[int]bool, b board) (bool, []int) {
	rows := 5
	cols := 5
	for i := 0; i < rows; i += 1 {
		start := i * cols
		end := start + cols
		arr := b[start:end]
		won := hasWon(values, arr)

		if won {
			return true, arr
		}
	}
	return false, nil
}

func checkColumns(values map[int]bool, b board) (bool, []int) {
	cols := 5
	boardLength := len(b)
	for i := 0; i < cols; i += 1 {
		var arr []int

		for j := i; j < boardLength; j += cols {
			arr = append(arr, b[j])
		}

		won := hasWon(values, arr)

		if won {
			return true, arr
		}
	}
	return false, nil
}

func hasBingo(values map[int]bool, b board) (bool, []int) {

	// check rows
	won, winningRow := checkRows(values, b)
	if won {
		return won, winningRow
	}

	// check columns
	won, winningCol := checkColumns(values, b)
	if won {
		return won, winningCol
	}
	return false, nil
}

func checkAllBoards(values map[int]bool, bs []board) (bool, board, int) {
	for boardIndex, b := range bs {
		res, _ := hasBingo(values, b)
		if res {
			return true, b, boardIndex
		}
	}
	return false, nil, 0
}

func playGame(game bingo) (bool, board, map[int]bool, int) {
	values := make(map[int]bool)
	for idx, num := range game.Numbers {
		values[num] = true
		if idx > 5 {
			complete, board, _ := checkAllBoards(values, game.Boards)
			if complete {
				return true, board, values, num
			}
		}
	}
	return false, nil, nil, 0
}

type winningBoard struct {
	Data   []int
	Number int
}

func remove(slice []board, index int) []board {
	return append(slice[:index], slice[index+1:]...)
}

func cloneMap(m map[int]bool) map[int]bool {
	clone := make(map[int]bool)
	for key, value := range m {
		clone[key] = value
	}
	return clone
}

func inputDataKeys(in map[int]bool) []int {
	var keys []int
	for k, _ := range in {
		keys = append(keys, k)
	}
	return keys
}

func playGameUntilEnd(game bingo) (board, map[int]bool, int) {
	var currValues map[int]bool
	var winner board
	var finalNumber int

	candidateBoards := game.Boards
	hasWinner := true

	for hasWinner == true {
		hasWinner = false
		values := make(map[int]bool)
		for idx, num := range game.Numbers {
			values[num] = true
			if idx > 5 && len(candidateBoards) > 0 {
				complete, board, boardIndex := checkAllBoards(values, candidateBoards)
				if complete {
					winner = board
					finalNumber = num
					candidateBoards = remove(candidateBoards, boardIndex)
					currValues = cloneMap(values)
					hasWinner = true
					break
				}
			}
		}
	}
	return winner, currValues, finalNumber
}

func calculateScore(b board, marked map[int]bool, finalNumber int) int {
	sum := 0
	for _, v := range b {
		_, isMarked := marked[v]
		if !isMarked {
			sum += v
		}
	}
	return sum * finalNumber
}

func parseBoard(input []string) board {
	var b board
	for _, num := range input {
		v, err := strconv.Atoi(num)
		if err != nil {
			v = 0
		}
		b = append(b, v)
	}
	return b
}

func parseNumbers(line string) []int {
	var arr []int
	s := strings.Split(line, ",")
	for _, str := range s {
		v, err := strconv.Atoi(string(str))
		if err == nil {
			arr = append(arr, v)
		}
	}
	return arr
}

func linesToFlatArray(str string) []string {
	return strings.Split(strings.Replace(strings.TrimSpace((str)), "  ", " ", -1), " ")
}

func combineLinesIntoString(str string) string {
	return strings.Join(strings.Split(str, "\n"), " ")
}

func parseBingoInput(in string) bingo {
	arrs := strings.Split(in, "\n\n")
	var nums []int
	var boards []board

	for idx, input := range arrs {
		if idx == 0 {
			nums = parseNumbers(input)
		} else {
			if len(input) > 0 {
				boards = append(boards, parseBoard(linesToFlatArray(combineLinesIntoString(input))))
			}
		}
	}
	return bingo{Numbers: nums, Boards: boards}
}

func part2(game bingo) {
	b, marked, finalNumber := playGameUntilEnd(game)
	score := calculateScore(b, marked, finalNumber)
	fmt.Println("Part 2: Bingo", score)
}

func part1(game bingo) {
	ok, b, marked, finalNumber := playGame(game)
	score := 0
	if ok {
		score = calculateScore(b, marked, finalNumber)
	}
	fmt.Println("Part 1: Bingo", score)
}

func main() {
	cfg := utils.BannerConfig{Year: 2021, Day: 4}
	utils.Banner(cfg)

	// Read all the numbers
	str := utils.LoadDataAsString("4.txt")
	game := parseBingoInput(str)

	part1(game)
	part2(game)
}
