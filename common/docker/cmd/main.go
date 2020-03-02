package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func main() {
	p, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	p = path.Join(p, "static")
	http.Handle("/file/",
		http.StripPrefix("/file/", http.FileServer(http.Dir(p))))
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	err := http.ListenAndServe(":8088", nil)
	if err != nil {
		fmt.Println(err)
	}
}
