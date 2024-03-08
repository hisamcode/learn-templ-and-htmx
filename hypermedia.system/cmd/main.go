package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/hisamcode/try-htmx/hs/components"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("../vendor/")))

	mux.HandleFunc("GET /contacts", func(w http.ResponseWriter, r *http.Request) {
		components.Contacts([]string{"contact1", "contacts 2"}).Render(r.Context(), w)
	})

	mux.HandleFunc("POST /contacts", func(w http.ResponseWriter, r *http.Request) {
		components.Contacts([]string{"post 1", "post 2", r.FormValue("q")}).Render(r.Context(), w)
	})

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		cs := []templ.Component{
			components.ButtonGetContact(),
			components.ButtonGetContact2(),
			components.ButtonGetContact3(),
		}
		components.Layout("htmx", cs).Render(r.Context(), w)
	})

	fmt.Println("listen on 127.0.0.1:8000")
	if err := http.ListenAndServe("127.0.0.1:8000", mux); err != nil {
		panic(err)
	}
}
