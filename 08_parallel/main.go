package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(true),
	)

	// Limit the max parallelism to 2
	// necessary to control the limit of simultaneous requests if goroutines are created dynamically
	// can also be controlled by spawning fixed no. of goroutines
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println(link)
		e.Request.Visit(link)
	})

	c.Visit("https://en.wikipedia.org/")
	c.Wait()
}
