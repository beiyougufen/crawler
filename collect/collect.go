package collect

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/beiyougufen/crawler/proxy"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

type Fetcher interface {
	Get(r *Request) ([]byte, error)
}

type BaseFetch struct {
}

func (BaseFetch) Get(r *Request) ([]byte, error) {

	resp, err := http.Get(r.Url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("url: %v, status code: %v", r.Url, resp.StatusCode)
		return nil, err
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())

	return ioutil.ReadAll(utf8Reader)
}

type BrowserFetch struct {
	Timeout time.Duration
	Proxy   proxy.ProxyFunc
}

// 模拟浏览器访问
func (b BrowserFetch) Get(r *Request) ([]byte, error) {

	client := &http.Client{
		Timeout: b.Timeout,
	}
	if b.Proxy != nil {
		transport := http.DefaultTransport.(*http.Transport)
		transport.Proxy = b.Proxy
		client.Transport = transport
	}

	req, err := http.NewRequest("GET", r.Url, nil)
	if err != nil {
		return nil, fmt.Errorf("new req failed: %w", err)
	}

	if len(r.Cookie) > 0 {
		req.Header.Set("Cookie", r.Cookie)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("url: %v, status code: %v", r.Url, resp.StatusCode)
		return nil, err
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
