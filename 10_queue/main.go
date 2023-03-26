package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
)

func main() {
	url := "https://httpbin.org/delay/1"
	c := colly.NewCollector()

	// create a request queue with 2 consumer threads
	q, _ := queue.New(
		2, // no. of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // queue storage
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	for i := 0; i < 5; i++ {
		// add urls to queue
		q.AddURL(fmt.Sprintf("%s?n=%d", url, i))
	}
	// consuming the urls
	q.Run(c)
}
