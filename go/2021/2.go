package main

import (
	"fmt"
	"strconv"
	s "strings"

	"github.com/theatlasroom/advent-of-code/go/utils"
)

/**
--- Day 2: Dive! ---

Now, you need to figure out how to pilot this thing.

It seems like the submarine can take a series of commands like forward 1, down 2, or up 3:

    forward X increases the horizontal position by X units.
    down X increases the depth by X units.
    up X decreases the depth by X units.

Note that since you're on a submarine, down and up affect your depth, and so they have the opposite result of what you might expect.

The submarine seems to already have a planned course (your puzzle input). You should probably figure out where it's going. For example:

forward 5
down 5
forward 8
up 3
down 8
forward 2

Your horizontal position and depth both start at 0. The steps above would then modify them as follows:

    forward 5 adds 5 to your horizontal position, a total of 5.
    down 5 adds 5 to your depth, resulting in a value of 5.
    forward 8 adds 8 to your horizontal position, a total of 13.
    up 3 decreases your depth by 3, resulting in a value of 2.
    down 8 adds 8 to your depth, resulting in a value of 10.
    forward 2 adds 2 to your horizontal position, a total of 15.

After following these instructions, you would have a horizontal position of 15 and a depth of 10. (Multiplying these together produces 150.)

Calculate the horizontal position and depth you would have after following the planned course. What do you get if you multiply your final horizontal position by your final depth?

*/

type vector struct {
	XPos, ZPos int
}

type direction int

const (
	DIR_FORWARD direction = iota
	DIR_UP
	DIR_DOWN
)

type instruction struct {
	Direction direction
	Delta     int
}

func (v *vector) Forward(delta int) {
	v.XPos += delta
}

func (v *vector) Up(delta int) {
	v.ZPos -= delta
}

func (v *vector) Down(delta int) {
	v.ZPos += delta
}

func (v vector) Position() int {
	return v.ZPos * v.XPos
}

func parseInstruction(str string) instruction {
	strArr := s.Split(str, " ")
	var d direction
	switch strArr[0] {
	case "down":
		d = DIR_DOWN
	case "up":
		d = DIR_UP
	default:
		d = DIR_FORWARD
	}

	delta, err := strconv.Atoi(strArr[1])
	if err != nil {
		delta = 0
	}
	return instruction{
		Direction: d,
		Delta:     delta,
	}
}

func part1(data []string) {
	var submarine vector

	for _, line := range data {
		i := parseInstruction(line)
		switch i.Direction {
		case DIR_DOWN:
			submarine.Down(i.Delta)
		case DIR_UP:
			submarine.Up(i.Delta)
		default:
			submarine.Forward(i.Delta)
		}
	}

	fmt.Printf("Part 1: Final position %d\n", submarine.Position())
}

func main() {
	cfg := utils.BannerConfig{Year: 2021, Day: 2}
	utils.Banner(cfg)

	// Read all the numbers
	input := utils.LoadData("2.txt")

	part1(input)
}
