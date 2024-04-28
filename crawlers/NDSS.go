package crawlers

import (
	"fmt"
	"github.com/gocolly/colly"
	"regexp"
	"strings"
	"time"
)

func HandleNDSSAcceptedPapersUrl(url string) *[][]string {
	var NDSS_papers [][]string

	//file, err := os.Create("NDSSoutput.csv")
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()

	//// 创建CSV写入器
	//writer := csv.NewWriter(file)
	//writer.Write([]string{"文章分类", "文章地址", "文章题目", "文章作者", "文章摘要", "文章下载链接", "PPT下载链接"})
	//writer.Flush()

	ndss := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	ndss.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	ndss.OnError(func(_ *colly.Response, err error) {
		fmt.Println("Something went wrong:", err)
	})

	ndss.OnHTML("div.paper-list", func(e *colly.HTMLElement) { //回调函数，查找每篇文章的子链接
		e.ForEach("div.tag-box", func(i int, element *colly.HTMLElement) {
			//遍历每个article标签
			article_url := element.ChildAttr("a", "href")
			temp_paper_info := HandleNDSSPaperUrl(article_url)
			NDSS_papers = append(NDSS_papers, *temp_paper_info)
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

	ndss.Visit(url)
	return &NDSS_papers
}

func HandleNDSSPaperUrl(url string) *[]string {
	res := []string{"NDSS2023"}
	res = append(res, url)
	ndss := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	ndss.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	ndss.OnHTML("h1.entry-title", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		var paper_title = e.Text
		res = append(res, paper_title)
	})

	ndss.OnHTML("div.paper-data", func(e *colly.HTMLElement) {
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
		var paper_download_url = e.Attr("href")
		res = append(res, paper_download_url)
	})

	ndss.Visit(url)
	return &res
}

func HandleNDSSProgramUrl(url string) *[]string {
	res := []string{"NDSS2023"}
	res = append(res, url)
	ndss := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"),
	)

	ndss.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	ndss.OnHTML("a.list-group-item.list-group-item-warning.card-subheading-session div.col-5", func(e *colly.HTMLElement) {
		var ori_tag = e.ChildText("strong")
		//fmt.Println(ori_tag)
		pattern := `:\s(.*)$`
		re := regexp.MustCompile(pattern)
		match := re.FindStringSubmatch(ori_tag)
		var grouping = false
		if len(match) > 1 {
			tag := match[1]
			tag_arr := strings.Split(tag, " ")
			switch tag_arr[len(tag_arr)-1] {
			case "I", "II", "III", "IV", "V":
				grouping = true
			}
			if grouping {
				tag = strings.Join(tag_arr[0:len(tag_arr)-1], " ")
			}
			fmt.Println(tag)
		} else {
			fmt.Println("未找到匹配的内容")
		}
	})

	ndss.OnHTML("ul.list-group.list-group-session.card-collapse.collapse.show", func(e *colly.HTMLElement) {
		e.ForEach("div.col-10", func(i int, element *colly.HTMLElement) {
			var title = element.ChildText("a")
			fmt.Println("题目", title)
			var paper_url = element.ChildAttr("a", "href")
			fmt.Println("文章链接", paper_url)
		})
	})

	ndss.Visit(url)
	return &res
}
