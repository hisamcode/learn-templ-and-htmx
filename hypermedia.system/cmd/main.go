package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("../vendor/")))

	t, err := template.ParseGlob("./ui/html/**.html")
	if err != nil {
		panic(err)
	}
	// index
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		t.Execute(w, nil)
	})
	mux.HandleFunc("GET /tt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "test@")
	})

	fmt.Println("listen on 127.0.0.1:8000")
	if err := http.ListenAndServe("127.0.0.1:8000", mux); err != nil {
		panic(err)
	}
}
