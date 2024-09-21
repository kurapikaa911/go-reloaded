package tools

import (
	"strings"
	"unicode"
)

func toUpper(word string) string {
	return strings.ToUpper(word)
}

// Convert word to lowercase
func toLower(word string) string {
	return strings.ToLower(word)
}

// Convert word to capitalized (first letter uppercase, rest lowercase)
func toCapitalize(word string) string {
	if len(word) == 0 {
		return word
	}

	runes := []rune(word)
	for i, r := range runes {
		if unicode.IsLetter(r) {
			runes[i] = unicode.ToUpper(r)
			// Convert the rest of the runes to lowercase
			for j := i + 1; j < len(runes); j++ {
				runes[j] = unicode.ToLower(runes[j])
			}
			break
		}
	}
	return string(runes)
}
