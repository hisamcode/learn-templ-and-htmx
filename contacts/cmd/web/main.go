package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"slices"
	"strconv"

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

func (app App) detailHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		app.log.Info("required id")
		return
	}

	newId, err := strconv.Atoi(id)
	if err != nil {
		app.log.Info("Cant convert int to id, err %v", err)
		return
	}

	app.render(components.PageDetailContact((*app.contacts)[newId]), w, r)
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

func (app App) deleteHandler(w http.ResponseWriter, r *http.Request) {
	pathID := r.PathValue("id")
	if pathID == "" {
		app.log.Info("id required")
		return
	}

	id, err := strconv.Atoi(pathID)
	if err != nil {
		app.log.Info(err.Error())
		return
	}

	(*app.contacts) = slices.Delete(*app.contacts, id, id+1)

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
	mux.HandleFunc("GET /{id}", app.detailHandler)
	mux.HandleFunc("POST /{$}", app.createHandler)
	mux.HandleFunc("DELETE /{id}", app.deleteHandler)

	var handler http.Handler = mux

	server := new(http.Server)
	server.Addr = "127.0.0.1:8000"
	server.Handler = handler

	app.log.Info(fmt.Sprintf("listening on %s", server.Addr))
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
