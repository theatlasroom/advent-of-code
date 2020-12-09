package main

import (
	"fmt"
	"strings"

	"github.com/theatlasroom/advent-of-code/go/2020/utils"
)

/**
--- Day 6: Custom Customs ---

As your flight approaches the regional airport where you'll switch to a much larger plane,
customs declaration forms are distributed to the passengers.

The form asks a series of 26 yes-or-no questions marked a through z. All you need to do is identify the questions
for which anyone in your group answers "yes". Since your group is just you, this doesn't take very long.

However, the person sitting next to you seems to be experiencing a language barrier and asks if you can help.
For each of the people in their group, you write down the questions for which they answer "yes", one per line. For example:

abcx
abcy
abcz

In this group, there are 6 questions to which anyone answered "yes": a, b, c, x, y, and z.
(Duplicate answers to the same question don't count extra; each question counts at most once.)

Another group asks for your help, then another, and eventually you've collected answers from every group on the plane (your puzzle input).
Each group's answers are separated by a blank line, and within each group, each person's answers are on a single line. For example:

abc

a
b
c

ab
ac

a
a
a
a

b

This list represents answers from five groups:

    The first group contains one person who answered "yes" to 3 questions: a, b, and c.
    The second group contains three people; combined, they answered "yes" to 3 questions: a, b, and c.
    The third group contains two people; combined, they answered "yes" to 3 questions: a, b, and c.
    The fourth group contains four people; combined, they answered "yes" to only 1 question, a.
    The last group contains one person who answered "yes" to only 1 question, b.

In this example, the sum of these counts is 3 + 3 + 3 + 1 + 1 = 11.

For each group, count the number of questions to which anyone answered "yes". What is the sum of those counts?
*/

type answers map[string]bool

type group struct {
	responses string
	answers
	completeAnswers int
}

type customsGroups []group

func (g group) countAnswers() int {
	return len(g.answers)
}

func (cg customsGroups) sumAnswers() int {
	count := 0
	for _, g := range cg {
		count += g.countAnswers()
	}
	return count
}

func (cg customsGroups) sumCompleteAnswers() int {
	count := 0
	for _, g := range cg {
		count += g.completeAnswers
	}
	return count
}

func parseGroupResponses(str string) (answers, int) {
	participants := len(strings.Split(str, " "))
	completed := 0

	a := make(answers)
	unique := make(map[string]int)

	for _, x := range str {
		char := string(x)
		if char != " " {
			a[char] = true
			unique[char]++
		}
	}

	for _, y := range unique {
		if y >= participants {
			completed++
		}
	}

	return a, completed
}

func parseData(data string) customsGroups {
	var cg customsGroups

	groupAnswers := strings.Split(strings.Replace(data, "\n", " ", -1), "  ")
	for _, responses := range groupAnswers {

		answers, completeAnswers := parseGroupResponses(responses)
		cg = append(cg, group{responses, answers, completeAnswers})
	}

	return cg
}

func main() {
	utils.Banner(6)
	data := utils.LoadDataAsString("6.txt")
	groups := parseData(data)

	fmt.Println(groups.sumAnswers())
	fmt.Println(groups.sumCompleteAnswers())
}
