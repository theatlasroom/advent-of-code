package main

import (
	"errors"
	"fmt"
	"sort"

	"github.com/theatlasroom/advent-of-code/go/2020/utils"
)

/**
--- Day 1: Report Repair ---

After saving Christmas five years in a row, you've decided to take a vacation at a nice resort on a tropical island. Surely, Christmas will go on without you.

The tropical island has its own currency and is entirely cash-only. The gold coins used there have a little picture of a starfish; the locals just call them stars.
None of the currency exchanges seem to have heard of them, but somehow, you'll need to find fifty of these coins by the time you arrive so you can pay the deposit on your room.

To save your vacation, you need to get all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first.
Each puzzle grants one star. Good luck!

Before you leave, the Elves in accounting just need you to fix your expense report (your puzzle input); apparently, something isn't quite adding up.

Specifically, they need you to find the two entries that sum to 2020 and then multiply those two numbers together.

For example, suppose your expense report contained the following:

1721
979
366
299
675
1456

In this list, the two entries that sum to 2020 are 1721 and 299. Multiplying them together produces 1721 * 299 = 514579, so the correct answer is 514579.

Of course, your expense report is much larger. Find the two entries that sum to 2020; what do you get if you multiply them together?
**/

// func isCandidate() {}

const target = 2020

func getValues(asc, desc []int) (int, int, error) {
	for _, a := range desc {
		for _, b := range asc {
			sum := a + b
			fmt.Printf("\na %v b %v sum %v", a, b, sum)
			if sum == target {
				return a, b, nil
			}
			if sum > target {
				break
			}
		}
	}
	return 0, 0, errors.New("No matching values")
}

func compute(data []int) int {
	sort.Ints(data)

	// gross
	desc := sort.IntSlice(data)
	sort.Sort(sort.Reverse(desc))
	fmt.Println(data)
	fmt.Println(desc)

	// a, b, err := getValues(data, desc)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// return a * b
}

func main() {
	// Read all the numbers
	input := utils.LoadDataAsString("1.txt")

	program := utils.StrToIntArr(input)
	compute(program)
}
