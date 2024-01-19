package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	inputFileName := flag.String("c", "", "Provide the input file name")
	lineCount := flag.String("l", "", "Provide the input file name")
	wordCount := flag.String("w", "", "Provide the input file name")
	characterCount := flag.String("m", "", "Provide the input file name")
	flag.Parse()

	var fileName *string
	if *inputFileName != "" {
		fileName = inputFileName
	} else if *lineCount != "" {
		fileName = lineCount
	} else if *wordCount != "" {
		fileName = wordCount
	} else if *characterCount != "" {
		fileName = characterCount
	}

	allOption := false
	if *inputFileName == "" && *lineCount == "" && *wordCount == "" && *characterCount == "" && len(os.Args) == 2 {
		allOption = true
		fileName = &os.Args[1]
	}

	var file *os.File
	if *characterCount != "" {
		var err error
		file, err = OpenFile(*fileName)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}

		defer file.Close()
	}

	if allOption {
		
	}
	

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

		if err := scanner.Err(); err != nil {
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
	} else if *characterCount != "" {
		// 0644 is a permission which allows the user to manipulate the different type of user in different ways.
		file, err := os.OpenFile(*characterCount, os.O_RDONLY, 0644)
		if err != nil {
			err.Error()
			os.Exit(1)
		}
		defer file.Close()

		// It reads 1024 bytes at a time
		buffer := make([]byte, 1024)
		var numChar int

		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}

			numChar += n
		}

	fmt.Println(numChar, " ", *characterCount)
	}

}

func OpenFile(inputFileName string) (*os.File, error) {
	file, err := os.Open(inputFileName)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func NumBytes(file *os.File, inputFileName *string) {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(fileInfo.Size(), " ", *inputFileName)
}
