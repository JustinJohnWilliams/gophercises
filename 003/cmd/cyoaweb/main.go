package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/justinjohnwilliams/cyoa"
)

func main() {
	port := flag.Int("port", 9000, "port to serve on")
	file := flag.String("file", "gopher.json", "the JSON file with the story")

	flag.Parse()
	fmt.Printf("using file %s\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JSONStory(f)
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("").Parse(fooBarTemplate))

	h := cyoa.NewHandler(story,
		cyoa.WithTemplate(tpl),
		cyoa.WithPathFunc(someOtherPathFn))

	mux := http.NewServeMux()
	mux.Handle("/foobar/", h)
	fmt.Printf("Starting the server on %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func someOtherPathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/foobar" || path == "/foobar/" {
		path = "/foobar/intro"
	}

	return path[len("/foobar/"):]
}

var fooBarTemplate = `
<!DOCTYPE HTML>
<html>
    <head>
        <meta charset="utf-8"/>
        <title>Choose Your Own Adventure</title>
    </head>
    <body>
	<section class="page">
        <h1>FOOBAR {{.Title}}</h1>
        {{range .Paragraphs}}
        <p>{{.}}</p>
        {{end}}
        <ul>
            {{range .Options}}
            <li><a href="foobar/{{.Chapter}}">{{.Text}}</a></li>
            {{end}}
        </ul>
	</section>
	    <style>
      body {
        font-family: helvetica, arial;
      }
      h1 {
        text-align:center;
        position:relative;
      }
      .page {
        width: 80%;
        max-width: 500px;
        margin: auto;
        margin-top: 40px;
        margin-bottom: 40px;
        padding: 80px;
        background: #FFFCF6;
        border: 1px solid #eee;
        box-shadow: 0 10px 6px -6px #777;
      }
      ul {
        border-top: 1px dotted #ccc;
        padding: 10px 0 0 0;
        -webkit-padding-start: 0;
      }
      li {
        padding-top: 10px;
      }
      a,
      a:visited {
        text-decoration: none;
        color: #6295b5;
      }
      a:active,
      a:hover {
        color: #7792a2;
      }
      p {
        text-indent: 1em;
      }
    </style>
    </body>
</html>`
