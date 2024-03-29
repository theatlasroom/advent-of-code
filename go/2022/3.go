package main

import (
	"fmt"
	"sort"

	"github.com/theatlasroom/advent-of-code/go/utils"
)

/**
--- Day 3: Rucksack Reorganization ---

One Elf has the important job of loading all of the rucksacks with supplies for the jungle journey. Unfortunately, that Elf didn't quite follow the packing instructions, and so a few items now need to be rearranged.

Each rucksack has two large compartments. All items of a given type are meant to go into exactly one of the two compartments. The Elf that did the packing failed to follow this rule for exactly one item type per rucksack.

The Elves have made a list of all of the items currently in each rucksack (your puzzle input), but they need your help finding the errors. Every item type is identified by a single lowercase or uppercase letter (that is, a and A refer to different types of items).

The list of items for each rucksack is given as characters all on a single line. A given rucksack always has the same number of items in each of its two compartments, so the first half of the characters represent items in the first compartment, while the second half of the characters represent items in the second compartment.

For example, suppose you have the following list of contents from six rucksacks:

vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw

    The first rucksack contains the items vJrwpWtwJgWrhcsFMMfFFhFp, which means its first compartment contains the items vJrwpWtwJgWr, while the second compartment contains the items hcsFMMfFFhFp. The only item type that appears in both compartments is lowercase p.
    The second rucksack's compartments contain jqHRNqRjqzjGDLGL and rsFMfFZSrLrFZsSL. The only item type that appears in both compartments is uppercase L.
    The third rucksack's compartments contain PmmdzqPrV and vPwwTWBwg; the only common item type is uppercase P.
    The fourth rucksack's compartments only share item type v.
    The fifth rucksack's compartments only share item type t.
    The sixth rucksack's compartments only share item type s.

To help prioritize item rearrangement, every item type can be converted to a priority:

    Lowercase item types a through z have priorities 1 through 26.
    Uppercase item types A through Z have priorities 27 through 52.

In the above example, the priority of the item type that appears in both compartments of each rucksack is 16 (p), 38 (L), 42 (P), 22 (v), 20 (t), and 19 (s); the sum of these is 157.

Find the item type that appears in both compartments of each rucksack. What is the sum of the priorities of those item types?
**/

type boolset map[string]bool

func scoreItem(str string) int {
	ascii := int([]rune(str)[0])

	ucaseindex, ucaseoffset := 64, 26
	lcaseindex := 96
	if ascii > lcaseindex {
		return ascii - lcaseindex
	}
	return ascii - ucaseindex + ucaseoffset
}

func findCommonItem(midpoint int, left, right string) string {
	lmap, rmap := boolset{}, boolset{}
	var lc, rc, dup string

	for i := 0; i < midpoint; i++ {
		lc, rc = string(left[i]), string(right[i])

		if lc == rc {
			dup = lc
			break
		}

		lmap[lc], rmap[rc] = true, true
	}

	if dup == "" {
		for key := range lmap {
			if rmap[key] {
				dup = key
				break
			}
		}
	}

	return dup
}

func findLargestString(strs []string) []string {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	return strs
}

func compileBoolmaps(rucksacks ...string) []boolset {
	boolmap := make([]boolset, len(rucksacks))

	for index, sack := range rucksacks {
		boolmap[index] = boolset{}

		for _, s := range sack {
			boolmap[index][string(s)] = true
		}
	}

	return boolmap
}

func part2(data []string) {
	sum := 0

	for i := 0; i < len(data); i += 3 {
		one, two, three := data[i:i+1], data[i+1:i+2], data[i+2:i+3]
		strmap := compileBoolmaps(one[0], two[0], three[0])
		str := findLargestString([]string{one[0], two[0], three[0]})[2]

		common := ""
		for _, s := range str {
			found := true

			for _, m := range strmap {
				if !m[string(s)] {
					found = false
				}
			}

			if found {
				common = string(s)
				break
			}
		}

		sum += scoreItem(common)
	}

	fmt.Printf("Part 2: Sum %d\n", sum)
}

func part1(data []string) {
	sum := 0

	var midpoint int
	var item string

	for _, rucksack := range data {
		midpoint = len(rucksack) / 2
		item = findCommonItem(midpoint, rucksack[:midpoint], rucksack[midpoint:])
		sum += scoreItem(item)
	}

	fmt.Printf("Part 1: Sum %d\n", sum)
}

func main() {
	// Read all the numbers
	data := utils.LoadData("3.txt")
	cfg := utils.BannerConfig{Year: 2022, Day: 3}
	utils.Banner(cfg)
	part1(data)
	part2(data)
}
