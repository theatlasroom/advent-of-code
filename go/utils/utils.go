package utils

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	s "strings"
)

type BannerConfig struct {
	Year, Day int
}

// CheckAndPanic checks for an error and panics if one is found
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

// LoadDataAsString reads each line of input into a single string
func LoadDataAsString(filename string) string {
	buf, err := ioutil.ReadFile(fmt.Sprintf("data/%s", filename))
	CheckAndPanic(err)

	return string(buf)
}

// Banner prints a text heading for the day specified
func Banner(cfg BannerConfig) {
	fmt.Println("==============================")
	fmt.Printf(" Advent of code %d - Day %d\n", cfg.Year, cfg.Day)
	fmt.Println("==============================")
}

const defaultDelim = "\n"

// StrToIntArr converts a delim separated string into a list of ints
func StrToIntArr(str string, rest ...string) []int {
	var delim string
	if len(rest) > 0 {
		delim = rest[0]
	} else {
		delim = defaultDelim
	}

	a := s.Split(str, delim)
	var b []int

	for _, v := range a {
		if len(v) < 1 {
			continue
		}
		r, err := strconv.Atoi(v)
		CheckAndPanic(err)

		b = append(b, r)
	}
	return b
}

func CountBlankLines(data string) {
	countLines := 0
	lines := s.Split(data, "\n")
	for _, line := range lines {
		if line == "" {
			countLines++
		}
	}
	fmt.Println("Total: " + strconv.Itoa(countLines))
}
