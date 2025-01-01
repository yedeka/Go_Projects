package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/yedeka/Go_Projects/cmd/cyoaweb/model"
)

func ReturnStoryHandler(story model.Story, chapterToRender string) (http.Handler, error) {

	storypageHandler, err := NewstoryPageHandler(
		story,
		WithTemplateString("./ui/templates/index.html"),
		WithChapterDecider(setRenderingChapter),
	)
	if nil != err {
		return nil, fmt.Errorf("error while creating Handler for Chapter %s", chapterToRender)
	}
	return storypageHandler, nil
}

func setRenderingChapter(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	// Check for any path that is not intro and update the path in struct to render that path's details
	if path == "/story" || path == "/story/" {
		fmt.Println("Got index page")
		return "intro"

	}
	fmt.Println("Did not get index page")
	return path[len("/story/"):]
}
