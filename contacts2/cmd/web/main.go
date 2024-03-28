package main

import (
	"contacts2/components"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type App struct {
	contacts *components.Contacts
}

func (app App) pageListHandler(w http.ResponseWriter, r *http.Request) {

	var err error

	pagination := *components.NewPagination()
	pagination.MaxPage = int(math.Ceil(float64(app.contacts.Count) / float64(pagination.Limit)))

	if r.URL.Query().Has("page") {
		pageStr := r.URL.Query().Get("page")
		pagination.Page, err = strconv.Atoi(pageStr)
		if err != nil {
			http.Redirect(w, r, "/contacts", http.StatusSeeOther)
			return
		}
	}

	hxTrigger, ok := r.Header[http.CanonicalHeaderKey("hx-trigger")]
	if ok {
		if hxTrigger[0] == "button-pagination-next" || hxTrigger[0] == "button-pagination-prev" {
			components.PageContacts(*app.contacts.Paging(pagination), pagination).Render(r.Context(), w)
			return
		}
	}

	if r.URL.Query().Has("q") {
		q := r.URL.Query().Get("q")

		if len(q) > 0 {
			cs := app.contacts.Search(q)
			if len(cs.Data) < 0 {
				// TODO: not found
			}

			components.PageContacts(*cs, pagination).Render(r.Context(), w)
			// components.ListContacts(*cs).Render(r.Context(), w)
			return
		}
	}

	components.Layout(components.PageContacts(*app.contacts.Paging(pagination), pagination)).Render(r.Context(), w)
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
		form.Errors["email"] = "duplicate email"
		errorCount++
	}

	if len(form.Values["name"]) < 1 || form.Values["name"] == "" {
		form.Errors["name"] = "name cant empty"
		errorCount++
	}
	if len(form.Values["email"]) < 1 || form.Values["email"] == "" {
		form.Errors["email"] = "email cant empty"
		errorCount++
	}
	if len(form.Values["phone"]) < 1 || form.Values["phone"] == "" {
		form.Errors["phone"] = "phone cant empty"
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

func (app App) validateEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")

	if len(email) < 1 || email == "" {
		fmt.Fprint(w, "email cant empty")
		return
	}

	contact := app.contacts.FindByEmail(email)
	if contact != nil {
		hxTrigger := r.Header.Get("hx-trigger")
		if hxTrigger == "edit-email" {
			hxCurrentURL, ok := r.Header[http.CanonicalHeaderKey("hx-current-url")]
			if !ok {
				fmt.Fprint(w, "hx-current-url cant empty")
				return
			}

			idStr := strings.Split(hxCurrentURL[0], "/")[4]
			id, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Fprint(w, "Error convert string to int")
				return
			}
			targetContact := app.contacts.FindByID(id)
			if contact.Email != targetContact.Email {
				fmt.Fprint(w, "duplicate email")
				return
			}
		}

		if hxTrigger == "create-email" {
			fmt.Fprint(w, "duplicate email")
			return
		}
	}

	fmt.Fprint(w, "")
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
	form.Values["ID"] = idStr
	form.Values["name"] = contact.Name
	form.Values["email"] = contact.Email
	form.Values["phone"] = contact.Phone

	components.Layout(components.PageEditContact(*form)).Render(r.Context(), w)
}
func (app App) editHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		return
	}

	form := components.NewForm()
	form.Values["id"] = idStr
	form.Values["name"] = r.PostFormValue("name")
	form.Values["email"] = r.PostFormValue("email")
	form.Values["phone"] = r.PostFormValue("phone")

	contact := components.NewContact(
		form.Values["name"],
		form.Values["email"],
		form.Values["phone"],
	)
	contact.ID = id

	app.contacts.Edit(*contact)

	http.Redirect(w, r, fmt.Sprintf("/contacts/%d", id), http.StatusSeeOther)

}
func (app App) deleteHandler(w http.ResponseWriter, r *http.Request) {

	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		return
	}

	app.contacts.Delete(id)

	hxTrigger := r.Header.Get("hx-trigger")
	if hxTrigger == "button-delete-redirect" {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		return
	}

	fmt.Fprint(w, "")

}

func (app App) TotalHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	fmt.Fprintf(w, "Total %d contact", app.contacts.Count)
}

func (app App) bulkDelete(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	selected := r.Form["bulk-delete"]

	form := components.NewForm()

	for _, idStr := range selected {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			form.Values[idStr] = fmt.Sprintf("error convert id string to int, id = %d", id)
		} else {
			app.contacts.Delete(id)
		}
	}

	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}

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
	mux.HandleFunc("POST /contacts/validate-email", app.validateEmailHandler)
	mux.HandleFunc("POST /contacts", app.createHandler)
	mux.HandleFunc("GET /contacts/{id}/edit", app.pageEditHandler)
	mux.HandleFunc("PUT /contacts/{id}", app.editHandler)
	mux.HandleFunc("DELETE /contacts/{id}", app.deleteHandler)

	mux.HandleFunc("GET /contacts/total", app.TotalHandler)
	mux.HandleFunc("POST /contacts/bulk-delete", app.bulkDelete)

	var handler http.Handler = mux

	server := new(http.Server)
	server.Addr = "127.0.0.1:8000"
	server.Handler = handler

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
