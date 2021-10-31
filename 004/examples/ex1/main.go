package main

import (
	"fmt"
	"strings"

	"github.com/justinjohnwilliams/link"
)

var exampleHtml = `
<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">
  	A link to another page
	<span>some span shenanigans</span>
  </a>
  <a href="/page-two">second page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)

	links, err := link.Parse(r)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
