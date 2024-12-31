package model

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []option `json:"options"`
}

type option struct {
	Text        string `json:"text"`
	ChapterName string `json:"arc"`
}
