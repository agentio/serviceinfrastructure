package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

var bots = []string{
	"ahrefs.com",
	"facebook.com",
	"opensiteexplorer.org",
	"bytedance.com",
	"yandex.com",
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fs := &InterceptingHandler{
		fs: http.FileServer(http.Dir("public")),
	}

	log.Fatal(http.ListenAndServe(":"+port, fs))
}

type InterceptingHandler struct {
	fs http.Handler
}

func (h *InterceptingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h.handle(w, r) {
		return
	}
	h.fs.ServeHTTP(w, r)
}

func (h *InterceptingHandler) handle(w http.ResponseWriter, r *http.Request) bool {
	// aggressively block bad user agents
	userAgents := r.Header["User-Agent"]
	if len(userAgents) == 0 {
		http.Error(w, "", http.StatusForbidden)
		return true
	}
	if r.URL.Path != "/robots.txt" {
		for _, userAgent := range userAgents {
			for _, bot := range bots {
				if strings.Contains(userAgent, bot) {
					http.Redirect(w, r, "/robots.txt", http.StatusSeeOther)
					return true
				}
			}
		}
	}

	p := r.URL.Path
	if strings.HasSuffix(p, ".js") ||
		strings.HasSuffix(p, ".png") ||
		strings.HasSuffix(p, ".ico") ||
		strings.HasSuffix(p, ".css") ||
		strings.HasSuffix(p, ".svg") {
		return false
	}

	filename := "public" + p
	if filename[len(filename)-1] != '/' {
		filename += "/"
	}
	filename += "index.html"
	log.Printf("%s", filename)
	b, err := os.ReadFile(filename)
	if err != nil {
		return false
	}
	//b = bytes.Replace(b, []byte(search), []byte(signin), 1)
	//b = bytes.Replace(b, []byte(footer), []byte(comments), 1)

	w.Write(b)
	return true
}

const search = `<div class="book-search hidden"><input type=text id=book-search-input placeholder=Search aria-label=Search maxlength=64 data-hotkeys=s/><div class="book-search-spinner hidden"></div><ul id=book-search-results></ul></div><script>document.querySelector(".book-search").classList.remove("hidden")</script>`

const signin = `<a href="/signin" class="active">Sign in</a> to comment` + search

const footer = `<footer `

const comments = `<hr/><div>
<h4>Got something to add?</h4>
<p><a href="/signin">Sign in</a> to comment.</a></p>
</div>
<footer `
