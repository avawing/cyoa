package cyoa

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

type Story map[string]Chapter

var defaultHandlerTemplate = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>CHOOSE YOUR OWN ADVENTURE</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    {{range .Paragraphs}}
        <p>{{.}}</p>
    {{end}}
    <ul>
        {{range .Options}}
        <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</body>
</html>`

var tpl *template.Template

//if tmpl doesn't compile, then it's bad -> Must enforces compilation
func init() {
	tpl = template.Must(template.New("").Parse(defaultHandlerTemplate))
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSpace(r.URL.Path)
	if path == "" || path == "/" {
		path = "/intro"
	}
	// remove preceding slash
	path = path[1:]

	// anytime you check a map, returns value & boolean
	// ok acts as a failsafe - if it's not there, it won't do the thing
	// kind of opposite error first
	if chapter, ok := h.s[path]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

func NewHandler(s Story) http.Handler {
	//go idiom, return structs - accept interfaces
	return handler{s}
}

func JsonStory(reader io.Reader) (Story, error) {
	d := json.NewDecoder(reader)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
