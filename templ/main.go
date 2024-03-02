package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/a-h/templ"
)

type Person struct {
	Name string
	URL  string
}

const exclamation = "!"

func main() {
	component := Hello("Hisam")

	http.Handle("GET /quickstart", templ.Handler(component))
	http.Handle("GET /basic", templ.Handler(headerTemplate("Sapi")))
	http.HandleFunc("GET /elements", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		button("hisam", "login").Render(r.Context(), os.Stdout)
		fmt.Println()
		tagMustBeClosed().Render(r.Context(), os.Stdout)
		fmt.Println()
		w.WriteHeader(http.StatusNoContent)
	})
	http.HandleFunc("GET /attributes", func(w http.ResponseWriter, r *http.Request) {
		constanAttribute().Render(r.Context(), os.Stdout)
		fmt.Println()
		booleanAttribute(false).Render(r.Context(), os.Stdout)
		fmt.Println()
		conditionalAttributes().Render(r.Context(), os.Stdout)
		fmt.Println()
		usageSpreadAttribute().Render(r.Context(), os.Stdout)
		fmt.Println()
		URLAttribute(Person{"Hisam", "github.com/hisamcode"}).Render(r.Context(), os.Stdout)
		fmt.Println()
		jsAttrButton("halo").Render(r.Context(), os.Stdout)
		jsAttrButton("hai").Render(r.Context(), w)
		fmt.Println()
	})

	http.HandleFunc("GET /expressions", func(w http.ResponseWriter, r *http.Request) {
		literals().Render(r.Context(), os.Stdout)
		fmt.Println()
		variables("Hello", Person{"Hisam", ""}).Render(r.Context(), os.Stdout)
		fmt.Println()
		functions().Render(r.Context(), os.Stdout)
		fmt.Println()
		escaping().Render(r.Context(), os.Stdout)

	})

	http.HandleFunc("GET /statements", func(w http.ResponseWriter, r *http.Request) {
		showHelloIfTrue(true).Render(r.Context(), os.Stdout)
		fmt.Println()
		display(250000.00, 3).Render(r.Context(), os.Stdout)
		fmt.Println()
	})

	http.HandleFunc("GET /ifelse", func(w http.ResponseWriter, r *http.Request) {
		ifelse(true).Render(r.Context(), w)
	})

	http.ListenAndServe("127.0.0.1:8000", nil)
}
