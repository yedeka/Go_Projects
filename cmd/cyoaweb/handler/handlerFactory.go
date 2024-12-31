package handler

import (
	"net/http"

	"github.com/yedeka/Go_Projects/cmd/cyoaweb/model"
)

func ReturnStoryHandler(story model.Story, chapterToRender string) http.Handler {

	return StoryPageHandler{
		currentStory:    story,
		renderedChapter: "intro",
	}
}
