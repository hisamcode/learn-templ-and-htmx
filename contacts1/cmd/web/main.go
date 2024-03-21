package main

import (
	"contacts1/components"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
)

type App struct {
	contacts *components.Contacts
}

func (app App) listPageHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page, err := strconv.Atoi(q.Get("page"))
	if err != nil {
		page = 1
	}
	search := q.Get("q")
	var contacts components.Contacts
	if len(search) > 0 {
		contacts = app.contacts.Search(search)
	} else {
		contacts = app.contacts.All(page)
	}

	app.render(components.PageList(contacts, page), w, r)

}
func (app App) detailPageHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.GetID(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	contact, err := app.contacts.FindByID(id)
	if err != nil {
		fmt.Println("not found")
		return
	}
	app.render(components.PageDetail(*contact), w, r)
}
func (app App) createPageHandler(w http.ResponseWriter, r *http.Request) {
	app.render(components.PageCreate(*components.NewFormContact()), w, r)
}
func (app App) createHandler(w http.ResponseWriter, r *http.Request) {

	name := r.PostFormValue("name")
	email := r.PostFormValue("email")

	form := components.NewFormContact()
	form.Values["name"] = name
	form.Values["email"] = email

	foundError := 0

	if len(name) < 1 || len(email) < 1 {
		form.Errors["name"] = "name tidak boleh kosong"
		form.Errors["email"] = "email tidak boleh kosong"
		foundError++
	}

	// cek apakah ada email yang duplicate
	if _, err := app.contacts.FindByEmail(email); err == nil {
		form.Errors["email"] = "email duplicate"
		foundError++
	}

	if foundError > 0 {
		app.render(components.PageCreate(*form), w, r)
		return
	}

	app.contacts.New(name, email)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (app App) deleteHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.GetID(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	app.contacts.DeleteByID(id)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (app App) editPageHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.GetID(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	contact, err := app.contacts.FindByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	form := components.NewFormContact()
	form.Values["name"] = contact.Name
	form.Values["email"] = contact.Email

	app.render(components.PageEdit(id, *form), w, r)

}

func (app App) validateEmail(w http.ResponseWriter, r *http.Request) {

	id, err := app.GetID(r)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "cant get ID")
		return
	}

	email := r.PostFormValue("email")
	if email == "" || len(email) < 1 {
		fmt.Fprint(w, "email cant be empty")
		return
	}

	contact, err := app.contacts.FindByID(id)
	if err != nil {
		fmt.Println(err)
		fmt.Fprint(w, "cant find contact by id")
		return
	}

	if contact.Email != email {
		if _, err := app.contacts.FindByEmail(email); err == nil {
			fmt.Fprint(w, "duplicate email")
			return
		}
	}
}

func (app App) editHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.GetID(r)
	if err != nil {
		fmt.Println(err)
		return
	}

	name := r.PostFormValue("name")
	email := r.PostFormValue("email")

	form := components.NewFormContact()
	form.Values["name"] = name
	form.Values["email"] = email

	foundError := 0

	if len(name) < 1 || len(email) < 1 {
		form.Errors["name"] = "name tidak boleh kosong"
		form.Errors["email"] = "email tidak boleh kosong"
		foundError++
	}

	// cek apakah ada email yang duplicate
	contact, err := app.contacts.FindByID(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	if email != contact.Email {
		if _, err := app.contacts.FindByEmail(email); err == nil {
			form.Errors["email"] = "email duplicate"
			foundError++
		}
	}

	if foundError > 0 {
		app.render(components.PageEdit(id, *form), w, r)
		return
	}

	contact.Email = email
	contact.Name = name

	app.contacts.UpdateByID(id, contact)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app App) render(content templ.Component, w http.ResponseWriter, r *http.Request) {
	components.Layout("Contacts", content).Render(r.Context(), w)
}

func (app App) GetID(r *http.Request) (int, error) {
	stringID := r.PathValue("id")

	if stringID == "" {
		fmt.Println("delete id required")
		return -1, errors.New("id cant empty")
	}

	id, err := strconv.Atoi(stringID)
	if err != nil {
		fmt.Println(err)
		return -1, errors.New(err.Error())
	}

	return id, nil
}

func main() {

	app := App{
		contacts: &components.Contacts{},
	}

	app.contacts.Init()

	mux := http.NewServeMux()

	http.Handle("/", http.FileServer(http.Dir("../vendor/")))

	mux.HandleFunc("GET /{$}", app.listPageHandler)
	mux.HandleFunc("GET /contacts/{id}", app.detailPageHandler)
	mux.HandleFunc("GET /contacts/create", app.createPageHandler)
	mux.HandleFunc("POST /contacts", app.createHandler)
	mux.HandleFunc("DELETE /contacts/{id}", app.deleteHandler)
	mux.HandleFunc("GET /contacts/{id}/edit", app.editPageHandler)
	mux.HandleFunc("PUT /contacts/{id}", app.editHandler)
	// validate email
	mux.HandleFunc("PUT /contacts/{id}/email", app.validateEmail)

	var handler http.Handler = mux

	server := new(http.Server)
	server.Addr = "127.0.0.1:8000"
	server.Handler = handler

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
