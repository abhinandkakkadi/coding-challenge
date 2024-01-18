package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	inputFileName := flag.String("c", "", "Provide the input file name")
	lineCount := flag.String("l", "", "Provide the input file name")
	wordCount := flag.String("w", "", "Provide the input file name")
	flag.Parse()

	var fileName *string
	if *inputFileName != "" {
		fileName = inputFileName
	} else if *lineCount != "" {
		fileName = lineCount
	} else if *wordCount != "" {
		fileName = wordCount
	}

	file, err := OpenFile(*fileName)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	defer file.Close()

	if *inputFileName != "" {
		
		fileInfo, err := file.Stat()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		fmt.Println(fileInfo.Size(), " ", *inputFileName)
	} else if *lineCount != "" {

		scanner := bufio.NewScanner(file)
		numLines := 0

		for scanner.Scan() {
			numLines++
		}

		if err = scanner.Err(); err != nil {
			fmt.Println("error scanning: "+err.Error())
			os.Exit(1)
		}

		fmt.Println(numLines, " ", *lineCount)
	} else if *wordCount != "" {

		scanner := bufio.NewScanner(file)
		numWords := 0

		for scanner.Scan() {
			wordsInLIne := strings.Fields(scanner.Text())
			numWords += len(wordsInLIne)
		}

		if err := scanner.Err();  err != nil {
			fmt.Println("error scanning: "+err.Error())
			os.Exit(1)
		}

		fmt.Println(numWords, " ", *wordCount)
	}

}

func OpenFile(inputFileName string) (*os.File, error) {
	file, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}
