package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Initialized a collector
	c := colly.NewCollector(
		// limited the domains to be visited
		colly.AllowedDomains("hackerspace.org", "wiki.hackerspace.org"),
	)

	// on every a element which has an href , call the callback
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// print the link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		// visit the link found
		// only links allowed in AllowedDomains will be visited
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// call this callback before making a request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// scraping lesgooo
	c.Visit("https://hackerspace.org/")
}
