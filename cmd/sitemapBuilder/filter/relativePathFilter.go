package filter

import (
	"strings"

	"github.com/yedeka/Go_Projects/cmd/linkparser/link"
)

type RelativePathFilter struct {
	Domain string
}

func (relativePathFilter RelativePathFilter) IsApplicable(urlLink link.Link) bool {
	return strings.HasPrefix(urlLink.Href, "/")
}

func (relativePathFilter RelativePathFilter) Apply(urlLink link.Link) (link.Link, error) {
	var modifiedURL strings.Builder
	modifiedURL.WriteString(relativePathFilter.Domain)
	modifiedURL.WriteString(urlLink.Href)
	urlLink.Href = modifiedURL.String()
	return urlLink, nil
}
