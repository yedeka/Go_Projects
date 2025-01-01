package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/yedeka/Go_Projects/cmd/cyoaweb/model"
)

type HandlerOpts func(handler *StoryPageHandler) error

type StoryPageHandler struct {
	currentStory         model.Story
	RenderedTemplatePath string
	RenderedTemplate     *template.Template
	ChapterDecider       func(r *http.Request) string
}

func WithTemplateString(templateString string) HandlerOpts {
	return func(handler *StoryPageHandler) error {
		template, err := template.ParseFiles(templateString)
		if nil != err {
			return errors.New("error while rendering the page")
		}
		handler.RenderedTemplatePath = templateString
		handler.RenderedTemplate = template
		return nil
	}
}

func WithChapterDecider(deciderFunction func(r *http.Request) string) HandlerOpts {
	return func(handler *StoryPageHandler) error {
		handler.ChapterDecider = deciderFunction
		return nil
	}
}

func NewstoryPageHandler(story model.Story, opts ...HandlerOpts) (StoryPageHandler, error) {
	storyHandler := StoryPageHandler{
		currentStory: story,
	}

	for _, options := range opts {
		err := options(&storyHandler)
		if nil != err {
			return storyHandler, err
		}

	}
	return storyHandler, nil
}

// ServeHTTP method returns the starting page for story.
func (pageHandler StoryPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := pageHandler.ChapterDecider(r)
	fmt.Printf("Rendered Chapter => %s \n", path)
	if chapter, ok := pageHandler.currentStory[path]; ok {
		err := pageHandler.RenderedTemplate.Execute(w, chapter)
		if nil != err {
			log.Fatal(err.Error())
			http.Error(w, "Something went wrong ...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Could not find the Chapter provided in given story", http.StatusNotFound)
}
