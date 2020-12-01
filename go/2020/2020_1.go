package main

import (
	"errors"
	"fmt"
	"log"
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

const target = 2020

type pair struct {
	a, b int
}

func (p pair) Product() int {
	return p.a * p.b
}

func (p pair) toString() string {
	return fmt.Sprintf("%v %v %v", p.a, p.b, p.Product())
}

type triplet struct {
	pair // embeds the pair struct
	c    int
}

func (t triplet) Product() int {
	return t.a * t.b * t.c
}

func (t triplet) toString() string {
	return fmt.Sprintf("%v %v %v %v", t.a, t.b, t.c, t.Product())
}

func equalPairs(a int, asc []int) (int, bool) {
	for _, b := range asc {
		sum := a + b
		if sum > target {
			break
		}
		if sum == target {
			return b, true
		}
	}
	return 0, false
}

func findPairsEqualToTarget(asc, desc []int) (pair, error) {
	for _, a := range desc {
		b, ok := equalPairs(a, asc)
		if ok {
			return pair{a, b}, nil
		}
	}
	return pair{}, errors.New("No matching values")
}

func findPairsLessThanTarget(asc, desc []int) []pair {
	var pairs []pair
	for _, a := range desc {
		for _, b := range asc {
			sum := a + b
			if sum >= target {
				break
			}
			pairs = append(pairs, pair{a, b})
		}
	}
	return pairs
}

func part1(asc, desc []int) pair {
	p, err := findPairsEqualToTarget(asc, desc)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func part2(asc, desc []int) triplet {
	p := findPairsLessThanTarget(asc, desc)
	var t triplet
	for _, pv := range p {
		curr := pv.a + pv.b
		if curr >= target {
			continue
		}

		c, ok := equalPairs(curr, asc)

		if ok {
			return triplet{pv, c}
		}
	}

	if t == (triplet{}) {
		log.Fatal(errors.New("No matching values"))
	}
	return t
}

func main() {
	// Read all the numbers
	utils.Banner(1)
	input := utils.LoadDataAsString("1.txt")

	data := utils.StrToIntArr(input)

	// gross
	sort.Ints(data)
	desc := sort.IntSlice(append([]int(nil), data...))
	sort.Sort(sort.Reverse(desc))

	fmt.Println(part1(data, desc).toString())
	fmt.Println(part2(data, desc).toString())
}
