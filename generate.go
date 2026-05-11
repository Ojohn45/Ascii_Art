package main

import (
	"strings"
)

func GenerateArt(text string, banner map[rune][]string) string {
	var result strings.Builder
	word := strings.Split(text, "\\n")
	for i, words := range word {
		if words == "" {
			if i != len(word)-1 {
				result.WriteString("\n")
			}
			continue
		}
		for row := 0; row < 8; row++ {
			for _, char := range words {
				result.WriteString(banner[char][row])
				result.WriteString("\n")
			}
		}
	}
	return result.String()
}
