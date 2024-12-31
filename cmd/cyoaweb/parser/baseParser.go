package parser

type BaseParser interface {
	Parse(string) (any, error)
}
