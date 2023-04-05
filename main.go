package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/antchfx/htmlquery"
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

	doc, err := htmlquery.Parse(bytes.NewReader(body))
	if err != nil {
		fmt.Printf("url: %v, htmlquery.Parse failed: %v", url, err)
		return
	}

	nodes := htmlquery.Find(doc, `//div[@class="new_li"]/h2/a[@target="_blank"]`)

	for _, node := range nodes {
		fmt.Println("fetch card:", node.FirstChild.Data)
	}
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
