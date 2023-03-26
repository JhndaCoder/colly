package main

import (
	"net"
	"net/http"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	//! Collector Configuration

	c1 := colly.NewCollector()
	c2 := colly.NewCollector(
		colly.UserAgent("xy"),
		colly.AllowURLRevisit(),
	)
	//? or
	//? c2 := colly.NewCollector()
	//? c2.UserAgent = "xy"
	//? c2.AllowURLRevisit = true

	//! Environment config variables
	// 	ALLOWED_DOMAINS (comma separated list of domains)
	// CACHE_DIR (string)
	// DETECT_CHARSET (y/n)
	// DISABLE_COOKIES (y/n)
	// DISALLOWED_DOMAINS (comma separated list of domains)
	// IGNORE_ROBOTSTXT (y/n)
	// MAX_BODY_SIZE (int)
	// MAX_DEPTH (int - 0 means infinite)
	// PARSE_HTTP_ERROR_RESPONSE (y/n)
	// USER_AGENT (string)

	//! HTTP Configuration

	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	})
}
