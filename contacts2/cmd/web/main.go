package main

import (
	"contacts2/components"
	"net/http"
	"strconv"
)

type App struct {
	contacts *components.Contacts
}

func (app App) pageListHandler(w http.ResponseWriter, r *http.Request) {
	components.Layout(components.PageContacts(*app.contacts)).Render(r.Context(), w)
}
func (app App) pageDetailHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	if len(idStr) < 1 {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		return
	}

	contact := app.contacts.FindByID(id)

	if contact == nil {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		return
	}

	form := components.NewForm()

	form.Values["ID"] = idStr
	form.Values["name"] = contact.Name
	form.Values["email"] = contact.Email
	form.Values["phone"] = contact.Phone

	components.Layout(components.PageDetailContact(*form)).Render(r.Context(), w)
}
func (app App) pageCreateHandler(w http.ResponseWriter, r *http.Request) {
	components.Layout(components.PageCreateContact(components.Form{})).Render(r.Context(), w)
}
func (app App) createHandler(w http.ResponseWriter, r *http.Request) {

	form := components.NewForm()
	form.Values["name"] = r.PostFormValue("name")
	form.Values["email"] = r.PostFormValue("email")
	form.Values["phone"] = r.PostFormValue("phone")

	var errorCount int

	if app.contacts.FindByEmail(form.Values["email"]) != nil {
		form.Errors["email"] = "email has been registered"
		errorCount++
	}

	if errorCount > 0 {
		components.Layout(components.PageCreateContact(*form)).Render(r.Context(), w)
		return
	}

	app.contacts.Add(
		form.Values["name"],
		form.Values["email"],
		form.Values["phone"],
	)

	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}
func (app App) pageEditHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	if len(idStr) < 1 {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		return
	}

	contact := app.contacts.FindByID(id)
	if contact == nil {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		return
	}

	form := components.NewForm()
	form.Values["ID"] = r.PostFormValue("ID")
	form.Values["name"] = r.PostFormValue("name")
	form.Values["email"] = r.PostFormValue("email")
	form.Values["phone"] = r.PostFormValue("phone")

	components.Layout(components.PageEditContact(*form)).Render(r.Context(), w)
}
func (app App) editHandler(w http.ResponseWriter, r *http.Request)   {}
func (app App) deleteHandler(w http.ResponseWriter, r *http.Request) {}

func main() {

	contacts := new(components.Contacts)
	contacts.Init()

	app := App{
		contacts: contacts,
	}

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("../vendor/")))

	mux.HandleFunc("/$", func(w http.ResponseWriter, r *http.Request) {})

	mux.HandleFunc("GET /contacts", app.pageListHandler)
	mux.HandleFunc("GET /contacts/{id}", app.pageDetailHandler)
	mux.HandleFunc("GET /contacts/create", app.pageCreateHandler)
	mux.HandleFunc("POST /contacts", app.createHandler)
	mux.HandleFunc("GET /contacts/{id}/edit", app.pageEditHandler)
	mux.HandleFunc("PUT /contacts/{id}", app.editHandler)
	mux.HandleFunc("DELETE /contacts/{id}", app.deleteHandler)

	var handler http.Handler = mux

	server := new(http.Server)
	server.Addr = "127.0.0.1:8000"
	server.Handler = handler

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
