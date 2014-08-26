package main

import (
	"log"
	"net/http"
	"regexp"
)

var pathToFile = regexp.MustCompile("\\..+$")

func handler(w http.ResponseWriter, r *http.Request) {

	m := pathToFile.FindStringSubmatch(r.URL.Path)
	path := ""

	if m == nil {
		log.Printf(">>>> m: is nil")
		path = "index.html"
	} else {
		log.Printf(">>>> m: " + m[0])
		path = r.URL.Path[1:]
	}

	http.ServeFile(w, r, "static/"+path)
	//log.Printf(">>>>   path: " + path)

	// t, _ := template.ParseFiles("static/" + path)
	// t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
