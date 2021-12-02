package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/theatlasroom/advent-of-code/go/utils"
)

/**
--- Day 7: Handy Haversacks ---

You land at the regional airport in time for your next flight. In fact, it looks like you'll even have time to grab some food:
all flights are currently delayed due to issues in luggage processing.

Due to recent aviation regulations, many rules (your puzzle input) are being enforced about bags and their contents;
bags must be color-coded and must contain specific quantities of other color-coded bags. Apparently, nobody responsible for these regulations considered how long they would take to enforce!

For example, consider the following rules:

light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.

These rules specify the required contents for 9 bag types. In this example, every faded blue bag is empty,
every vibrant plum bag contains 11 bags (5 faded blue and 6 dotted black), and so on.

You have a shiny gold bag. If you wanted to carry it in at least one other bag, how many different bag colors would be
valid for the outermost bag? (In other words: how many colors can, eventually, contain at least one shiny gold bag?)

In the above rules, the following options would be available to you:

    A bright white bag, which can hold your shiny gold bag directly.
    A muted yellow bag, which can hold your shiny gold bag directly, plus some other bags.
    A dark orange bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.
    A light red bag, which can hold bright white and muted yellow bags, either of which could then hold your shiny gold bag.

So, in this example, the number of bag colors that can eventually contain at least one shiny gold bag is 4.

How many bag colors can eventually contain at least one shiny gold bag? (The list of rules is quite long; make sure you get all of it.)

*/

type bagID struct {
	col, shade string
}

type bagContent struct {
	id    string
	count int
}

type bagContents []bagContent

type bag struct {
	id       string
	contents bagContents // ids of bags that can be contained in this bag
}

func (bc bagContents) contains(id string) bool {
	for _, item := range bc {
		if id == item.id {
			return true
		}
	}
	return false
}

func (bc bagContents) toArray() []string {
	var arr []string
	for _, item := range bc {
		arr = append(arr, item.id)
	}
	return arr
}

type bags map[string]bag

func generateBagID(b bagID) string {
	return strings.Join([]string{b.shade, b.col}, "_")
}

func extractContents(rawContents []string) bagContents {
	var contents bagContents
	for _, content := range rawContents {
		csplit := strings.Split(strings.Trim(content, " "), " ")
		count, err := strconv.Atoi(csplit[0])
		if err != nil {
			count = 0
		}
		contents = append(contents, bagContent{id: generateBagID(bagID{col: csplit[2], shade: csplit[1]}), count: count})
	}
	return contents
}

func parseBag(str string) (string, bag) {
	bagAndContents := strings.Split(str, "contain")
	b, c := bagAndContents[0], bagAndContents[1]

	bagAttributes := strings.Split(b, " ")
	shade, col := bagAttributes[0], bagAttributes[1]

	var contents bagContents
	if !strings.Contains(c, "no other") {
		contents = extractContents(strings.Split(c, ","))
	}
	id := generateBagID(bagID{col: col, shade: shade})
	return id, bag{id: id, contents: contents}
}

func extractBags(data []string) bags {
	b := make(bags)
	for _, str := range data {
		key, nextBag := parseBag(str)
		b[key] = nextBag
	}
	return b
}

func contains(s []string, str string) bool {
	for _, v := range s {
		fmt.Println("v", v, "str", str)
		if v == str {
			return true
		}
	}
	return false
}

func mergeArraysWithoutDuplicates(left, right []string) []string {
	for _, el := range left {
		if !contains(right, el) {
			right = append(left, el)
		}
	}
	return right
}

// func canContainTargetBag(allBags bags, candidates []string, targetBagID string) bool {
// 	// for each bag, check all its children
// 	// fmt.Println(currentBag.id, targetBagID, currentBag.contents)
// 	// if contains(candidates, targetBagID) {
// 	// 	return true
// 	// }

// 	// travese the contents of this bag
// 	for _, nextBagID := range candidates {
// 		fmt.Println("candidates", candidates)
// 		if contains(candidates, targetBagID) {
// 			return true
// 		}
// 		nextBag := allBags[nextBagID]
// 		bagContents := nextBag.contents
// 		if len(bagContents) < 1 {
// 			fmt.Println(nextBag.id, " is empty")
// 			continue
// 		}
// 		fmt.Println(nextBagID, "=>", bagContents.toArray())
// 		fmt.Println("candidates", candidates)
// 		return canContainTargetBag(allBags, bagContents.toArray(), targetBagID)
// 		// candidates = mergeArraysWithoutDuplicates(candidates, bagContents.toArray())
// 	}

// 	return false
// }

func canContainTargetBag(allBags bags, candidates []string, targetBagID string) bool {
	for _, nextBagID := range candidates {
		if contains(candidates, targetBagID)
	}
	return false
}

func p20200701(bs bags, targetBag string) int {
	count := 0
	for _, b := range bs {
		fmt.Println("")
		fmt.Println(b.id, " => ", b.contents)
		if len(b.contents) > 0 {
			ok := canContainTargetBag(bs, b.contents.toArray(), targetBag)
			fmt.Println("can contain: ", ok)
			if ok {
				count++
			}
		}
	}
	return count
}

// TODO: should redo with proper trees
func main() {
	utils.Banner(utils.BannerConfig{ Year: 2020, Day: 7 })
	data := utils.LoadData("7.txt")
	bs := extractBags(data)
	targetID := generateBagID(bagID{col: "gold", shade: "shiny"})
	// fmt.Println(len(bs))
	fmt.Println(p20200701(bs, targetID))
}
