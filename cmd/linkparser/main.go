package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yedeka/Go_Projects/cmd/linkparser/link"
)

func main() {
	fileName := flag.String("htmlFilePath", "text1.html", "Path of the HTML file to parse for links")
	flag.Parse()
	fmt.Println(*fileName)
	file, err := os.Open(*fileName)
	if nil != err {
		fmt.Printf("error while opening the file %s", *fileName)
		os.Exit(3)
	}
	links, err := link.Parse(file)
	if nil != err {
		fmt.Printf("error while Genearting links from file %s", *fileName)
		os.Exit(3)
	}
	fmt.Printf("%+v", links)
}
