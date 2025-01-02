package link

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Links specifies a link in a HTML document.
type Link struct {
	Href string
	Text string
}

// Parse takes a reader of HTML document and returns either a slice of links or an error.
func Parse(reader io.Reader) ([]Link, error) {
	doc, err := html.Parse(reader)
	if nil != err {
		return nil, fmt.Errorf("error while parsing html file for creating link slice")
	}
	linkNodes := extractLinkNodes(doc)
	var links []Link
	for _, node := range linkNodes {
		links = append(links, buildLinksFromNodes(node))
	}
	return links, nil
}

// extractLinkNodes function extracts all the link nodes from the given HTML root node and returns a slice of only link nodes
// link nodes are extracted by performing DFS on each html node so that we can get to anchor tags in entire dom tree.
func extractLinkNodes(node *html.Node) []*html.Node {
	// Current node is exactly an achor tag
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	var ret []*html.Node
	// For the current node parse all of it's children to perform DFS
	for element := node.FirstChild; element != nil; element = element.NextSibling {
		ret = append(ret, extractLinkNodes(element)...)
	}
	return ret
}

// buildLinksFromNodes function estracts the Href and text attributes from the anchor tags in the page and returns a corresponding
// Link object.
func buildLinksFromNodes(node *html.Node) Link {
	var result Link
	for _, attribute := range node.Attr {
		if attribute.Key == "href" {
			result.Href = attribute.Val
			break
		}
	}
	result.Text = extractTextFromNodeTree(node)
	return result
}

func extractTextFromNodeTree(node *html.Node) string {
	// Found the text node so return the text of that
	if node.Type == html.TextNode {
		return node.Data
	}
	// Non text and non HTML element
	if node.Type != html.ElementNode {
		return ""
	}
	var textResult string
	// Run DFS of passed node to extract all the text from the node
	for element := node.FirstChild; element != nil; element = element.NextSibling {
		textResult += extractTextFromNodeTree(element) + " "
	}
	return strings.Join(strings.Fields(textResult), " ")
}
