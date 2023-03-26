package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func main() {
	url := "https://httpbin.org/delay/2"

	c := colly.NewCollector(
		// attached a debugger to collector
		colly.Debugger(&debug.LogDebugger{}),
		colly.Async(true),
	)

	// limit the no. of threads to 2
	// when visiting links which domains matches "*httpbin.*" glob
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*httpbin.",
		Parallelism: 2,
		RandomDelay: time.Second * 5,
	})

	// Start scraping in four threads on https://httpbin.org/delay/2
	for i := 0; i < 4; i++ {
		c.Visit(fmt.Sprintf("%s?n=%d", url, i))
	}

	// Start scraping on https://httpbin.org/delay/2
	c.Visit(url)
	// Wait until threads are finished
	c.Wait()
}
