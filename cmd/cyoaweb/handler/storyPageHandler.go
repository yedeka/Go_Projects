package handler

import (
	"fmt"
	"net/http"
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
	fmt.Println("%v\n", pageHandler.currentStory)
	fmt.Println("%s\n", pageHandler.renderedChapter)
	err = t.Execute(w, pageHandler.currentStory[pageHandler.renderedChapter])
	if nil != err {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
