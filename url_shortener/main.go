package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"

	url_shortener "github.com/Zeddling/gophercises/url_shortener/shortener"
)

func main() {
	mux := defaultMux()

	//	Build the MapHandler using the mux as the handler
	paths := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	m_handler := url_shortener.MapHandler(paths, mux)

	//	Build YAMLHandler using map handler as fallback

	var path string
	flag.StringVar(&path, "f", "data.yml", "Defaulted to saved string")

	data := url_shortener.Read(path)

	ext := filepath.Ext(path)

	handler, err := url_shortener.FileHandler([]byte(data), ext, m_handler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", handler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world")
}
