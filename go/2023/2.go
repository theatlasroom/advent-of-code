package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/theatlasroom/advent-of-code/go/utils"
)

/*
--- Day 2: Cube Conundrum ---

You're launched high into the atmosphere! The apex of your trajectory just barely reaches the surface of a large island floating in the sky. You gently land in a fluffy pile of leaves. It's quite cold, but you don't see much snow. An Elf runs over to greet you.

The Elf explains that you've arrived at Snow Island and apologizes for the lack of snow. He'll be happy to explain the situation, but it's a bit of a walk, so you have some time. They don't get many visitors up here; would you like to play a game in the meantime?

As you walk, the Elf shows you a small bag and some cubes which are either red, green, or blue. Each time you play this game, he will hide a secret number of cubes of each color in the bag, and your goal is to figure out information about the number of cubes.

To get information, once a bag has been loaded with cubes, the Elf will reach into the bag, grab a handful of random cubes, show them to you, and then put them back in the bag. He'll do this a few times per game.

You play several games and record the information from each game (your puzzle input). Each game is listed with its ID number (like the 11 in Game 11: ...) followed by a semicolon-separated list of subsets of cubes that were revealed from the bag (like 3 red, 5 green, 4 blue).

For example, the record of a few games might look like this:

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

In game 1, three sets of cubes are revealed from the bag (and then put back again). The first set is 3 blue cubes and 4 red cubes; the second set is 1 red cube, 2 green cubes, and 6 blue cubes; the third set is only 2 green cubes.

The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

In the example above, games 1, 2, and 5 would have been possible if the bag had been loaded with that configuration. However, game 3 would have been impossible because at one point the Elf showed you 20 red cubes at once; similarly, game 4 would also have been impossible because the Elf showed you 15 blue cubes at once. If you add up the IDs of the games that would have been possible, you get 8.

Determine which games would have been possible if the bag had been loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?

--- Part Two ---

The Elf says they've stopped producing snow because they aren't getting any water! He isn't sure why the water stopped; however, he can show you how to get to the water source to check it out for yourself. It's just up ahead!

As you continue your walk, the Elf poses a second question: in each game you played, what is the fewest number of cubes of each color that could have been in the bag to make the game possible?

Again consider the example games from earlier:

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green

    In game 1, the game could have been played with as few as 4 red, 2 green, and 6 blue cubes. If any color had even one fewer cube, the game would have been impossible.
    Game 2 could have been played with a minimum of 1 red, 3 green, and 4 blue cubes.
    Game 3 must have been played with at least 20 red, 13 green, and 6 blue cubes.
    Game 4 required at least 14 red, 3 green, and 15 blue cubes.
    Game 5 needed no fewer than 6 red, 3 green, and 2 blue cubes in the bag.

The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together. The power of the minimum set of cubes in game 1 is 48. In games 2-5 it was 12, 1560, 630, and 36, respectively. Adding up these five powers produces the sum 2286.

For each game, find the minimum set of cubes that must have been present. What is the sum of the power of these sets?

*/

const GAME_DELIMITER = ":"
const ROLL_DELIMITER = ";"
const DICE_DELIMITER = ","

var digitRe = regexp.MustCompile(`(\d+)`)

type Roll struct {
	Red, Green, Blue int
}

type GameConstraints struct {
	Red, Green, Blue int
}

type Game map[int][]Roll

func extractGameID(str string) int {
	idStr := digitRe.FindAllString(str, -1)

	id, err := strconv.Atoi(idStr[0])
	utils.CheckAndPanic(err)

	return id
}

func extractGameRoll(str string) Roll {
	r, g, b := 0, 0, 0
	for _, roll := range strings.Split(str, DICE_DELIMITER) {
		value, err := strconv.Atoi(digitRe.FindAllString(roll, -1)[0])
		if err != nil {
			value = 0
		}

		if strings.Contains(roll, "red") {
			r = value
			continue
		}

		if strings.Contains(roll, "green") {
			g = value
			continue
		}

		if strings.Contains(roll, "blue") {
			b = value
			continue
		}
	}
	return Roll{Red: r, Green: g, Blue: b}
}

func isPossible(cg GameConstraints, r Roll) bool {
	if r.Red > cg.Red || r.Blue > cg.Blue || r.Green > cg.Green {
		return false
	}
	return true
}

func (r Roll) power() int {
	base := 1

	if r.Red > 0 {
		base *= r.Red
	}

	if r.Green > 0 {
		base *= r.Green
	}

	if r.Blue > 0 {
		base *= r.Blue
	}

	return base
}

func findPossibleGames(cg GameConstraints, str string) int {
	var rolls []Roll

	res := strings.Split(str, GAME_DELIMITER)

	gameIDStr, rollsStr := res[0], res[1]
	gameID := extractGameID(gameIDStr)

	rollStrs := strings.Split(rollsStr, ROLL_DELIMITER)
	for _, roll := range rollStrs {
		rolls = append(rolls, extractGameRoll(roll))
	}

	for _, roll := range rolls {
		if !isPossible(cg, roll) {
			return 0
		}
	}

	return gameID
}

func findMinDicePerRoll(rolls []Roll) Roll {
	r, g, b := 0, 0, 0
	for _, roll := range rolls {
		if roll.Red > r {
			r = roll.Red
		}
		if roll.Green > g {
			g = roll.Green
		}
		if roll.Blue > b {
			b = roll.Blue
		}
	}

	return Roll{Red: r, Green: g, Blue: b}
}

func findMinimumCubes(str string) int {
	var rolls []Roll

	res := strings.Split(str, GAME_DELIMITER)
	rollStrs := strings.Split(res[1], ROLL_DELIMITER)

	for _, roll := range rollStrs {
		rolls = append(rolls, extractGameRoll(roll))
	}

	minDice := findMinDicePerRoll(rolls)

	return minDice.power()
}

func part1(input []string) {
	sum := 0

	c := GameConstraints{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	for _, line := range input {
		sum += findPossibleGames(c, line)
	}

	fmt.Println("Sum:", sum)
}

func part2(input []string) {
	sum := 0

	for _, line := range input {
		sum += findMinimumCubes(line)
	}

	fmt.Println("Sum:", sum)
}

func main() {
	utils.Banner(utils.BannerConfig{Year: 2023, Day: 2})
	// Read all the calibration values
	input := utils.LoadData("2.txt")

	part1(input)
	part2(input)
}
