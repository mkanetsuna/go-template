// src/main.go
package main

import (
	"fmt"
	"net/http"
	"html/template"
	"github.com/fsnotify/fsnotify"
)

var reloadChan = make(chan bool)

func loadTemplates() (*template.Template, error) {
	return template.ParseFiles("src/templates/index.html")
}

func handler(w http.ResponseWriter, r *http.Request) {
	templates, err := loadTemplates()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func sseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	for {
		select {
		case <-reloadChan:
			fmt.Fprintf(w, "data: reload\n\n")
			w.(http.Flusher).Flush()
		}
	}
}

func watchFiles() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("Error creating watcher:", err)
		return
	}
	defer watcher.Close()

	err = watcher.Add("src/templates")
	if err != nil {
		fmt.Println("Error adding watcher:", err)
		return
	}

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				reloadChan <- true
			}
		case err := <-watcher.Errors:
			fmt.Println("Error:", err)
		}
	}
}

func main() {
	fs := http.FileServer(http.Dir("src/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler)
	http.HandleFunc("/reload", sseHandler)
	go watchFiles()

	fmt.Println("Starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
