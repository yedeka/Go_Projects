package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"

	"github.com/yedeka/Go_Projects/cmd/linkparser/link"
	"github.com/yedeka/Go_Projects/cmd/sitemapBuilder/filter"
)

func main() {

	linksList, err := startProcess()
	if nil != err {
		fmt.Println(err.Error())
	}
	fmt.Printf("%+v", linksList)
}

func startProcess() (map[string]struct{}, error) {
	urlVal, deapth := processFlags()
	linksList, err := PrepareSiteMap(urlVal, deapth)
	if nil != err {
		return nil, err
	}
	return linksList, nil
}

func processFlags() (string, int) {
	seedUrl := flag.String("url", "https://gophercises.com", "URL of the site to start building Site Map")
	depth := flag.Int("deapth", 3, "Deapth of parsing links on a page")
	flag.Parse()
	return *seedUrl, *depth
}

func PrepareSiteMap(seedURL string, depth int) (map[string]struct{}, error) {
	visitedURLs := make(map[string]struct{})
	var queue map[string]struct{}
	levelQueue := map[string]struct{}{
		seedURL: {},
	}

	for i := 0; i <= depth; i++ {
		queue, levelQueue = levelQueue, make(map[string]struct{})
		for url := range queue {
			// Parse the URL at hand only if it is not visited
			if _, ok := visitedURLs[url]; !ok {
				childrenLinks, err := processURL(url)
				if nil != err {
					return nil, err
				}
				// Parse the links to extract text to be given in SiteMap
				for _, link := range childrenLinks {
					levelQueue[link.Href] = struct{}{}
				}
				visitedURLs[url] = struct{}{}
			}
		}
	}
	return visitedURLs, nil
}

// processURL function takes a URL, fires a get call to fetch the cotent of the web page and return corresponding reader to caller.
func processURL(siteMapUrl string) ([]link.Link, error) {
	resp, err := http.Get(siteMapUrl)
	if err != nil {
		return nil, fmt.Errorf("error while getting response from %s", siteMapUrl)
	}
	filteredLinks, err := extractSanitizeLinks(resp)
	if nil != err {
		return nil, err
	}
	defer resp.Body.Close()
	return filteredLinks, nil
}

func extractSanitizeLinks(responseHTML *http.Response) ([]link.Link, error) {
	requestUrl := responseHTML.Request.URL
	baseUrl := &url.URL{
		Scheme: requestUrl.Scheme,
		Host:   requestUrl.Host,
	}
	base := baseUrl.String()
	links, err := link.Parse(responseHTML.Body)
	if nil != err {
		return nil, fmt.Errorf("could not extract links from %s", requestUrl)
	}
	return sanitizeLinks(links, base), nil
}

func sanitizeLinks(linksList []link.Link, domain string) []link.Link {
	filterList := []filter.UrlFilter{
		filter.RelativePathFilter{
			Domain: domain,
		}, filter.SameDomainFilter{
			Domain:           domain,
			ValidationRegExp: `https?://[^\s]+`,
		},
	}
	var sanitizedList []link.Link

	for _, link := range linksList {
		for _, filter := range filterList {
			if filter.IsApplicable(link) {
				filteredLink, err := filter.Apply(link)
				if nil != err {
					//fmt.Printf("Could not filter the link, skipping Link %+v", filteredLink)
					break
				}
				sanitizedList = append(sanitizedList, filteredLink)
				break
			}
		}
	}
	return sanitizedList
}
