package cyoa

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"text/template"
)

func init() {
	tpl = template.Must(template.New("").Parse(DefaultTemplate))
}

var tpl *template.Template

//READ: https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis
type HandlerOption func(h *handler)

func WithTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.t = t
	}
}

func WithPathParserFunc(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.pathParserFn = fn
	}
}

func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := handler{s, tpl, DefaultPathParserFn} // default options in handler
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

type handler struct {
	s            Story
	t            *template.Template
	pathParserFn func(r *http.Request) string
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.pathParserFn(r)

	if chapter, ok := h.s[path]; ok {
		err := h.t.Execute(w, chapter)

		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "OH NO'S!!!", http.StatusInternalServerError)
		}

		return
	}

	http.Error(w, "chapter not found", http.StatusNotFound)
}

func JSONStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
