package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	// banner := "standard.txt"
	// if len(os.Args) == 3 {
	// 	banner = os.Args[2] + ".txt"
	// }

	font := make(map[rune][]string)
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return font, errors.New("inavalid charater")
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) < 856 {
		return font, errors.New("invalid")
	}
	for ch := ' '; ch <= '~'; ch++ {
		start := (int(ch) - 32) * 9
		font[ch] = lines[start+1 : start+9]
	}
	return font, nil
}
