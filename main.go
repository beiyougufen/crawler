package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

	numLinks := strings.Count(string(body), "<a")
	fmt.Printf("homepage has %d links!\n", numLinks)

	numLinks = bytes.Count(body, []byte("<a"))
	fmt.Printf("homepage has %d links!\n", numLinks)

	exist := strings.Contains(string(body), "疫情")
	fmt.Printf("是否存在疫情：%v\n", exist)

	exist = bytes.Contains(body, []byte("疫情"))
	fmt.Printf("是否存在疫情：%v\n", exist)

}
