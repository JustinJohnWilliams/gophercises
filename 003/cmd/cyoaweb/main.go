package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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

	tpl := template.Must(template.New("").Parse(cyoa.FooBarTemplate))

	h := cyoa.NewHandler(story,
		cyoa.WithTemplate(tpl),
		cyoa.WithPathParserFunc(cyoa.FooBarPathParserFn))

	mux := http.NewServeMux()
	mux.Handle("/foobar/", h)
	fmt.Printf("Starting the server on %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}
