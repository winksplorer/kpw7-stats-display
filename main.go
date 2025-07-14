package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
)

var frontendDir string = "./frontend"
var port string = ":42799"

func main() {
	http.HandleFunc("/", templater)
	http.HandleFunc("/ping", ping)

	log.Printf("kpw7-stats-display on port %s\n", port)
	if err := http.ListenAndServe(":42799", nil); err != nil {
		log.Println("http error:", err)
	}
}

func templater(w http.ResponseWriter, r *http.Request) {
	// if it's not html don't touch it
	if !strings.HasSuffix(r.URL.Path, ".html") {
		http.FileServer(http.Dir(frontendDir)).ServeHTTP(w, r)
		return
	}

	// template together base + the page
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/index.html", frontendDir))
	if err != nil {
		log.Printf("template parse error for: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, map[string]interface{}{
		// "Title": pageName + " - " + hostname,
	})
	if err != nil {
		log.Printf("template exec error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
