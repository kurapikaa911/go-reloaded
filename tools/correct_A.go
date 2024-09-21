package tools

import "strings"

// correctA replaces "a" with "an" and "A" with "An" if the following word starts with a vowel or 'h'
func correctA(text string) string {
	// Split the text into words
	words := strings.Split(text, " ")

	for i := 0; i < len(words)-1; i++ {
		currentWord := words[i]
		nextWord := words[i+1]

		// Check if the current word is "a" or "A"
		if strings.ToLower(currentWord) == "a" {
			// Check if the next word starts with a vowel or 'h'
			if len(nextWord) > 0 && strings.Contains("aeiouh", strings.ToLower(string(nextWord[0]))) {
				// Replace "a" with "an" or "A" with "An" while preserving original case
				if currentWord == "A" {
					words[i] = "An"
				} else {
					words[i] = "an"
				}
			}
		}
	}

	return strings.Join(words, " ")
}
