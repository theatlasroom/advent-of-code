package main

import (
	"fmt"

	"github.com/theatlasroom/advent-of-code/go/2020/utils"
)

/**
--- Day 3: Toboggan Trajectory ---

With the toboggan login problems resolved, you set off toward the airport.
While travel by toboggan might be easy, it's certainly not safe: there's very minimal steering and the area is covered in trees.
You'll need to see which angles will take you near the fewest trees.

Due to the local geology, trees in this area only grow on exact integer coordinates in a grid.
You make a map (your puzzle input) of the open squares (.) and trees (#) you can see. For example:

..##.......
#...#...#..
.#....#..#.
..#.#...#.#
.#...##..#.
..#.##.....
.#.#.#....#
.#........#
#.##...#...
#...##....#
.#..#...#.#

These aren't the only trees, though; due to something you read about once involving arboreal
genetics and biome stability, the same pattern repeats to the right many times:

..##.........##.........##.........##.........##.........##.......  --->
#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
.#....#..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..#...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.##.......#.##.......#.##.......#.##.......#.##.....  --->
.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........#.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...##....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#.#..#...#.#  --->

You start on the open square (.) in the top-left corner and need to reach the bottom (below the bottom-most row on your map).

The toboggan can only follow a few specific slopes (you opted for a cheaper model that prefers rational numbers);
start by counting all the trees you would encounter for the slope right 3, down 1:

From your starting position at the top-left, check the position that is right 3 and down 1.
Then, check the position that is right 3 and down 1 from there, and so on until you go past the bottom of the map.

The locations you'd check in the above example are marked here with O where there was an open square and X where there was a tree:

..##.........##.........##.........##.........##.........##.......  --->
#..O#...#..#...#...#..#...#...#..#...#...#..#...#...#..#...#...#..
.#....X..#..#....#..#..#....#..#..#....#..#..#....#..#..#....#..#.
..#.#...#O#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#..#.#...#.#
.#...##..#..X...##..#..#...##..#..#...##..#..#...##..#..#...##..#.
..#.##.......#.X#.......#.##.......#.##.......#.##.......#.##.....  --->
.#.#.#....#.#.#.#.O..#.#.#.#....#.#.#.#....#.#.#.#....#.#.#.#....#
.#........#.#........X.#........#.#........#.#........#.#........#
#.##...#...#.##...#...#.X#...#...#.##...#...#.##...#...#.##...#...
#...##....##...##....##...#X....##...##....##...##....##...##....#
.#..#...#.#.#..#...#.#.#..#...X.#.#..#...#.#.#..#...#.#.#..#...#.#  --->

In this example, traversing the map using this slope would cause you to encounter 7 trees.

Starting at the top-left corner of your map and following a slope of right 3 and down 1, how many trees would you encounter?

*/

const patternLength = 31
const xInc = 3
const yInc = 1
const treeSymbol = "#"

type treeLocations = [][]int
type coords struct {
	x, y int
}

type coordArr = []coords

var slopes coordArr = []coords{
	coords{1, 1},
	coords{3, 1},
	coords{5, 1},
	coords{7, 1},
	coords{1, 2},
}

func (c *coords) move(x, y int) {
	c.moveX(x)
	c.moveY(y)
}

func (c *coords) moveX(x int) {
	c.x = c.x + x
}

func (c *coords) moveY(y int) {
	c.y = c.y + y
}

func (c coords) current() {
	fmt.Printf("\n(%v, %v)", c.x, c.y)
}

func newCoords() coords {
	return coords{x: 0, y: 0}
}

func collision(xPos int, treeMap []int) bool {
	clampPosition := xPos % patternLength
	for _, pos := range treeMap {
		if clampPosition == pos {
			return true
		}
	}
	return false
}

func p202031(offset coords, trees treeLocations) int {
	count := 0
	p := newCoords()

	p.move(offset.x, offset.y)
	for p.y < len(trees) {
		if collision(p.x, trees[p.y]) {
			count++
		}
		p.move(offset.x, offset.y)
	}
	return count
}

func p202032(trees treeLocations) ([]int, int) {
	var counts []int
	for _, offset := range slopes {
		counts = append(counts, p202031(offset, trees))
	}

	total := 1
	for _, c := range counts {
		total *= c
	}
	return counts, total
}

func main() {
	utils.Banner(3)
	data := utils.LoadData("3.txt")
	trees := make([][]int, len(data))

	for i, treeMap := range data {
		t := []int{}
		for j, input := range treeMap {
			if string(input) == treeSymbol {
				t = append(t, j)
			}
		}
		trees[i] = t
	}

	fmt.Println("Collisions:", p202031(coords{3, 1}, trees))
	counts, product := p202032(trees)
	fmt.Println("Collisions:", counts, product)
}
