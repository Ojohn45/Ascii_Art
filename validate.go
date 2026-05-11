package main

func ValidateInput(input string) (rune, error) {
	for _, val := range input {
		for val < ' ' || val > '~' {

		}
	}
	return 0, nil
}