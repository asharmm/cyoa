package cyoa


import (
	"encoding/json"
	"io"
	"net/http"
	"html/template"
	"strings"	
	"log"
)


func init() {
	tmplt = template.Must(template.New("").Parse(DefaultHandlerTmplt))

}

var tmplt *template.Template

var DefaultHandlerTmplt = `
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Choose Your Own Adventure</title>
</head>

<body>
    <div class="wrapper">
        <div class="story-wrapper">
            <h1>{{.Title}}</h1>
            {{range .Paragraphs}}
            <p>{{.}}</p>
            {{end}}

            <ul>
                {{range .Options}}
                <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
                {{end}}
            </ul>
        </div>
    </div>
    <style>
        body {
            background-position: center center;
            background-image: url('https://images.unsplash.com/photo-1519681393784-d120267933ba');
            background-repeat: no-repeat;
            background-size: cover;
            height: 100vh;
            backdrop-filter: blur(5px);
            margin: 0;
        }

        .wrapper {

            width: 80%;
            margin: auto;
            padding-top: 10%;
        }

        .story-wrapper {
            color: white;
            padding: 1rem;
            backdrop-filter: blur(13px) saturate(200%);
            -webkit-backdrop-filter: blur(13px) saturate(200%);
            background-color: rgba(17, 25, 40, 0.75);
            border-radius: 12px;
            border: 1px solid rgba(255, 255, 255, 0.125);
        }
		a {
			color: white; 
			text-decoration: none; 
			font-weight: bold; 
		}
    </style>
</body>

</html>
`

func NewHandler(s Story) http.Handler{
	return handler{s}
}

type handler struct {
	s Story
}


func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {


	path := strings.TrimSpace(r.URL.Path)

	if path == "" || path == "/" {
		path = "/intro"
	}

	path = path[1:]

	if chapter, ok := h.s[path]; ok {
		err := tmplt.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Somthing Went Wrong...", http.StatusInternalServerError)
			panic(err)
		}	

		return
	}

	http.Error(w, "Chapter not found", http.StatusNotFound)
}


func JsonStory(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)
	var story Story
	
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil

}

type Story map[string]Chapter

type Chapter struct {
	Title string `json:"title"`
	Paragraphs []string `json:"story"`
	Options []Option `json:"options"`

}


type Option struct {
	Text string `json:"text"`
	Chapter string `json:"chapter"`
}