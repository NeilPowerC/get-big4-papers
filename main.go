package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

func handleNDSSPaperUrl(url string) *[]string {
	var res []string
	res = append(res, url)
	ndss := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	ndss.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	//ndss.OnError(func(_ *colly.Response, err error) {
	//	fmt.Println("Something went wrong:", err)
	//})

	//ndss.OnResponse(func(r *colly.Response) {
	//	fmt.Println("Visited", r.Request.URL)
	//})

	ndss.OnHTML("h1.entry-title", func(e *colly.HTMLElement) { //回调函数，查找每篇文章的子链接
		fmt.Println(e.Text)
		var paper_title = e.Text
		res = append(res, paper_title)
	})

	ndss.OnHTML("div.paper-data", func(e *colly.HTMLElement) { //回调函数，查找每篇文章的子链接
		e.ForEach("p", func(i int, element *colly.HTMLElement) {
			switch i {
			case 0:
				var authors = element.ChildText("strong")
				res = append(res, authors)
			case 2:
				var abstruct = element.Text
				res = append(res, abstruct)
			}
		})
	})

	ndss.OnHTML("div.paper-buttons a[href]", func(e *colly.HTMLElement) { //回调函数，查找每篇文章的子链接
		fmt.Println("find paper url")
		fmt.Println(e.Attr("href"))
		var paper_download_url = e.Attr("href")
		res = append(res, paper_download_url)
	})

	ndss.Visit(url)
	return &res
}

func handleBig4(url, conf_name string) {
	if conf_name == "NDSS" {

	}
}

func main() {
	var NDSS_papers [][]string

	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	c.OnHTML("div.paper-list", func(e *colly.HTMLElement) { //回调函数，查找每篇文章的子链接
		e.ForEach("div.tag-box", func(i int, element *colly.HTMLElement) {
			//遍历每个article标签
			article_url := element.ChildAttr("a", "href")
			temp_paper_info := handleNDSSPaperUrl(article_url)
			NDSS_papers = append(NDSS_papers, *temp_paper_info)
			time.Sleep(time.Second)
		})

	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	c.Visit("https://www.ndss-symposium.org/ndss2023/accepted-papers/")
}
