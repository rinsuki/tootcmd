package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	host := os.Getenv("MASTODON_HOST")
	if host == "" {
		fmt.Println("MASTODON_HOST env is missing.")
		fmt.Println("Usage: ")
		fmt.Printf("  MASTODON_HOST=mastodon.example.com MASTODON_ACCESS_TOKEN=xxx %s toot content\n", os.Args[0])
		os.Exit(1)
	}
	requestUrl := fmt.Sprintf("https://%s/api/v1/statuses", os.Getenv("MASTODON_HOST"))

	values := url.Values{}
	values.Add("status", strings.Join(os.Args[1:], " "))

	req, err := http.NewRequest("POST", requestUrl, strings.NewReader(values.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("MASTODON_ACCESS_TOKEN")))
	req.Header.Set("User-Agent", "go-tootcmd/1.0 (+https://github.com/rinsuki/tootcmd)")
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}
