package main

import (
	"flag"
	"fmt"
)

func main() {
	seedUrl := flag.String("url", "https://gophercises.com", "URL of the site to start building Site Map")
	flag.Parse()
	fmt.Printf("URL to start building site map => %s\n", *seedUrl)
}
