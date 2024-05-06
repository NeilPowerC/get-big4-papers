package crawlers

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

func HandleUSENIXAcceptedPapers(url string) *[][]string {
	var USENIX_papers [][]string

	//file, err := os.Create("NDSSoutput.csv")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()

	//// 创建CSV写入器
	//writer := csv.NewWriter(file)
	//writer.Write([]string{"文章分类", "文章地址", "文章题目", "文章作者", "文章摘要", "文章下载链接", "PPT下载链接"})
	//writer.Flush()

	usenix := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	usenix.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	usenix.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	usenix.OnHTML("li.expanded.last", func(e *colly.HTMLElement) { //回调函数，查找每篇文章的子链接

		//err := e.DOM.Find("a").Click()
	})

	usenix.OnHTML("article.node.node-paper", func(e *colly.HTMLElement) { //回调函数，查找每篇文章的子链接
		e.ForEach("div.tag-box", func(i int, element *colly.HTMLElement) {
			//遍历每个article标签
			article_url := element.ChildAttr("a", "href")
			fmt.Println(article_url)
			//temp_paper_info := HandleUSENIXPaperUrl(article_url)
			//USENIX_papers = append(USENIX_papers, *temp_paper_info)
			time.Sleep(time.Second)
		})
	})

	//c.OnScraped(func(r *colly.Response) {
	//	writer := csv.NewWriter(file)
	//	for _, row := range NDSS_papers {
	//		err := writer.Write(row)
	//		if err != nil {
	//			panic(err)
	//		}
	//	}
	//	writer.Flush()
	//})

	usenix.Visit(url)
	return &USENIX_papers
}

func HandleUSENIXSessionsPapers(url string) {

}

func HandleUSENIXPaperUrl(url string) {

}
