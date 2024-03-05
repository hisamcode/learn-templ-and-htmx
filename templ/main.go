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

type Item struct {
	Name string
}

type TimeValue struct {
	Time  string  `json:"time"`
	Value float64 `json:"value"`
}

func main() {
	component := Hello("Hisam")

	mux := http.NewServeMux()

	mux.Handle("GET /quickstart", templ.Handler(component))
	mux.Handle("GET /basic", templ.Handler(headerTemplate("Sapi")))
	mux.HandleFunc("GET /elements", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println()
		button("hisam", "login").Render(r.Context(), os.Stdout)
		fmt.Println()
		tagMustBeClosed().Render(r.Context(), os.Stdout)
		fmt.Println()
		w.WriteHeader(http.StatusNoContent)
	})

	mux.HandleFunc("GET /attributes", func(w http.ResponseWriter, r *http.Request) {
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

	mux.HandleFunc("GET /expressions", func(w http.ResponseWriter, r *http.Request) {
		literals().Render(r.Context(), os.Stdout)
		fmt.Println()
		variables("Hello", Person{"Hisam", ""}).Render(r.Context(), os.Stdout)
		fmt.Println()
		functions().Render(r.Context(), os.Stdout)
		fmt.Println()
		escaping().Render(r.Context(), os.Stdout)

	})

	mux.HandleFunc("GET /statements", func(w http.ResponseWriter, r *http.Request) {
		showHelloIfTrue(true).Render(r.Context(), os.Stdout)
		fmt.Println()
		display(250000.00, 3).Render(r.Context(), os.Stdout)
		fmt.Println()
	})

	mux.HandleFunc("GET /ifelse", func(w http.ResponseWriter, r *http.Request) {
		ifelse(true).Render(r.Context(), w)
	})

	mux.HandleFunc("GET /switch", func(w http.ResponseWriter, r *http.Request) {
		userTypeDisplay("Other").Render(r.Context(), os.Stdout)
	})

	mux.HandleFunc("GET /forloops", func(w http.ResponseWriter, r *http.Request) {
		items := []Item{}
		items = append(items, Item{"Hisam"})
		items = append(items, Item{"Maulana"})
		nameList(items).Render(r.Context(), os.Stdout)
	})

	mux.HandleFunc("GET /template-composition", func(w http.ResponseWriter, r *http.Request) {
		showAll().Render(r.Context(), os.Stdout)
		fmt.Println()
		c := paragraph("Dynamic contenst")
		layout(c).Render(r.Context(), os.Stdout)
		fmt.Println()
		root().Render(r.Context(), w)
	})

	mux.HandleFunc("GET /css", func(w http.ResponseWriter, r *http.Request) {

		cs := []templ.Component{
			button("test", "test"),
			button2("click me"),
			button3("login", "login-red"),
			button4("register", "blue"),
			button5("button 5 clickme", true),
			button6("Click me button 6", true),
			divLoading(),
		}

		cssLayout(cs).Render(r.Context(), w)

	})

	mux.HandleFunc("GET /javascript", func(w http.ResponseWriter, r *http.Request) {

		data := []TimeValue{
			{Time: "2019-04-11", Value: 80.01},
			{Time: "2019-04-12", Value: 96.63},
			{Time: "2019-04-13", Value: 76.64},
			{Time: "2019-04-14", Value: 81.89},
			{Time: "2019-04-15", Value: 74.43},
			{Time: "2019-04-16", Value: 80.01},
			{Time: "2019-04-17", Value: 96.63},
			{Time: "2019-04-18", Value: 76.64},
			{Time: "2019-04-19", Value: 81.89},
			{Time: "2019-04-20", Value: 74.43},
		}

		jsPage(data, "hello templ").Render(r.Context(), w)
	})

	http.ListenAndServe("127.0.0.1:8000", templ.NewCSSMiddleware(mux, primaryClassName(), className()))
}
