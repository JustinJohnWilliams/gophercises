package main

import (
	"fmt"
	"strings"

	"github.com/justinjohnwilliams/link"
)

var exampleHtml = `
<html>
<body>
  <a href="/dog-cat">dog cat <!-- commented text SHOULD NOT be included! --></a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)

	links, err := link.Parse(r)

	if err != nil {
		panic(err)
	}

	for _, link := range links {
		fmt.Printf("%+v\n", link)
	}
}
