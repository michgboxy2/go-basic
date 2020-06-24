package cms

import (
	"net/http"
	"time"
)

//HandleNew preview logic
func HandleNew(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		Tmpl.ExecuteTemplate(w, "new", nil)	
	case "POST":
		title := r.FormValue("title")
		content := r.FormValue("content")
		contentType := r.FormValue("content-type")
		r.ParseForm()

		if contentType == "page" {
			Tmpl.ExecuteTemplate(w, "page", &Page{
				Title: title,
				Content: content,
				})
				return
			}

		if contentType == "post" {
			Tmpl.ExecuteTemplate(w, "post", &Post{
				Title: title,
				Content: content,
			})
			return
		}
	default:
		http.Error(w, "Method not supported: "+r.Method, http.StatusMethodNotAllowed)
		
	}
}

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	p := &Page{
		Title: "Go Projects CMS",
		Content: "Welcome to our home page",
		Posts: []*Post{
			&Post{
				Title: "Hello World",
				Content: "Hello World, thanks for coming around",
				DatePublished: time.Now(),
			},
			&Post{
				Title: "A post with comments",
				Content: "Here is a controversial post, It's sure to attract comments.",
				DatePublished: time.Now().Add(-time.Hour),
				Comments: []*Comment{
					&Comment{
						Author:        "Ben Tranter",
						Comment:       "Nevermind, I guess I just commented on my own post...",
						DatePublished: time.Now().Add(-time.Hour / 2),
					},
				},
			},
		},
	}

	Tmpl.ExecuteTemplate(w, "page", p);
}