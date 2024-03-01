package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("../vendor/")))

	fmt.Println("listen on 127.0.0.1:8000")
	if err := http.ListenAndServe("127.0.0.1:8000", mux); err != nil {
		panic(err)
	}
}
