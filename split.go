package main

import (
	"strings"
)

func SplitInput(input string) []string {
	words := strings.Split(input, "\\n")
	return words
}
