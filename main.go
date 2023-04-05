package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://www.thepaper.cn"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("url: %v, error: %v", url, err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("url: %v, status code: %v", url, resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("url: %v, read body failed: %v", url, err)
		return
	}

	fmt.Printf("url: %v, body: %v", url, string(body))
}
