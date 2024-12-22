package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/yedeka/Go_Projects/cmd/ccwc/wctool"
)

func count(fileContent []byte,
	byteCntFlag bool,
	lineCntFlag bool,
	wordCntFlag bool,
	charCntFlag bool) {

	if byteCntFlag {
		fmt.Printf("%d\t", len(fileContent))
	}
	if wordCntFlag {
		fmt.Printf("%d\t", len(bytes.Fields(fileContent)))
	}
	if charCntFlag {
		fmt.Printf("%d\t", len(bytes.Runes(fileContent)))
	}
	if lineCntFlag {
		lineCount := 0
		for i := 0; i < len(fileContent); i++ {
			if string(fileContent[i]) == "\n" {
				lineCount++
			}
		}
		fmt.Printf("%d\t", lineCount)
	}
}

func main() {
	var byteCntFlag, lineCntFlag, wordCntFlag, charCntFlag bool
	flag.BoolVar(&byteCntFlag, "c", false, "Flag to output number of bytes in a file")
	flag.BoolVar(&lineCntFlag, "l", false, "Flag to output number of lines in a file")
	flag.BoolVar(&wordCntFlag, "w", false, "Flag to output number of words in a file")
	flag.BoolVar(&charCntFlag, "m", false, "Flag to output number of characters in a file")
	flag.Parse()

	if !byteCntFlag && !charCntFlag && !lineCntFlag && !wordCntFlag {
		lineCntFlag = true
		wordCntFlag = true
		byteCntFlag = true
	}
	numargs := len(flag.CommandLine.Args())

	if numargs == 0 {
		// No arguments passed hence read from console
		var fileContent []byte
		fileContent, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println("Error while reading from console:\n", err)
			os.Exit(3)
		}
		count(fileContent, byteCntFlag, lineCntFlag, wordCntFlag, charCntFlag)

	} else {
		filename := flag.CommandLine.Arg(0)

		if byteCntFlag {
			fmt.Printf("%d\t", wctool.HandleByteCount(filename))
		}
		if lineCntFlag {
			fmt.Printf("%d\t", wctool.HandleLineCount(filename))
		}
		if wordCntFlag {
			fmt.Printf("%d\t", wctool.HandleWordCount(filename))
		}
		if charCntFlag {
			fmt.Printf("%d\t", wctool.HandleCharCount(filename))
		}
		fmt.Printf("%s\n", filename)
	}
}
