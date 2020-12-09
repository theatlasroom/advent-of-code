package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/theatlasroom/advent-of-code/go/2020/utils"
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

type bagContent struct {
	id    string
	count int
}

type bag struct {
	id       string
	contents []bagContent // ids of bags that can be contained in this bag
}

type bags []bag

func bagID(col, shade string) string {
	return strings.Join([]string{shade, col}, "_")
}

func extractContents(rawContents []string) []bagContent {
	var contents []bagContent
	for _, content := range rawContents {
		fmt.Println(content)
		csplit := strings.Split(strings.Trim(content, " "), " ")
		count, err := strconv.Atoi(csplit[0])
		fmt.Println(csplit)
		if err != nil {
			count = 0
		}
		contents = append(contents, bagContent{id: bagID(csplit[2], csplit[1]), count: count})
	}
	return contents
}

func parseBag(str string) bag {
	bagAndContents := strings.Split(str, "contain")
	b, c := bagAndContents[0], bagAndContents[1]

	bagAttributes := strings.Split(b, " ")
	shade, col := bagAttributes[0], bagAttributes[1]

	var contents []bagContent
	if !strings.Contains(c, "no other") {
		contents = extractContents(strings.Split(c, ","))
	}
	return bag{id: bagID(col, shade), contents: contents}
}

func extractBags(data []string) bags {
	var b bags
	for _, str := range data {
		b = append(b, parseBag(str))
	}
	return b
}

// func canContainBag(bagID string) int {
// }

// TODO: should redo with trees
func main() {
	utils.Banner(7)
	data := utils.LoadData("7.txt")
	bags := extractBags(data)
	fmt.Println(data)
	for _, bag := range bags {
		fmt.Println(bag)
	}
}
