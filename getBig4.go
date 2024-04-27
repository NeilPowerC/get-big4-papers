package main

import (
	"encoding/csv"
	"fmt"
	"get-big4-papers/crawlers"
	"github.com/gocolly/colly"
	"os"
	"time"
)

func handleBig4(url, conf_name string) {
	if conf_name == "NDSS" {

	}
}

func main() {
	var NDSS_papers [][]string
	file, err := os.Create("NDSSoutput.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// 创建CSV写入器
	writer := csv.NewWriter(file)
	writer.Write([]string{"文章分类", "文章地址", "文章题目", "文章作者", "文章摘要", "文章下载链接", "PPT下载链接"})
	writer.Flush()
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
			temp_paper_info := crawlers.HandleNDSSPaperUrl(article_url)
			NDSS_papers = append(NDSS_papers, *temp_paper_info)
			time.Sleep(time.Second)
		})
	})

	c.OnScraped(func(r *colly.Response) {
		writer := csv.NewWriter(file)
		for _, row := range NDSS_papers {
			err := writer.Write(row)
			if err != nil {
				panic(err)
			}
		}
		writer.Flush()
	})

	c.Visit("https://www.ndss-symposium.org/ndss2023/accepted-papers/")
}
