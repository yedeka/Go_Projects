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
	seedUrl := flag.String("url", "https://gophercises.com", "URL of the site to start building Site Map")
	flag.Parse()
	htmlResponse, err := processURL(*seedUrl)
	if nil != err {
		fmt.Println(err.Error())
	}
	linksList, err := extractSanitizeLinks(htmlResponse)
	if nil != err {
		fmt.Println(err.Error())
	}
	defer htmlResponse.Body.Close()
	fmt.Printf("%+v", linksList)
}

// processURL function takes a URL, fires a get call to fetch the cotent of the web page and return corresponding reader to caller.
func processURL(siteMapUrl string) (*http.Response, error) {
	resp, err := http.Get(siteMapUrl)
	if err != nil {
		return nil, fmt.Errorf("error while getting response from %s", siteMapUrl)
	}
	return resp, nil
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
