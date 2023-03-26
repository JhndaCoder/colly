package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		// MaxDepth is 1 so links on scraped page will be visited only
		colly.MaxDepth(1),
	)
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// print link
		fmt.Println(link)
		// visit the links found
		e.Request.Visit(link)
	})

	c.Visit("https://en.wikipedia.org/")
}
