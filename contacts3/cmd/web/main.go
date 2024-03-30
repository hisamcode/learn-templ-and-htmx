package main

import (
	"contacts3/components"
	"net/http"
	"strconv"
)

type App struct {
	contacts components.Contacts
}

func (app App) handleContacts(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
	if err != nil {
		limit = 10
	}

	pagination := components.NewPagination(page, limit, &app.contacts.Total)
	form := components.NewForm()
	form.Values["q"] = r.URL.Query().Get("q")

	hxTrigger := r.Header.Get("hx-trigger")

	if hxTrigger == "search" {

		if len(form.Values["q"]) > 0 {
			// time.Sleep(1 * time.Second)
			cs, err := app.contacts.Search(form.Values["q"])
			if err != nil {
				form.Errors["q"] = err.Error()
				components.MainSearchErr(*form).Render(r.Context(), w)
				return
			}

			pagination := components.NewPagination(page, limit, &cs.Total)
			components.TableContacts(*cs, *pagination, true, false).Render(r.Context(), w)
			return

		} else {
			components.TableContacts(*app.contacts.Paging(pagination), *pagination, true, true).Render(r.Context(), w)
			return
		}
	}

	components.Layout(components.TableContacts(*app.contacts.Paging(pagination), *pagination, false, true), *pagination, *form).Render(r.Context(), w)
}

func main() {
	contacts := components.NewContacts()

	app := App{
		contacts: *contacts,
	}

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("../vendor/")))
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/contacts", http.StatusSeeOther)
	})

	mux.HandleFunc("GET /contacts", app.handleContacts)

	http.ListenAndServe("127.0.0.1:8000", mux)

}
