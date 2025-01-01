package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/yedeka/Go_Projects/cmd/cyoaweb/handler"
	"github.com/yedeka/Go_Projects/cmd/cyoaweb/parser"
)

func main() {
	storyFile := flag.String("fileName", "gopher.json", "The JSON file containing text for Choose Your Own Adventure web application")
	port := flag.Int("serverPort", 3000, "The port to start server")
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
	storyHandler, err := handler.ReturnStoryHandler(story, "intro")
	if nil != err {
		errors.New("Application startup failed to start Story Handler")
		os.Exit(3)
	}
	fmt.Printf("Starting the server on %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), storyHandler))
}
