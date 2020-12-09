package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/theatlasroom/advent-of-code/go/2020/utils"
)

/**
--- Day 2: Password Philosophy ---

Your flight departs in a few days from the coastal airport; the easiest way down to the coast from here is via toboggan.

The shopkeeper at the North Pole Toboggan Rental Shop is having a bad day.
"Something's wrong with our computers; we can't log in!" You ask if you can take a look.

Their password database seems to be a little corrupted: some of the passwords wouldn't have been allowed by the Official Toboggan Corporate Policy
that was in effect when they were chosen.

To try to debug the problem, they have created a list (your puzzle input) of passwords (according to the corrupted database)
and the corporate policy when that password was set.

For example, suppose you have the following list:

1-3 a: abcde
1-3 b: cdefg
2-9 c: ccccccccc

Each line gives the password policy and then the password. The password policy indicates the lowest and highest number of times a given letter must appear for the password to be valid. For example, 1-3 a means that the password must contain a at least 1 time and at most 3 times.

In the above example, 2 passwords are valid. The middle password, cdefg, is not; it contains no instances of b, but needs at least 1. The first and third passwords are valid: they contain one a or nine c, both within the limits of their respective policies.

How many passwords are valid according to their policies?
**/

type rule struct {
	min, max  int
	character string
}

type policy struct {
	rule
	password string
}

var re = regexp.MustCompile(`^(\d+)-(\d+)\s([a-z]):\s(.*)$`)

func newPolicy(input string) policy {
	str := re.FindStringSubmatch(input)
	minStr, maxStr, character, password := str[1], str[2], str[3], str[4]

	min, err := strconv.Atoi(minStr)
	utils.CheckAndPanic(err)

	max, err := strconv.Atoi(maxStr)
	utils.CheckAndPanic(err)

	return policy{
		rule{
			min,
			max,
			character,
		},
		password,
	}
}

func extractPolicies(data []string) []policy {
	var policies []policy
	for _, d := range data {
		policies = append(policies, newPolicy(d))
	}
	return policies
}

type validatorFn = func(policy) bool

func isValidRentalRule(p policy) bool {
	tally := 0
	for _, c := range p.password {
		if string(c) == p.character {
			tally++
		}
	}
	if p.min <= tally && tally <= p.max {
		return true
	}
	return false
}

func isValidTobogganRule(p policy) bool {
	x := string(p.password[p.min-1]) == p.character
	y := string(p.password[p.max-1]) == p.character
	if (x || y) && !(x && y) {
		return true
	}
	return false
}

func findValidPasswords(policies []policy, validator validatorFn) int {
	tally := 0
	for _, p := range policies {
		if validator(p) {
			tally++
		}
	}
	return tally
}

func main() {
	utils.Banner(2)
	data := utils.LoadData("2.txt")

	policies := extractPolicies(data)
	fmt.Printf("%d valid rental passwords\n", findValidPasswords(policies, isValidRentalRule))
	fmt.Printf("%d valid toboggan passwords\n", findValidPasswords(policies, isValidTobogganRule))
}
