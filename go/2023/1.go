package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/theatlasroom/advent-of-code/go/utils"
)

/*
--- Day 1: Trebuchet?! ---

Something is wrong with global snow production, and you've been selected to take a look.
The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar;
the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

You try to ask why they can't just use a weather machine ("not powerful enough") and where they're even sending you ("the sky")
and why your map looks mostly blank ("you sure ask a lot of questions") and hang on did you just say the sky
("of course, where do you think snow comes from") when you realize that the Elves are already loading you into a trebuchet
("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been
amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are
having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration
value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit
(in that order) to form a single two-digit number.

For example:

1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet

In this example, the calibration values of these four lines are 12, 38, 15, and 77. Adding these together produces 142.

Consider your entire calibration document. What is the sum of all of the calibration values?

*/
var digitStrings = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func findFirstDigit(target, sVal, siVal string) int {
	fStr, fiStr := strings.Index(target, sVal), strings.Index(target, siVal)

	if fStr > -1 && fiStr > -1 {
		if fStr < fiStr {
			return fStr
		} else {
			return fiStr
		}
	}

	if fStr > -1 {
		return fStr
	}

	return fiStr
}

func findLastDigit(target, sVal, siVal string) int {
	lStr, liStr := strings.LastIndex(target, sVal), strings.LastIndex(target, siVal)

	if lStr > -1 && liStr > -1 {
		if lStr > liStr {
			return lStr
		} else {
			return liStr
		}
	}

	if lStr > -1 {
		return lStr
	}

	return liStr
}

func naiveParseString(currString string) []int {
	first, last := 0, 0
	fidx, lidx := -1, -1

	for key, value := range digitStrings {
		nfidx := findFirstDigit(currString, key, strconv.Itoa(value))
		if fidx < 0 {
			fidx = nfidx
		}

		if nfidx > -1 && nfidx <= fidx {
			first = value
			fidx = nfidx
		}

		nlidx := findLastDigit(currString, key, strconv.Itoa(value))
		if lidx < 0 {
			lidx = nlidx
		}

		if nlidx > -1 && nlidx >= lidx {
			last = value
			lidx = nlidx
		}
	}

	return []int{first, last}
}

type finderFn = func(string) []int

func findPossibleDigits(str string) []int {
	// Note: Workarounds because positive/negative lookaheads are not supported
	// Link - https://github.com/golang/go/issues/18868

	re := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	// split via the regex and return all substrings
	result := re.FindAllString(str, -1)
	fmt.Println("result", result)

	var ints []int
	for _, str := range result {
		if val, ok := digitStrings[str]; ok {
			ints = append(ints, val)
		} else {
			i, err := strconv.Atoi(str)
			utils.CheckAndPanic(err)

			ints = append(ints, i)
		}
	}

	fmt.Print("ints", ints)
	return ints
}

func findDigits(str string) []int {
	re := regexp.MustCompile(`(\d)`)
	// split via the regex and return all substrings
	result := re.FindAllString(str, -1)

	var ints []int

	for _, str := range result {
		i, err := strconv.Atoi(str)
		utils.CheckAndPanic(err)

		ints = append(ints, i)

	}
	return ints
}

func calculateCalibrationValue(a, b int) int {
	str := fmt.Sprintf("%d%d", a, b)
	v, err := strconv.Atoi(str)
	utils.CheckAndPanic(err)

	// fmt.Printf(" -> %d\n", v)
	return v
}

func calibrate(str string, fn finderFn) int {
	ints := fn(str)

	switch l := len(ints); {
	case l > 1:
		return calculateCalibrationValue(ints[0], ints[l-1])
	case l == 1:
		return calculateCalibrationValue(ints[0], ints[0])
	default:
		return 0
	}
}

func parseInput(data []string, index int, val int, fn finderFn) int {
	if index < len(data) {
		sum := val + calibrate(data[index], fn)
		return parseInput(data, index+1, sum, fn)
	}
	return val
}

func part2(data []string) {
	val := parseInput(data, 0, 0, naiveParseString)
	// val := parseInput(data, 0, 0, findPossibleDigits)
	utils.PrintResult(2, fmt.Sprintf("sum %d", val))
}

func part1(data []string) {
	val := parseInput(data, 0, 0, findDigits)
	utils.PrintResult(1, fmt.Sprintf("sum %d", val))
}

func main() {
	utils.Banner(utils.BannerConfig{Year: 2023, Day: 1})
	// Read all the calibration values
	input := utils.LoadData("1.txt")

	part1(input)
	part2(input)
}
