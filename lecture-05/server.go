package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		fmt.Fprint(w, "this only accepts get methods")
	}

	fmt.Fprint(w, "noooooooooo")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered")
		}
	}()

	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
