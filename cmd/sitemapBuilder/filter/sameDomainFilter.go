package filter

import (
	"fmt"
	"regexp"

	"github.com/yedeka/Go_Projects/cmd/linkparser/link"
)

type SameDomainFilter struct {
	Domain           string
	ValidationRegExp string
}

func (sdFilter SameDomainFilter) IsApplicable(urlLink link.Link) bool {
	// Simple check to see if we have a URL starting with http:// or https:// followed by non white space characters
	// regex := regexp.MustCompile(`https?://[^\s]+`)
	regex := regexp.MustCompile(sdFilter.ValidationRegExp)
	return regex.MatchString(urlLink.Href)
}

func (sdFilter SameDomainFilter) Apply(urlLink link.Link) (link.Link, error) {
	// extract the domain from passed URL
	regex := regexp.MustCompile(`^http(?:s?):\/\/([^\/]+?)(?:\/|$)`)
	tokensList := regex.FindStringSubmatch(urlLink.Href)
	candiateDomain := &tokensList[1]
	sourceDomain := &sdFilter.Domain
	if *candiateDomain == *sourceDomain {
		return urlLink, nil
	}
	return urlLink, fmt.Errorf("input URL domain %s is not of the same domain %s as Seed URL", *candiateDomain, *sourceDomain)
}
