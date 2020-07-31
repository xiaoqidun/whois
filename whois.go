package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	client := &http.Client{Transport: &http.Transport{
		DisableKeepAlives: true,
	}}
	resp, err := client.Get("https://whois.aite.xyz/?ajax&client&domain=" + url.QueryEscape(os.Args[1]))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println(string(body))
}
