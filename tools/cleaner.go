package tools

import "strings"

// Text_cleaner cleans the input text by removing extra spaces, newlines, and adding a space at the beginning and end of each line.
func Text_cleaner(data string) string {
	// Remove extra spaces and join words with a single space.
	cleanedSpaces := SplitWhiteSpaces(data)
	splitedString := strings.Join(cleanedSpaces, " ")

	// Clean extra newlines and add a space to the beginning and end of each line.
	return cleannewlines(splitedString)
}

// SplitWhiteSpaces splits the string by spaces and tabs, removing extra spaces.
func SplitWhiteSpaces(s string) []string {
	str := ""
	table := []string{}
	isSpaced := false
	for i, v := range s {
		if v == ' ' && !isSpaced || v == '\t' && !isSpaced {
			table = append(table, str)
			str = ""
			isSpaced = true
		} else if v != ' ' && v != '\t' {
			str += string(v)
			isSpaced = false
		}
		if i == len(s)-1 {
			table = append(table, str)
		}
	}
	return table
}

// cleannewlines removes extra newlines from the string.
func cleannewlines(ss string) string {
	cleanString := ""
	isClean := false
	for i, v := range ss {
		if i != len(ss)-1 {
			if ss[i] == '\n' && ss[i+1] == '\n' {
				isClean = false
			} else {
				isClean = true
			}
			if isClean {
				cleanString += string(v)
			}
		}
	}
	return cleanString + string(ss[len(ss)-1])
}

// addSpaceToLines adds a space to the beginning of each line.
func addSpaceToLines(text string) string {
	lines := strings.Split(text, "\n")
	for i := range lines {
		lines[i] = " " + lines[i]
	}
	return strings.Join(lines, "\n")
}

// removeLeadingSpace removes the leading space from each line of the text.
func removeLeadingSpace(text string) string {
	lines := strings.Split(text, "\n")
	for i := range lines {
		lines[i] = strings.TrimLeft(lines[i], " ")
	}
	return strings.Join(lines, "\n")
}

// removeTrailingSpace removes the trailing space from each line of the text.
