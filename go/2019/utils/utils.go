package utils

import (
	"bufio"
	"io/ioutil"
	"fmt"
	"os"
	"strconv"
	"strings"
	s "strings"
)

// Check checks for an error and panics if one is found
func CheckAndPanic(e error) {
	if e != nil {
		panic(e)
	}
}

// LoadData reads each line of input into a string array
func LoadData(filename string) []string {
	data := []string{}

	handler, err := os.Open(fmt.Sprintf("data/%s", filename))
	CheckAndPanic(err)

	scanner := bufio.NewScanner(handler)
	for scanner.Scan() {
		data = append(data, s.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return data
}

// LoadData reads each line of input into a string array
func LoadDataAsString(filename string) string {
	buf, err := ioutil.ReadFile(fmt.Sprintf("data/%s", filename))
	CheckAndPanic(err)

	return string(buf)
}

// Banner prints a text heading for the day specified
func Banner(day int) {
	fmt.Println("==============================")
	fmt.Printf(" Advent of code 2019 - Day %v\n", day)
	fmt.Println("==============================")
}

// StrToIntArr converts a delim separated string into a list of ints
func StrToIntArr(str string, rest ...string) []int {
	var delim string
	if len(rest) > 0 {
		delim = rest[0]
	} else {
		delim = ","
	}

	a := strings.Split(str, delim)
	b := make([]int, len(a))
	for i, v := range a {
		r, err := strconv.Atoi(v)
		CheckAndPanic(err)

		b[i] = r
	}
	return b
}
