package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	// router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, Vedant")
	// })

	// router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, Mrunal")
	// })

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("templates/home.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, "Prajwal")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, World")
	// })

	// http.ListenAndServe(":8080", nil)

	http.ListenAndServe(":8080", router)
}
