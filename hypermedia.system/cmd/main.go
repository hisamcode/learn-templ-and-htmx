package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/hisamcode/try-htmx/hs/components"
)

var Contacts []components.Contact

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("../vendor/")))

	mux.HandleFunc("GET /contacts", func(w http.ResponseWriter, r *http.Request) {
		components.Layout("Contacts", []templ.Component{components.Contacts(Contacts)}).Render(r.Context(), w)
	})

	mux.HandleFunc("GET /contacts/new", func(w http.ResponseWriter, r *http.Request) {
		components.Layout("New Contact", []templ.Component{components.ContactNew()}).Render(r.Context(), w)
	})

	mux.HandleFunc("DELETE /contacts/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			fmt.Println(err)
			return
		}

		key, err := findKeyContactById(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		newContacts := []components.Contact{}
		newContacts = append(newContacts, Contacts[:*key]...)
		newContacts = append(newContacts, Contacts[*key+1:]...)
		Contacts = newContacts

		// components.Layout("Contacts", []templ.Component{components.Contacts(Contacts)}).Render(r.Context(), w)
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)

	})

	mux.HandleFunc("POST /contacts", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			fmt.Println(err)
			return
		}

		if !r.PostForm.Has("id") {
			fmt.Println("required id")
			return
		}

		id := r.PostForm.Get("id")
		if len(id) < 1 {
			fmt.Println("id must > 1")
			return
		}

		newId, err := strconv.Atoi(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		if !r.PostForm.Has("name") {
			fmt.Println("required name")
		}

		Name := r.PostForm.Get("name")
		fmt.Println("create contacts", newId, Name)

		Contacts = append(Contacts, components.Contact{ID: newId, Name: Name})
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
		// components.Contacts([]string{"post 1", "post 2", r.FormValue("q")}).Render(r.Context(), w)
	})

	mux.HandleFunc("GET /contacts/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			fmt.Println(err)
			return
		}

		contact, err := findContactById(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		components.PageContact(*contact).Render(r.Context(), w)

	})

	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		cs := []templ.Component{
			components.ButtonGetContact(),
			components.ButtonGetContact2(),
			components.ButtonGetContact3(),
		}
		components.Layout("htmx", cs).Render(r.Context(), w)
	})

	mux.HandleFunc("GET /boosted-link", func(w http.ResponseWriter, r *http.Request) {
		cs := []templ.Component{
			components.BoostedLink(),
		}

		components.Layout("boosted link", cs).Render(r.Context(), w)
	})

	mux.HandleFunc("GET /settings", func(w http.ResponseWriter, r *http.Request) {
		components.PageSettings().Render(r.Context(), w)
	})

	mux.HandleFunc("GET /boosted-forms", func(w http.ResponseWriter, r *http.Request) {
		components.BoostedForm().Render(r.Context(), w)
	})

	mux.HandleFunc("POST /message", func(w http.ResponseWriter, r *http.Request) {
		msg := r.FormValue("message")
		w.Write([]byte(msg))
	})

	fmt.Println("listen on 127.0.0.1:8000")
	if err := http.ListenAndServe("127.0.0.1:8000", mux); err != nil {
		panic(err)
	}
}

func findKeyContactById(id int) (*int, error) {
	for k, v := range Contacts {
		if v.ID == id {
			return &k, nil
		}
	}
	return nil, errors.New("not found")

}

func findContactById(id int) (*components.Contact, error) {
	for _, v := range Contacts {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, errors.New("not found")
}
