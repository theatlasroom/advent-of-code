package utils

import (
	"bufio"
	"os"
	"fmt"
	s "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func LoadData(filename string) []string {
	data := []string{}

	handler, err := os.Open(fmt.Sprintf("data/%s", filename))
	check(err)

	scanner := bufio.NewScanner(handler)
	for scanner.Scan() {
		data = append(data, s.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return data
}