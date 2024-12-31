package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/yedeka/Go_Projects/cmd/cyoaweb/model"
)

type StoryParser struct {
	ParserType string
	File       *os.File
}

func (parser *StoryParser) Parse() (any, error) {
	fmt.Printf("Type of parser used is %s \n", parser.ParserType)
	decoder := json.NewDecoder(parser.File)
	var story model.Story

	if err := decoder.Decode(&story); err != nil {
		fmt.Println("Error during mapping JSON to Story object")
		return nil, errors.New("error while creating JSON decoder from input file")
	}
	return story, nil
}
