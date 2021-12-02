package main

import (
	"errors"
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/theatlasroom/advent-of-code/go/2019/utils"
)

/**
--- Day 5: Binary Boarding ---

You board your plane only to discover a new problem: you dropped your boarding pass!
You aren't sure which seat is yours, and all of the flight attendants are busy with the
flood of people that suddenly made it through passport control.

You write a quick program to use your phone's camera to scan all of the nearby
boarding passes (your puzzle input); perhaps you can find your seat through process of elimination.

Instead of zones or groups, this airline uses binary space partitioning to seat people.
A seat might be specified like FBFBBFFRLR, where F means "front", B means "back", L means "left", and R means "right".

The first 7 characters will either be F or B; these specify exactly one of the 128 rows on
the plane (numbered 0 through 127). Each letter tells you which half of a region the given seat is in.
Start with the whole list of rows; the first letter indicates whether the seat is in the front (0 through 63)
or the back (64 through 127).
The next letter indicates which half of that region the seat is in, and so on until you're left with exactly one row.

For example, consider just the first seven characters of FBFBBFFRLR:

    Start by considering the whole range, rows 0 through 127.
    F means to take the lower half, keeping rows 0 through 63.
    B means to take the upper half, keeping rows 32 through 63.
    F means to take the lower half, keeping rows 32 through 47.
    B means to take the upper half, keeping rows 40 through 47.
    B keeps rows 44 through 47.
    F keeps rows 44 through 45.
    The final F keeps the lower of the two, row 44.

The last three characters will be either L or R; these specify exactly one of the 8 columns of seats on the plane (numbered 0 through 7).
The same process as above proceeds again, this time with only three steps. L means to keep the lower half, while R means to keep the upper half.

For example, consider just the last 3 characters of FBFBBFFRLR:

    Start by considering the whole range, columns 0 through 7.
    R means to take the upper half, keeping columns 4 through 7.
    L means to take the lower half, keeping columns 4 through 5.
    The final R keeps the upper of the two, column 5.

So, decoding FBFBBFFRLR reveals that it is the seat at row 44, column 5.

Every seat also has a unique seat ID: multiply the row by 8, then add the column. In this example, the seat has ID 44 * 8 + 5 = 357.

Here are some other boarding passes:

    BFFFBBFRRR: row 70, column 7, seat ID 567.
    FFFBBBFRRR: row 14, column 7, seat ID 119.
    BBFFBBFRLL: row 102, column 4, seat ID 820.

As a sanity check, look through your list of boarding passes. What is the highest seat ID on a boarding pass?

*/
type seat struct {
	seatID, row, column int
}

type seats []seat

// implement the sort.Interface
func (s seats) Len() int           { return len(s) }
func (s seats) Less(i, j int) bool { return s[i].seatID < s[j].seatID }
func (s seats) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s seat) toString() string {
	return fmt.Sprintf("Row: %v Column: %v SeatID: %v", s.row, s.column, s.seatID)
}

func generateSeatID(row, column int) int {
	return (row * 8) + column
}

func parse(str, leftSymbol, rightSymbol string) int {
	var char string
	bits := ""
	for _, c := range str {
		char = string(c)
		switch char {
		case leftSymbol:
			bits += "0"
			break
		case rightSymbol:
			bits += "1"
			break
		default:
			break
		}
	}
	i, err := strconv.ParseInt(bits, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return int(i)
}

func findSeat(str string) seat {
	row := parse(str[:7], "F", "B")
	column := parse(str[7:], "L", "R")
	seatID := generateSeatID(row, column)
	return seat{seatID, row, column}
}

func calculateAllSeats(data []string) seats {
	var s seats
	for _, str := range data {
		s = append(s, findSeat(str))
	}
	return s
}

func part01(data []string) seat {
	s := calculateAllSeats(data)
	sort.Sort(sort.Reverse(s))
	return s[0]
}

func part02(data []string) (int, error) {
	s := calculateAllSeats(data)
	sort.Sort(s)
	prevSeatID := s[0].seatID
	for i := 1; i < s.Len(); i++ {
		if s[i].seatID-prevSeatID > 1 {
			return ((s[i].seatID + prevSeatID) / 2), nil
		}
		prevSeatID = s[i].seatID
	}
	return -1, errors.New("Couldnt find the seat")
}

func main() {
	utils.Banner(utils.BannerConfig{Year: 2020, Day: 5})
	data := utils.LoadData("5.txt")
	fmt.Println(part01(data).toString())
	seatID, err := part02(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Missing seat", seatID)
}
