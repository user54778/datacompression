package utils

import (
	"bufio"
	"os"
	"unicode"
)

// CleanInput returns a rune slice containing only the unicode characters a-z or A-Z
// from some file.
func CleanInput(filePath string) ([]rune, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanRunes)

	var cleaned []rune

	for scanner.Scan() {
		r := []rune(scanner.Text())[0]
		if unicode.IsLetter(r) && (unicode.IsLower(r) || unicode.IsUpper(r)) {
			cleaned = append(cleaned, r)
		}
		if unicode.IsNumber(r) {
			cleaned = append(cleaned, r)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return cleaned, nil
}
