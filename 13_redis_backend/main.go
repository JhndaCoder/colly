package main

import (
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/gocolly/redisstorage"
)

func main() {
	urls := []string{
		"http://httpbin.org/",
		"http://httpbin.org/ip",
		"http://httpbin.org/cookies/set?a=b&c=d",
		"http://httpbin.org/cookies",
	}

	c := colly.NewCollector()

	storage := &redisstorage.Storage{
		Address:  "127.0.0.1:6379",
		Password: "",
		DB:       0,
		Prefix:   "httpbin_test",
	}

	err := c.SetStorage(storage)
	if err != nil {
		panic(err)
	}

	if err := storage.Clear(); err != nil {
		log.Fatal(err)
	}

	defer storage.Client.Close()

	q, _ := queue.New(2, storage)

	c.OnResponse(func(r *colly.Response) {
		log.Println("Cookies:", c.Cookies(r.Request.URL.String()))
	})

	for _, u := range urls {
		q.AddURL(u)
	}

	q.Run(c)
}
