package main

import (
	"fmt"
	"strconv"

	"github.com/theatlasroom/advent-of-code/go/utils"
)

/**
--- Day 3: Binary Diagnostic ---

The submarine has been making some odd creaking noises, so you ask it to produce a diagnostic report just in case.

The diagnostic report (your puzzle input) consists of a list of binary numbers which, when decoded properly, can tell you many useful things about the conditions of the submarine. The first parameter to check is the power consumption.

You need to use the binary numbers in the diagnostic report to generate two new binary numbers (called the gamma rate and the epsilon rate). The power consumption can then be found by multiplying the gamma rate by the epsilon rate.

Each bit in the gamma rate can be determined by finding the most common bit in the corresponding position of all numbers in the diagnostic report. For example, given the following diagnostic report:

00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010

Considering only the first bit of each number, there are five 0 bits and seven 1 bits. Since the most common bit is 1, the first bit of the gamma rate is 1.

The most common second bit of the numbers in the diagnostic report is 0, so the second bit of the gamma rate is 0.

The most common value of the third, fourth, and fifth bits are 1, 1, and 0, respectively, and so the final three bits of the gamma rate are 110.

So, the gamma rate is the binary number 10110, or 22 in decimal.

The epsilon rate is calculated in a similar way; rather than use the most common bit, the least common bit from each position is used. So, the epsilon rate is 01001, or 9 in decimal. Multiplying the gamma rate (22) by the epsilon rate (9) produces the power consumption, 198.

Use the binary numbers in your diagnostic report to calculate the gamma rate and epsilon rate, then multiply them together. What is the power consumption of the submarine? (Be sure to represent your answer in decimal, not binary.)
*/

func generateMask(maskLength int) uint32 {
	var str string
	for len(str) < maskLength {
		str += "1"
	}
	return binaryStringAsUint32(str)
}

func flipBits(v uint32, mask uint32) uint32 {
	// bitwise negate the value v, then perform a bitwise AND using the provided mask
	// given v = 1100100010, mask 0b111111111111, returns 110011011101
	return ^v & mask
}

func binaryStringAsUint32(s string) uint32 {
	v, err := strconv.ParseUint(s, 2, 32)
	if err != nil {
		return 0
	}
	return uint32(v)
}

func calculatePower(counts []int, threshhold int, mask uint32) uint32 {
	str := generateRates(counts, threshhold, mask)
	gamma := binaryStringAsUint32(str)
	epsilon := flipBits(gamma, mask)
	return gamma * epsilon
}

func generateRates(counts []int, threshhold int, mask uint32) string {
	str := ""
	for _, i := range counts {
		if i > threshhold {
			str += "1"
			continue
		}
		str += "0"
	}
	return str
}

type lifeSupportRating struct {
	Zeros, Ones []string
}

type lifeSupportComparatorFn = func(lfr lifeSupportRating) []string

func matchCriteria(input []string, bit int, comparator lifeSupportComparatorFn) string {
	if len(input) == 1 {
		return input[0]
	}
	var zeros []string
	var ones []string

	for _, str := range input {
		if str[bit] == '1' {
			ones = append(ones, str)
		} else {
			zeros = append(zeros, str)
		}
	}

	rating := lifeSupportRating{Zeros: zeros, Ones: ones}
	return matchCriteria(comparator(rating), bit+1, comparator)
}

func part2(data []string) {
	oxygen := matchCriteria(data, 0, func(lfr lifeSupportRating) []string {
		if len(lfr.Ones) >= len(lfr.Zeros) {
			return lfr.Ones
		}
		return lfr.Zeros
	})

	c02 := matchCriteria(data, 0, func(lfr lifeSupportRating) []string {
		if len(lfr.Ones) < len(lfr.Zeros) {
			return lfr.Ones
		}
		return lfr.Zeros
	})
	fmt.Printf("Part 2: life support rating %v\n", binaryStringAsUint32(oxygen)*binaryStringAsUint32(c02))
}

func part1(data []string) {
	size := 12
	mask := generateMask(size)
	threshhold := len(data) / 2

	diagnostics := make([]int, size, size)

	for _, str := range data {
		for i, c := range str {
			if c == '1' {
				diagnostics[i] += 1
			}
		}
	}

	power := calculatePower(diagnostics[:], threshhold, mask)
	fmt.Printf("Part 1: power consumption %d\n", power)
}

func main() {
	cfg := utils.BannerConfig{Year: 2021, Day: 3}
	utils.Banner(cfg)

	// Read all the numbers
	input := utils.LoadData("3.txt")

	part1(input)
	part2(input)
}
