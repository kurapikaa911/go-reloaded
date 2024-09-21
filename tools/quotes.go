package tools

import (
	"strings"
)

// correctSingleQuotesSpacing adjusts the spacing around text within single quotes.
// It removes spaces immediately inside the single quotes while preserving the rest of the text.
func correctSingleQuotesSpacing(input string) string {
	// Split the input string into parts using single quote as the delimiter.
	// This creates an array where even indices are outside quotes, and odd indices are inside quotes.
	parts := strings.Split(input, "'")

	// Iterate through the parts, focusing on the odd indices (inside quotes)
	for i := 1; i < len(parts); i += 2 {
		// Trim spaces from the beginning and end of each part inside quotes
		parts[i] = strings.TrimSpace(parts[i])
	}

	// Join all the parts back together, inserting single quotes between them
	// This reconstructs the string with correct spacing around quoted text
	return strings.Join(parts, "'")
}

// Single_Cote processes the content to adjust spacing around single quotes.
// func Single_Cote(content string) string {
// 	new_content := Split_Single_Cote(content)

// 	if len(new_content) == 0 {
// 		return content
// 	}

// 	str := ""
// 	for i, arg := range new_content {
// 		if arg[0] == '\'' && arg[len(arg)-1] == '\'' && arg != "h'" {
// 			if i < len(new_content)-1 && strings.ContainsAny(string(new_content[i+1][0]), ".,?!:;") {
// 				str += "'" + RemoveSpace(`arg[1:len(arg)-1]) + "'"
// 			} else {
// 				str += "'" + RemoveSpace(arg[1:len(arg)-1]) + "' "
// 			}
// 		} else {
// 			str += arg + " "
// 		}
// 	}

// 	return str
// }

// Split_Single_Cote splits the content based on single quotes.
// func Split_Single_Cote(Data string) []string {
// 	Words := []string{}
// 	firstindex := 0
// 	lastindex := 0
// 	inQuote := false // Indicator if we are inside quotes
// 	counter := 0

// 	for i, char := range Data {
// 		if char == '\'' {
// 			// Check if the quote is between two letters
// 			if i > 0 && i < len(Data)-1 && unicode.IsLetter(rune(Data[i-1])) && unicode.IsLetter(rune(Data[i+1])) {
// 				continue // Skip the quote as it's part of the word
// 			}

// 			if inQuote {
// 				// Closing the quote
// 				inQuote = false
// 				lastindex = i + 1
// 				Words = append(Words, Data[firstindex:lastindex]) // Add text between quotes
// 				counter++
// 			} else {
// 				// Opening the quote
// 				inQuote = true
// 				firstindex = i
// 				if lastindex < firstindex {
// 					Words = append(Words, Data[lastindex:firstindex]) // Add text outside quotes
// 				}
// 			}
// 		}
// 	}

// 	// If the quote is open and not closed
// 	if inQuote {
// 		Words = append(Words, Data[firstindex:])
// 	} else if len(Data) > lastindex {
// 		Words = append(Words, Data[lastindex:])
// 	}

// 	// println(counter) // Print the number of closed quotes
// 	return Words
// }

// RemoveSpace trims whitespace from the start and end of the string.
// func RemoveSpace(str string) string {
// 	return strings.TrimSpace(str)
// }
