package main

import (
	"fmt"
	"os"

	"go_reloaded/tools"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Invalid command. Please provide two filenames.")
		return
	}

	firstFilename := os.Args[1]
	secondFilename := os.Args[2]

	if !tools.Checktype(firstFilename) || !tools.Checktype(secondFilename) {
		fmt.Println("Invalid file type.")
		return
	}

	textToCorrect, err := os.ReadFile(firstFilename)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", firstFilename, err)
		return
	}
	if len(textToCorrect) == 0 {
		fmt.Println("sample.txt is empty ")
	} else if len(textToCorrect) != 0 {

		correctedString := tools.Text_corrector(string(textToCorrect))

		resultFile, err := os.Create("result.txt")
		if err != nil {
			fmt.Printf("Error creating result file: %v\n", err)
			return
		}
		defer resultFile.Close()

		data := []byte(correctedString)
		_, write_error := resultFile.Write(data)
		if write_error != nil {
			fmt.Printf("Error writing to result file: %v\n", err)
			return
		}

		expectedResult := getExpectedResult(firstFilename)
		fmt.Printf("Expected Text:\n%s\n", expectedResult)

		fmt.Println(`
╔════════════════════════════════════════╗
║   🎉 File processed successfully! 🎉   ║
╠════════════════════════════════════════╣
║   📂 Check 'result.txt' for output.    ║
╚════════════════════════════════════════╝
	`)
	}
}

// Helper function to return the expected result based on the sample file
func getExpectedResult(filename string) string {
	switch filename {
	case "sample1.txt":
		return "If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?"
	case "sample2.txt":
		return "I have to pack 5 outfits. Packed 26 just to be sure"
	case "sample3.txt":
		return "Don not be sad, because sad backwards is das. And das not good"
	case "sample4.txt":
		return "Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'"
	default:
		return ` ____
        / _\ \
      .'\/  \ \
    ,'   \   \ \
     / /-'    \ \ .
    / /       ,\ '|
   / /        '-._|
  / /_.'|________\_\
  \/_<  ___________/
      '.|`
	}
}
