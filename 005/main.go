package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "url of the site to build sitemap for")
	flag.Parse()

	fmt.Println(*urlFlag)

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	links, _ := link.Parse(resp.Body)

	/*
		TODO:
			1. GET the webpage (done)
			2. parse all the links on the page (use the package from 004)
			3. build proper url's for each link
			4. remove links with diff domain
			5. find all the pages (bfs)
			(steps 1 - 4 will be on each page)
			6. export XML
	*/
}
