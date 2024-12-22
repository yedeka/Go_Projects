package wctool

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yedeka/Go_Projects/filehandling"
)

func HandleByteCount(filePath string) int64 {
	absPath := FindFileAbsolutePath(filePath)
	return filehandling.ReadFileSize(absPath)
}

func HandleLineCount(filePath string) int {
	absPath := FindFileAbsolutePath(filePath)
	return filehandling.ReadFileLines(absPath)
}

func HandleWordCount(filePath string) int {
	absPath := FindFileAbsolutePath(filePath)
	return filehandling.ReadFileWords(absPath)
}

func HandleCharCount(filePath string) int {
	absPath := FindFileAbsolutePath(filePath)
	return filehandling.ReadFileChars(absPath)
}

func FindFileAbsolutePath(filePath string) string {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Println("Error getting absolute path for file:", filePath)
		os.Exit(3)
	}
	return absPath
}
