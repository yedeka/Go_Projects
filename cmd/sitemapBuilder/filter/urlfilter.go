package filter

import "github.com/yedeka/Go_Projects/cmd/linkparser/link"

type UrlFilter interface {
	IsApplicable(link.Link) bool
	Apply(link.Link) (link.Link, error)
}
