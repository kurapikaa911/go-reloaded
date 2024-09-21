package tools

import (
	"strconv"
	"strings"
)

// Preprocess text by removing spaces inside flags
func preprocessFlags(text string) string {
	// Define flag patterns to search for
	flags := []string{"(cap,", "(low,", "(up,"}

	for _, flag := range flags {
		// Replace instances like "(cap, 2)" with "(cap,2)"
		text = strings.ReplaceAll(text, flag+" ", flag)
	}

	return text
}

func handleFlags(text string) string {
	text = preprocessFlags(text)
	lines := strings.Split(text, "\n")
	var processedLines []string

	for _, line := range lines {
		words := SplitWhiteSpaces(line)
		j := 0
		for j < len(words) {
			if strings.HasPrefix(words[j], "(") && strings.HasSuffix(words[j], ")") {
				flagContent := strings.Trim(words[j], "()")
				parts := strings.Split(flagContent, ",")

				// Handle flags with numbers (e.g. "up,2")
				if len(parts) == 2 {
					flag := strings.TrimSpace(parts[0])
					numberString := strings.TrimSpace(parts[1])
					number, err := strconv.Atoi(numberString)
					if err == nil && number > 0 && j > 0 {
						switch flag {
						case "up", "low", "cap":
							for k := j - 1; k >= 0 && number > 0; k-- {
								switch flag {
								case "up":
									words[k] = toUpper(words[k])
								case "low":
									words[k] = toLower(words[k])
								case "cap":
									words[k] = toCapitalize(words[k])
								}
								number--
							}
						}
					}
					// Remove the processed flag with the number (e.g. "(up,2)")
					words = append(words[:j], words[j+1:]...)
				} else {
					// Handle single flags (e.g. "up", "low", "cap", "bin", "hex")
					flag := flagContent
					if isValidFlag(flag) && j > 0 {
						switch flag {
						case "bin", "hex":
							words[j-1] = convertbase(words[j-1], flag)
						case "up":
							words[j-1] = toUpper(words[j-1])
						case "low":
							words[j-1] = toLower(words[j-1])
						case "cap":
							words[j-1] = toCapitalize(words[j-1])
						}
						// Remove the processed flag (e.g. "(up)")
						words = append(words[:j], words[j+1:]...)
					} else {
						// Move to next word if not a recognized flag
						j++
					}
				}
			} else {
				j++
			}
		}
		processedLines = append(processedLines, strings.Join(words, " "))
	}

	return strings.Join(processedLines, "\n")
}

// Check if a string is a valid flag (e.g. "up", "low", "cap", "bin", "hex")
func isValidFlag(flag string) bool {
	validFlags := []string{"up", "low", "cap", "bin", "hex"}
	for _, valid := range validFlags {
		if flag == valid {
			return true
		}
	}
	return false
}
