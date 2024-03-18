package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	countBytes := flag.Bool("c", false, "-c Option outputs the number of bytes in the file")
	countLines := flag.Bool("l", false, "-l option outputs the number of lines in the file")
	countWords := flag.Bool("w", false, "-w option outputs the number of words in the file")
	countChars := flag.Bool("m", false, "-m option outputs the number of characters in the file")

	flag.Parse()
	fileName := flag.Args()

	//No flags provided.  set all to true
	if !*countBytes && !*countLines && !*countWords {
		*countBytes = true
		*countLines = true
		*countWords = true
	}

	var rd *bufio.Reader

	if len(fileName) > 0 {
		file, err := os.Open(fileName[0])
		if err != nil {
			fmt.Println("Error Reading File", fileName[0], err)
			return
		}
		rd = bufio.NewReader(file)
	} else {
		rd = bufio.NewReader(os.Stdin)
	}

	byteCount := 0
	lineCount := 0
	charCount := 0
	wordCount := 0

	newLineRune, _ := utf8.DecodeRuneInString("\n")

	r, s, err := rd.ReadRune()
	prevRune := utf8.RuneError

	for err == nil {
		charCount++
		byteCount += s

		if r == newLineRune {
			lineCount++
		}

		if unicode.IsSpace(r) && !unicode.IsSpace(prevRune) {
			wordCount++
		}

		prevRune = r
		r, s, err = rd.ReadRune()
	}

	printArgs := make([]any, 0, 5)

	if *countLines {
		printArgs = append(printArgs, fmt.Sprintf("%d ", lineCount))
	}

	if *countWords {
		printArgs = append(printArgs, fmt.Sprintf("%d ", wordCount))
	}

	if *countBytes {
		printArgs = append(printArgs, fmt.Sprintf("%d ", byteCount))
	}

	if *countChars {
		printArgs = append(printArgs, fmt.Sprintf("%d ", charCount))
	}

	if len(fileName) > 0 {
		printArgs = append(printArgs, fileName[0])
	}

	fmt.Println(printArgs...)
}
