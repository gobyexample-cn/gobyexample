package main

import (
	"fmt"
	"net/http"
	"os"
)

type HTMLDir struct {
	d http.Dir
}

func (d HTMLDir) Open(name string) (http.File, error) {
	f, err := d.d.Open(name)
	if os.IsNotExist(err) {
		if f, err := d.d.Open(name + ".html"); err == nil {
			return f, nil
		}
	}
	return f, err
}

func main() {
	port := "8000"
	publicDir := "public"
	fmt.Printf("Serving Go by Example at http://127.0.0.1:%s\n", port)
	fs := http.FileServer(HTMLDir{http.Dir(publicDir)})
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(":"+port, nil)
}
