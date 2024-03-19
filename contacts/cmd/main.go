package main

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/hisamcode/try-htmx/contacts/components"
)

type App struct {
	log      *slog.Logger
	contacts components.Contacts
}

func (app App) listHandler(w http.ResponseWriter, r *http.Request) {
	app.log.Info("list handler")
	app.render(components.PageListContact(app.contacts), w, r)
}

func (app App) render(content templ.Component, w http.ResponseWriter, r *http.Request) {
	components.Layout(content).Render(r.Context(), w)
}

func main() {
	s := slog.New(slog.Default().Handler())

	app := App{
		log:      s,
		contacts: make(components.Contacts),
	}

	app.contacts[1] = components.Contact{2, "hisam"}

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("../vendor/")))

	mux.HandleFunc("GET /{$}", app.listHandler)

	app.log.Info("listening on 127.0.0.1:8000")
	err := http.ListenAndServe("127.0.0.1:8000", mux)
	if err != nil {
		panic(err)
	}
}
