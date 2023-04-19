package main

import (
	"fmt"
	"time"

	"github.com/beiyougufen/crawler/collect"
	"github.com/beiyougufen/crawler/proxy"
)

func main() {
	proxyURLs := []string{"http://127.0.0.1:8888", "http://127.0.0.1:8889"}
	p, err := proxy.RoundRobinProxySwitcher(proxyURLs...)
	if err != nil {
		fmt.Println("RoundRobinProxySwitcher failed")
	}
	url := "https://google.com"
	var f collect.Fetcher = collect.BrowserFetch{
		Timeout: 10 * time.Second,
		Proxy:   p,
	}

	body, err := f.Get(url)
	if err != nil {
		fmt.Printf("read content failed: %v", err)
		return
	}
	fmt.Println(string(body))
}
