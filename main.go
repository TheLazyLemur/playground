package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"sync"
)

//go:embed static/index.js
var indexjs string

type data struct {
	sync.Mutex
	Count int
}

var dataContainer data

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		TemplComponent().Render(r.Context(), w)
	})

	mux.HandleFunc("/page1", func(w http.ResponseWriter, r *http.Request) {
		dataContainer.Lock()
		defer dataContainer.Unlock()
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		ReactComponent(dataContainer.Count).Render(r.Context(), w)
	})

	mux.HandleFunc("/click", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		click().Render(r.Context(), w)
	})

	mux.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		dataContainer.Lock()
		defer dataContainer.Unlock()
		w.Header().Set("Content-Type", "application/json")
		dataContainer.Count++
		w.Write(fmt.Appendf(nil, `{"count":%d}`, dataContainer.Count))
	})

	mux.HandleFunc("/index.js", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/javascript")
		w.Write([]byte(indexjs))
	})

	fmt.Println("listening on localhost:8080")
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		log.Printf("error listening: %v", err)
	}
}
