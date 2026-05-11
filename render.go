package main

func RenderLine(text string, banner map[rune][]string) []string {
	result := make([]string, 8)
	if text == "" {
		return result
	}
	for _, char := range text {
		for row := 0; row < 8; row++ {
			result[row] += banner[char][row]
		}
	}
	return result
}
