package main

import (
	"context"
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
		button("hisam", "login").Render(context.Background(), os.Stdout)
		fmt.Println()
		tagMustBeClosed().Render(context.Background(), os.Stdout)
		fmt.Println()
		w.WriteHeader(http.StatusNoContent)
	})
	http.HandleFunc("GET /attributes", func(w http.ResponseWriter, r *http.Request) {
		constanAttribute().Render(context.Background(), os.Stdout)
		fmt.Println()
		booleanAttribute(false).Render(context.Background(), os.Stdout)
		fmt.Println()
		conditionalAttributes().Render(context.Background(), os.Stdout)
		fmt.Println()
		usageSpreadAttribute().Render(context.Background(), os.Stdout)
		fmt.Println()
		URLAttribute(Person{"Hisam", "github.com/hisamcode"}).Render(context.Background(), os.Stdout)
		fmt.Println()
		jsAttrButton("halo").Render(context.Background(), os.Stdout)
		jsAttrButton("hai").Render(context.Background(), w)
		fmt.Println()
	})

	http.HandleFunc("GET /expressions", func(w http.ResponseWriter, r *http.Request) {
		literals().Render(context.Background(), os.Stdout)
		fmt.Println()
		variables("Hello", Person{"Hisam", ""}).Render(context.Background(), os.Stdout)
		fmt.Println()
		functions().Render(context.Background(), os.Stdout)
		fmt.Println()
		escaping().Render(context.Background(), os.Stdout)

	})

	http.ListenAndServe("127.0.0.1:8000", nil)

	component.Render(context.Background(), os.Stdout)

}
