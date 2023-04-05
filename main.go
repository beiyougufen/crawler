package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

func main() {
	url := "https://www.thepaper.cn"
	body, err := Fetch(url)

	if err != nil {
		fmt.Printf("url: %v, read body failed: %v", url, err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	if err != nil {
		fmt.Printf("url: %v, goquery.NewDocumentFromReader failed: %v", url, err)
		return
	}

	doc.Find("div.new_li h2 a[target=_blank]").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}

func Fetch(url string) ([]byte, error) {

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("url: %v, status code: %v", url, resp.StatusCode)
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func DeterminEncoding(r *bufio.Reader) encoding.Encoding {

	bytes, err := r.Peek(1024)

	if err != nil {
		fmt.Printf("fetch error: %v", err)
		return unicode.UTF8
	}

	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}
