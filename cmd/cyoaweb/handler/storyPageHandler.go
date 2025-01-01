package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/yedeka/Go_Projects/cmd/cyoaweb/model"
)

type HandlerOpts func(handler *StoryPageHandler) error

type StoryPageHandler struct {
	currentStory         model.Story
	RenderedChapter      string
	RenderedTemplatePath string
	RenderedTemplate     *template.Template
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

func (pageHandler *StoryPageHandler) setRenderingChapter(renderingChapter string) {
	// Check for any path that is not intro and update the path in struct to render that path's details
	if renderingChapter != "" && renderingChapter != "/" {
		renderingChapter = renderingChapter[1:]
	} else {
		renderingChapter = "intro"
	}
	pageHandler.RenderedChapter = renderingChapter
}

// ServeHTTP method returns the starting page for story.
func (pageHandler StoryPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	pageHandler.setRenderingChapter(path)
	fmt.Printf("Rendered Chapter => %s \n", pageHandler.RenderedChapter)
	if chapter, ok := pageHandler.currentStory[pageHandler.RenderedChapter]; ok {
		err := pageHandler.RenderedTemplate.Execute(w, chapter)
		if nil != err {
			log.Fatal(err.Error())
			http.Error(w, "Something went wrong ...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Could not find the Chapter provided in given story", http.StatusNotFound)
}
