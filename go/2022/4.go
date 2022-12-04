package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/theatlasroom/advent-of-code/go/utils"
)

/**
--- Day 4: Camp Cleanup ---

Space needs to be cleared before the last supplies can be unloaded from the ships, and so several Elves have been assigned the job of cleaning up sections of the camp. Every section has a unique ID number, and each Elf is assigned a range of section IDs.

However, as some of the Elves compare their section assignments with each other, they've noticed that many of the assignments overlap. To try to quickly find overlaps and reduce duplicated effort, the Elves pair up and make a big list of the section assignments for each pair (your puzzle input).

For example, consider the following list of section assignment pairs:

2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8

For the first few pairs, this list means:

    Within the first pair of Elves, the first Elf was assigned sections 2-4 (sections 2, 3, and 4), while the second Elf was assigned sections 6-8 (sections 6, 7, 8).
    The Elves in the second pair were each assigned two sections.
    The Elves in the third pair were each assigned three sections: one got sections 5, 6, and 7, while the other also got 7, plus 8 and 9.

This example list uses single-digit section IDs to make it easier to draw; your actual list might contain larger numbers. Visually, these pairs of section assignments look like this:

.234.....  2-4
.....678.  6-8

.23......  2-3
...45....  4-5

....567..  5-7
......789  7-9

.2345678.  2-8
..34567..  3-7

.....6...  6-6
...456...  4-6

.23456...  2-6
...45678.  4-8

Some of the pairs have noticed that one of their assignments fully contains the other. For example, 2-8 fully contains 3-7, and 6-6 is fully contained by 4-6. In pairs where one assignment fully contains the other, one Elf in the pair would be exclusively cleaning sections their partner will already be cleaning, so these seem like the most in need of reconsideration. In this example, there are 2 such pairs.

In how many assignment pairs does one range fully contain the other?
**/

type assignment struct{ start, end int }
type assignmentComparator = func(assignment, assignment) bool

const (
	DELIM_ASSIGNMENT = "-"
	DELIM_PAIR       = ","
)

func newAssignment(str string) assignment {
	s := strings.Split(str, DELIM_ASSIGNMENT)
	start, startok := strconv.Atoi(s[0])
	if startok != nil {
		start = 0
	}

	end, endok := strconv.Atoi(s[1])
	if endok != nil {
		end = 0
	}

	return assignment{
		start: start,
		end:   end,
	}
}

func fullyContains(a, b assignment) bool {
	if (a.start <= b.start && a.end >= b.end) ||
		(b.start <= a.start && b.end >= a.end) {
		return true
	}

	return false
}

func overlaps(a, b assignment) bool {
	if fullyContains(a, b) {
		return true
	}

	if (b.start <= a.end && b.start >= a.start) ||
		(a.start <= b.end && a.start >= b.start) {
		return true
	}

	return false
}

func parseAssignmentPairs(data []string, compfn assignmentComparator) int {
	sum := 0
	for _, line := range data {
		d := strings.Split(line, DELIM_PAIR)
		l, r := d[0], d[1]

		al := newAssignment(l)
		ar := newAssignment(r)

		if compfn(al, ar) {
			sum++
		}
	}
	return sum
}

func part2(data []string) {
	sum := parseAssignmentPairs(data, overlaps)
	fmt.Println("Part 2: Sum", sum)
}

func part1(data []string) {
	sum := parseAssignmentPairs(data, fullyContains)
	fmt.Println("Part 1: Sum", sum)
}

func main() {
	// Read all the numbers
	data := utils.LoadData("4.txt")
	cfg := utils.BannerConfig{Year: 2022, Day: 4}
	utils.Banner(cfg)

	part1(data)
	part2(data)
}
