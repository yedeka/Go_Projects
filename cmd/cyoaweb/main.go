package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/yedeka/Go_Projects/cmd/cyoaweb/parser"
)

func main() {
	storyFile := flag.String("fileName", "gopher.json", "The JSON file containing text for Choose Your Own Adventure web application")
	flag.Parse()

	file, err := os.Open(*storyFile)
	if nil != err {
		fmt.Println("Error while opening the file, ", *storyFile)
		os.Exit(3)
	}
	jsonParser := parser.StoryParser{
		ParserType: "JSONParser",
		File:       file,
	}
	story, err := jsonParser.Parse()

	if err != nil {
		fmt.Println(err.Error())

	}
	fmt.Println(story)
}
