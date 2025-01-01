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

type StoryPageHandler struct {
	currentStory     model.Story
	RenderedChapter  string
	RenderedTemplate *template.Template
}

func (pageHandler *StoryPageHandler) setRenderingDetails(renderingChapter string, renderingTemplate string) error {
	// Check for any path that is not intro and update the path in struct to render that path's details
	if renderingChapter != "" && renderingChapter != "/" {
		renderingChapter = renderingChapter[1:]
	} else {
		renderingChapter = "intro"
	}
	pageHandler.RenderedChapter = renderingChapter
	template, err := template.ParseFiles(renderingTemplate)
	if nil != err {
		return errors.New("error while rendering the page")
	}
	pageHandler.RenderedTemplate = template
	return nil
}

// ServeHTTP method returns the starting page for story.
func (pageHandler StoryPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	renderingError := pageHandler.setRenderingDetails(path, "./ui/templates/index.html")
	if nil != renderingError {
		log.Fatal(renderingError.Error())
		http.Error(w, renderingError.Error(), http.StatusInternalServerError)
	}
	fmt.Printf("Rendered Chapter => %s", pageHandler.RenderedChapter)
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
