package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gocolly/colly"
)

type PageInfo struct {
	StatusCode int
	Links      map[string]int
}

func handler(w http.ResponseWriter, r *http.Request) {
	URL := r.URL.Query().Get("url")
	if URL == "" {
		log.Println("missing URL argument")
		return
	}
	log.Println("visiting", URL)

	c := colly.NewCollector()
	p := &PageInfo{Links: make(map[string]int)}

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))
		if link != "" {
			p.Links[link]++
		}
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("response received:", r.StatusCode)
		p.StatusCode = r.StatusCode
	})

	c.Visit(URL)

	b, err := json.Marshal(p)
	if err != nil {
		log.Println("Failed to serialize response:", err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	addr := ":7171"
	http.HandleFunc("/", handler)
	log.Println("listening on", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
