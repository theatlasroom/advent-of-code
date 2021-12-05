package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/theatlasroom/advent-of-code/go/utils"
)

/**
--- Day 5: Hydrothermal Venture ---

You come across a field of hydrothermal vents on the ocean floor! These vents constantly produce large, opaque clouds, so it would be best to avoid them if possible.

They tend to form in lines; the submarine helpfully produces a list of nearby lines of vents (your puzzle input) for you to review. For example:

0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2

Each line of vents is given as a line segment in the format x1,y1 -> x2,y2 where x1,y1 are the coordinates of one end the line segment and x2,y2 are the coordinates of the other end. These line segments include the points at both ends. In other words:

    An entry like 1,1 -> 1,3 covers points 1,1, 1,2, and 1,3.
    An entry like 9,7 -> 7,7 covers points 9,7, 8,7, and 7,7.

For now, only consider horizontal and vertical lines: lines where either x1 = x2 or y1 = y2.

So, the horizontal and vertical lines from the above list would produce the following diagram:

.......1..
..1....1..
..1....1..
.......1..
.112111211
..........
..........
..........
..........
222111....

In this diagram, the top left corner is 0,0 and the bottom right corner is 9,9. Each position is shown as the number of lines which cover that point or . if no line covers that point. The top-left pair of 1s, for example, comes from 2,2 -> 2,1; the very bottom row is formed by the overlapping lines 0,9 -> 5,9 and 0,9 -> 2,9.

To avoid the most dangerous areas, you need to determine the number of points where at least two lines overlap. In the above example, this is anywhere in the diagram with a 2 or larger - a total of 5 points.

Consider only horizontal and vertical lines. At how many points do at least two lines overlap?

*/

type point struct {
	X, Y int
}

type ventLines struct {
	Start point
	End   point
}

func (p *point) encode() string {
	return fmt.Sprintf("%d-%d", p.X, p.Y)
}

func (v *ventLines) isVertical() bool {
	if v.Start.X == v.End.X {
		return true
	}
	return false
}

func (v *ventLines) isHorizontal() bool {
	if v.Start.Y == v.End.Y {
		return true
	}
	return false
}

func parsePoint(s string) point {
	precision := 32
	coordDelimiter := ","
	sp := strings.Split(strings.TrimSpace(s), coordDelimiter)

	xpos, err := strconv.ParseInt(sp[0], 10, precision)
	if err != nil {
		xpos = 0
	}

	ypos, err := strconv.ParseInt(sp[1], 10, precision)
	if err != nil {
		ypos = 0
	}

	return point{
		X: int(xpos),
		Y: int(ypos),
	}
}

func parseVentInput(data []string) []ventLines {
	pointDelimiter := "->"
	var lines []ventLines
	for _, line := range data {
		points := strings.Split(line, pointDelimiter)

		l := ventLines{
			Start: parsePoint(points[0]),
			End:   parsePoint(points[1]),
		}

		lines = append(lines, l)
	}
	return lines
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x

}

func countPointsCovered(pc map[string]int) int {
	total := 0
	for _, count := range pc {
		if count > 1 {
			total += 1
		}
	}
	return total
}

func part1(lines []ventLines) {
	covered := make(map[string]int)

	for _, l := range lines {
		if l.isVertical() {
			delta := intAbs(l.Start.Y - l.End.Y)
			origin := l.Start
			if l.Start.Y > l.End.Y {
				origin = l.End
			}

			for i := 0; i <= delta; i += 1 {
				p := point{
					X: origin.X,
					Y: origin.Y + i,
				}

				key := p.encode()
				covered[key] += 1

			}
		}

		if l.isHorizontal() {
			delta := intAbs(l.Start.X - l.End.X)
			origin := l.Start
			if l.Start.X > l.End.X {
				origin = l.End
			}

			for i := 0; i <= delta; i += 1 {
				p := point{
					X: origin.X + i,
					Y: origin.Y,
				}
				key := p.encode()
				covered[key] = covered[key] + 1
			}
		}
	}

	total := countPointsCovered(covered)
	fmt.Printf("Part 1: points %d\n", total)
}

func main() {
	cfg := utils.BannerConfig{Year: 2021, Day: 5}
	utils.Banner(cfg)

	// Read all the numbers
	str := utils.LoadData("5.txt")
	lines := parseVentInput(str)

	part1(lines)
}
