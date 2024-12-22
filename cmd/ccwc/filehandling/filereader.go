package filehandling

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func ReadFileSize(filepath string) int64 {
	file := readFile(filepath)
	// Get file Info
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Error getting file information:", err)
		os.Exit(3)
	}
	defer file.Close()
	return fileInfo.Size()
}

func ReadFileLines(filePath string) int {
	file := readFile(filePath)
	// creates a new Scanner for the file
	scanner := bufio.NewScanner(file)
	lineCount := 0
	// reads all lines until there is nothing to read
	for scanner.Scan() {
		lineCount++ // increment line count
	}
	// checks if there was any error during Scan
	if err := scanner.Err(); err != nil {
		fmt.Println("Error while counting lines of file")
		os.Exit(3)
	}
	defer file.Close()
	return lineCount
}

func ReadFileWords(filePath string) int {
	file := readFile(filePath)
	// creates a new Scanner for the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	// count the words
	wordcnt := 0
	for scanner.Scan() {
		wordcnt += 1
	}
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(3)
	}
	return wordcnt
}

func ReadFileChars(filePath string) int {
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		os.Exit(3)
	}
	chars := len(bytes.Runes(file))
	return chars
}

func readFile(filepath string) *os.File {
	// Open the file at given path
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error reading file: %v", err)
		os.Exit(3)
	}
	return file
}
