package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type fileHandler struct {
	root http.FileSystem
	h    http.Handler
}

func fileServer(root http.FileSystem, h http.Handler) http.Handler {
	return &fileHandler{root, h}
}

func (f *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if _, err := os.Stat("public/" + path); os.IsNotExist(err) {
		http.ServeFile(w, r, "public/404.html")
		return
	}
	f.h.ServeHTTP(w, r)
}

func main() {
	// Simple static webserver:
	port := flag.String("port", ":8080", "localhost port to serve")
	path := flag.String("path", "public", "public files path")
	flag.Parse()
	log.Printf("Serving \x1b[1m%s\x1b[0m at: http://localhost\x1b[1m%s\x1b[0m\n", *path, *port)
	fs := http.Dir(*path)
	http.Handle("/", fileServer(&fs, http.FileServer(&fs)))
	err := http.ListenAndServe(*port, nil)
	if err != nil {
		log.Fatal("ListenAndServe fail:", err)
	}
}
