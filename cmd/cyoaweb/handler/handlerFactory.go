package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/yedeka/Go_Projects/cmd/cyoaweb/model"
)

func ReturnStoryHandler(story model.Story, chapterToRender string) (http.Handler, error) {

	storypageHandler, err := NewstoryPageHandler(
		story,
		WithTemplateString("./ui/templates/index.html"),
	)
	if nil != err {
		return nil, errors.New(fmt.Sprintf("error while creating Handler for Chapter %s", chapterToRender))
	}
	return storypageHandler, nil
}
