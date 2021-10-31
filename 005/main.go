package main

import (
	"flag"
	"fmt"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.ccom/", "url of the site to build sitemap for")
	flag.Parse()

	fmt.Println(*urlFlag)
	/*
		TODO:
			1. GET the webpage
			2. parse all the links on the page (use the package from 004)
			3. build proper url's for each link
			4. remove links with diff domain
			5. find all the pages (bfs)
			(steps 1 - 4 will be on each page)
			6. export XML
	*/
}
