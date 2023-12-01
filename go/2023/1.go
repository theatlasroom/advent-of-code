package main

import (
	"fmt"
	"strconv"

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

func calculateCalibrationValue(a, b int) int {
	str := fmt.Sprintf("%d%d", a, b)
	v, err := strconv.Atoi(str)
	utils.CheckAndPanic(err)
	return v
}

func calibrate(str string) int {
	ints := utils.FindDigitsInString(str)

	switch l := len(ints); {
	case l > 1:
		return calculateCalibrationValue(ints[0], ints[l-1])
	case l == 1:
		return calculateCalibrationValue(ints[0], ints[0])
	default:
		return 0
	}
}

func parseInput(data []string, index int, val int) int {
	if index < len(data) {
		sum := val + calibrate(data[index])
		return parseInput(data, index+1, sum)
	}
	return val
}

func part1(data []string) {
	val := parseInput(data, 0, 0)
	utils.PrintResult(1, fmt.Sprintf("sum %d", val))
}

func main() {
	utils.Banner(utils.BannerConfig{Year: 2023, Day: 1})
	// Read all the numbers
	input := utils.LoadData("1.txt")
	// fmt.Println("input", input)

	part1(input)
	// part2(data)
}
