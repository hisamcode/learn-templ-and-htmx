package main

import (
	"fmt"
	"net/http"
)

func integration(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("hx-trigger", "contacts-updated")

	fmt.Fprint(w, "from integration")
}

func table(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "from table")
}

func main() {

	mux := http.NewServeMux()

	mux.Handle("GET /", http.FileServer(http.Dir("../vendor/")))

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		index().Render(r.Context(), w)
	})

	mux.HandleFunc("POST /integrations/{id}", integration)
	mux.HandleFunc("GET /contacts/table", table)
	mux.HandleFunc("GET /not-found", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	err := http.ListenAndServe("127.0.0.1:8000", mux)
	panic(err)

}
