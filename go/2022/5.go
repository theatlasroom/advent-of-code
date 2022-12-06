package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/theatlasroom/advent-of-code/go/utils"
)

/*
*
--- Day 5: Supply Stacks ---

The expedition can depart as soon as the final supplies have been unloaded from the ships. Supplies are stored in stacks of marked crates, but because the needed supplies are buried under many other crates, the crates need to be rearranged.

The ship has a giant cargo crane capable of moving crates between stacks. To ensure none of the crates get crushed or fall over, the crane operator will rearrange them in a series of carefully-planned steps. After the crates are rearranged, the desired crates will be at the top of each stack.

The Elves don't want to interrupt the crane operator during this delicate procedure, but they forgot to ask her which crate will end up where, and they want to be ready to unload them as soon as possible so they can embark.

They do, however, have a drawing of the starting stacks of crates and the rearrangement procedure (your puzzle input). For example:

	[D]

[N] [C]
[Z] [M] [P]

	1   2   3

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2

In this example, there are three stacks of crates. Stack 1 contains two crates: crate Z is on the bottom, and crate N is on top. Stack 2 contains three crates; from bottom to top, they are crates M, C, and D. Finally, stack 3 contains a single crate, P.

Then, the rearrangement procedure is given. In each step of the procedure, a quantity of crates is moved from one stack to a different stack. In the first step of the above rearrangement procedure, one crate is moved from stack 2 to stack 1, resulting in this configuration:

[D]
[N] [C]
[Z] [M] [P]

	1   2   3

In the second step, three crates are moved from stack 1 to stack 3. Crates are moved one at a time, so the first crate to be moved (D) ends up below the second and third crates:

	       [Z]
	       [N]
	   [C] [D]
	   [M] [P]
	1   2   3

Then, both crates are moved from stack 2 to stack 1. Again, because crates are moved one at a time, crate C ends up below crate M:

	[Z]
	[N]

[M]     [D]
[C]     [P]

	1   2   3

Finally, one crate is moved from stack 1 to stack 2:

	[Z]
	[N]
	[D]

[C] [M] [P]

	1   2   3

The Elves just need to know which crate will end up on top of each stack; in this example, the top crates are C in stack 1, M in stack 2, and Z in stack 3, so you should combine these together and give the Elves the message CMZ.

After the rearrangement procedure completes, what crate ends up on top of each stack?

--- Part Two ---

As you watch the crane operator expertly rearrange the crates, you notice the process isn't following your prediction.

Some mud was covering the writing on the side of the crane, and you quickly wipe it away. The crane isn't a CrateMover 9000 - it's a CrateMover 9001.

The CrateMover 9001 is notable for many new and exciting features: air conditioning, leather seats, an extra cup holder, and the ability to pick up and move multiple crates at once.

Again considering the example above, the crates begin in the same configuration:

    [D]
[N] [C]
[Z] [M] [P]
 1   2   3

Moving a single crate from stack 2 to stack 1 behaves the same as before:

[D]
[N] [C]
[Z] [M] [P]
 1   2   3

However, the action of moving three crates from stack 1 to stack 3 means that those three moved crates stay in the same order, resulting in this new configuration:

        [D]
        [N]
    [C] [Z]
    [M] [P]
 1   2   3

Next, as both crates are moved from stack 2 to stack 1, they retain their order as well:

        [D]
        [N]
[C]     [Z]
[M]     [P]
 1   2   3

Finally, a single crate is still moved from stack 1 to stack 2, but now it's crate C that gets moved:

        [D]
        [N]
        [Z]
[M] [C] [P]
 1   2   3

In this example, the CrateMover 9001 has put the crates in a totally different order: MCD.

Before the rearrangement process finishes, update your simulation so that the Elves know where they should stand to be ready to unload the final supplies. After the rearrangement procedure completes, what crate ends up on top of each stack?

*
*/

type stack []string
type stacks []stack
type move struct{ Count, From, To int }
type moves []move

var moveRegex = regexp.MustCompile(`\w+\s(\d{1,2})`)
var endConfigurationRegex = regexp.MustCompile(`([\s][\d][\s])+`)

func newMove(s [][]string) move {
	count, countOk := strconv.Atoi(s[0][1])
	if countOk != nil {
		count = 0
	}

	from, fromOk := strconv.Atoi(s[1][1])
	if fromOk != nil {
		from = 0
	}

	to, toOk := strconv.Atoi(s[2][1])
	if toOk != nil {
		to = 0
	}

	return move{Count: count, From: from, To: to}
}

func reverse(s stack) stack {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func (s stack) push(v stack) stack {
	return append(s, v...)
}

func (s stack) pop(count int) (newSlice stack, result stack) {
	index := len(s) - count
	return s[index:], s[:index]
}

func parseMoves(data []string) moves {
	var m moves
	for _, line := range data {
		values := moveRegex.FindAllStringSubmatch(line, -1)
		if len(values) > 0 {
			m = append(m, newMove(values))
		}
	}
	return m
}

func sortCargo(s stacks, m moves, preserveOrder bool) stacks {
	for _, nextMove := range m {
		to := nextMove.To - 1
		from := nextMove.From - 1

		popped, resulting := s[from].pop(nextMove.Count)
		if !preserveOrder {
			popped = reverse(popped)
		}
		pushed := s[to].push(popped)

		s[to], s[from] = pushed, resulting
	}
	return s
}

func topItems(s stacks) string {
	str := ""

	for _, cargo := range s {
		str += cargo[len(cargo)-1]
	}
	return str
}

func part2(s stacks, m moves) {
	res := sortCargo(s, m, true)
	top := topItems(res)
	fmt.Println("Part 2: top stack", top)
}

func part1(s stacks, m moves) {
	res := sortCargo(s, m, false)
	top := topItems(res)
	fmt.Println("Part 1: top stack", top)
}

func prepend(s stack, char string) stack {
	return append([]string{char}, s...)
}

func parseConfiguration(lines []string) stacks {
	res := stacks{}

	startIndex := 1
	offset := 4

	// initialize the stacks
	for i := 0; i < len(lines[0]); i += offset {
		res = append(res, stack{})
	}

	for _, line := range lines {
		if endConfigurationRegex.MatchString(line) {
			break
		}

		for i := startIndex; i < len(line); i += offset {
			stackIndex := i / offset

			char := string(line[i])
			if char != " " {
				res[stackIndex] = prepend(res[stackIndex], char)
			}
		}
	}

	return res
}

func main() {
	// Read all the numbers
	data := utils.LoadDataWithSpacesPreserved("5.txt")
	cfg := utils.BannerConfig{Year: 2022, Day: 5}
	utils.Banner(cfg)

	config := parseConfiguration(data)
	m := parseMoves(data)

	part1(config, m)
	// part2(config, m)
}
