package tools

// Text_corrector processes the input text through several stages to clean and correct it.
func Text_corrector(txt string) string {
	// Step 1: Clean the text.
	// The Text_cleaner function removes unwanted characters and whitespace.
	cleanedText := Text_cleaner(txt)

	// Step 2: Handle flags.
	// The handleFlags function processes the cleaned text based on specific flags or formatting rules.
	flagHandledText := handleFlags(addSpaceToLines(cleanedText))

	// Step 3: Correct punctuation spacing.
	// The correctPunctuationSpacing function adjusts punctuation and fixes spacing issues.
	punctuationCorrectedText := correctPunctuationSpacing(flagHandledText)

	// Step 4: Correct 'a' and 'an' usage.
	// The correctA function ensures the correct usage of these articles based on the context.
	correctedAtext := correctA(punctuationCorrectedText)

	// Step 5: Fix spacing around single quotes.
	// The correctSingleQuotesSpacing function handles spacing corrections related to single quotes.
	finalText := removeLeadingSpace(correctSingleQuotesSpacing(correctedAtext))

	// Return the fully corrected text.
	return finalText
}
