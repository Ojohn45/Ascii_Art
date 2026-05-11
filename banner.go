package main

import (
	"errors"
	"os"
	"strings"
)

func LoadBanner(filename string) (map[rune][]string, error) {
	fontMap := make(map[rune][]string)
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return fontMap, errors.New("invalid character")
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) < 856 {
		return fontMap, errors.New("invalid")
	}
	return fontMap 
}
