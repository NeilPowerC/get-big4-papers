package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func handleNDSSPaperUrl(url string) {
	ndss := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	ndss.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	ndss.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	ndss.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	ndss.OnHTML(".paginator a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	ndss.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	ndss.Visit(url)
}

func handleBig4(url, conf_name string) {
	if conf_name == "NDSS" {

	}
}

func main() {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://www.ndss-symposium.org/ndss2023/accepted-papers/")
}
