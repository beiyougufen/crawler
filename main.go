package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/beiyougufen/crawler/collect"
)

func main() {
	url := "https://book.douban.com/subject/1007305/"
	var f collect.Fetcher = collect.BrowserFetch{
		Timeout: 10 * time.Second,
	}
	body, err := f.Get(url)
	if err != nil {
		fmt.Printf("read content failed: %v", err)
		return
	}
	fmt.Println(string(body))

	// 加载HTML文档
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Printf("read content failed: %v", err)
	}

	doc.Find("div.news_li h2 a[target=_blank]").Each(func(i int, s *goquery.Selection) {
		// 获取匹配元素的文本
		title := s.Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}
