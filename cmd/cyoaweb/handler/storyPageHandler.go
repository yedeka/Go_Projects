package handler

import (
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/yedeka/Go_Projects/cmd/cyoaweb/model"
)

type StoryPageHandler struct {
	currentStory    model.Story
	renderedChapter string
}

// ServeHTTP method returns the starting page for story.
func (pageHandler StoryPageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./ui/templates/index.html")
	if nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	path := strings.TrimSpace(r.URL.Path)
	// Check for any path that is not intro and update the path in struct to render that path's details
	if path != "" && path != "/" {
		path = path[1:]
		pageHandler.renderedChapter = path
	}

	if chapter, ok := pageHandler.currentStory[pageHandler.renderedChapter]; ok {
		err = t.Execute(w, chapter)
		if nil != err {
			log.Fatal(err.Error())
			http.Error(w, "Something went wrong ...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Could not find the Chapter provided in given story", http.StatusNotFound)
}
