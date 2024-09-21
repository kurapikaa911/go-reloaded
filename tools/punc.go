package tools

// correctPunctuationSpacing adjusts spacing around punctuation marks.
// It ensures there's no space before a punctuation mark and ensures exactly
// one space after a punctuation mark unless it's followed by another punctuation mark.
func correctPunctuationSpacing(text string) string {
	runes := []rune(text) // Convert text to a slice of runes to handle multi-byte characterss 
	var result []rune     // Create an empty slice to store the corrected text.

	for i := 0; i < len(runes); i++ { // Loop through each rune in the input text.
		currentChar := runes[i] // Get the current character (rune).

		// Check if the current character is a punctuation mark.
		if currentChar == '.' || currentChar == ',' || currentChar == '!' || currentChar == '?' || currentChar == ':' || currentChar == ';' {
			// If the previous character is a space, remove it (no space before punctuation).
			if len(result) > 0 && result[len(result)-1] == ' ' {
				result = result[:len(result)-1] // Remove the last character (a space) from result.
			}
			// Append the punctuation mark to the result.
			result = append(result, currentChar)

			// Check if there's no space after the punctuation and the next character is not another punctuation mark.
			if i+1 < len(runes) && runes[i+1] != ' ' && (runes[i+1] != '.' && runes[i+1] != ',' && runes[i+1] != '!' && runes[i+1] != '?' && runes[i+1] != ':' && runes[i+1] != ';') {
				// If no space follows the punctuation, insert one.
				result = append(result, ' ')
			}
		} else {
			// If the current character is not a punctuation mark, append it to the result.
			result = append(result, currentChar)
		}
	}

	// Convert the result slice of runes back to a string and return it.
	return string(result)
}
