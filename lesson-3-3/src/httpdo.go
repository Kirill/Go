package main

import (
	"net/http"
	"net/url"
)

func main() {
	reqArgs := url.Values{}
	reqArgs.Add("query", "go syntax")
	reqArgs.Add("limit", "5")
	reqUrl, _ := url.Parse("https://site.ru/search")
	reqUrl.RawQuery = reqArgs.Encode()

	req, err := http.NewRequest("GET", reqUrl.String(), nil)
	req.Header.Add("User-Agent", `Mozilla/5.0 Gecko/20100101 Firefox/39.0`)

	resp, err := http.DefaultClient.Do(req)
}
