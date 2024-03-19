package main

import (
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/hisamcode/try-htmx/contacts/components"
)

type App struct {
	log      *slog.Logger
	contacts *components.Contacts
}

// listHandler list contacts
func (app App) listHandler(w http.ResponseWriter, r *http.Request) {
	app.render(components.PageListContact(*app.contacts), w, r)
}

// createPageHandler page create contact
func (app App) createPageHandler(w http.ResponseWriter, r *http.Request) {
	app.render(components.PageCreateContact(components.NewFormData()), w, r)
}

// createHandler create post handler
func (app App) createHandler(w http.ResponseWriter, r *http.Request) {

	name := r.PostFormValue("name")
	email := r.PostFormValue("email")

	if app.contacts.HasEmail(email) {
		form := components.NewFormData()
		form.Values["name"] = name
		form.Values["email"] = email
		form.Errors["email"] = "email duplicate"

		app.render(components.PageCreateContact(form), w, r)
		return
	}

	app.contacts.New(name, email)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

// render render layout
func (app App) render(content templ.Component, w http.ResponseWriter, r *http.Request) {
	components.Layout(content).Render(r.Context(), w)
}

func main() {
	s := slog.New(slog.Default().Handler())

	app := App{
		log:      s,
		contacts: components.Contacts{}.Init(),
	}

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("../vendor/")))

	mux.HandleFunc("GET /{$}", app.listHandler)
	mux.HandleFunc("GET /create", app.createPageHandler)
	mux.HandleFunc("POST /{$}", app.createHandler)

	app.log.Info("listening on 127.0.0.1:8000")
	err := http.ListenAndServe("127.0.0.1:8000", mux)
	if err != nil {
		panic(err)
	}
}
