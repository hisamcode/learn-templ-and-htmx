package main

import (
	"fmt"
	"net/http"
)

type Form struct {
	Values map[string]string
	Errors map[string]string
}

func newForm() *Form {
	return &Form{
		Values: make(map[string]string),
		Errors: make(map[string]string),
	}
}

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
		index(*newForm()).Render(r.Context(), w)
	})

	mux.HandleFunc("POST /integrations/{id}", integration)
	mux.HandleFunc("GET /contacts/table", table)
	mux.HandleFunc("GET /not-found", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})

	mux.HandleFunc("POST /validate", func(w http.ResponseWriter, r *http.Request) {
		form := newForm()
		form.Values["name"] = r.PostFormValue("name")
		if form.Values["name"] == "kopi" {
			form.Errors["name"] = "kopi not allowed"
			formKopi(*form).Render(r.Context(), w)
			return
		}

		contentOOB("dari post validate").Render(r.Context(), w)
	})

	err := http.ListenAndServe("127.0.0.1:8000", mux)
	panic(err)

}
