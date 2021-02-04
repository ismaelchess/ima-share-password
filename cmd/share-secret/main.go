package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Secret struct {
	SecretValue string `json:"secretvalue"`
	SecretTime  int    `json:"secrettime"`
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := fmt.Fprintln(w, "hello world"); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	r.Handle("/secret", Apply(PostSecret(), CORSWithDefaults())).Methods(http.MethodPost)
	r.HandleFunc("/secret/{key}", sharesecretlink).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
	return nil
}

func PostSecret() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := struct {
			Secret string
		}{}
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data, err := json.Marshal(&struct {
			URI string `json:"uri"`
		}{
			URI: "https://example.com/secret/123456789",
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(data)
	}
}

func sharesecretlink(w http.ResponseWriter, r *http.Request) {

	fmt.Println("xxxxxxx")

	// var secretData Secret
	// json.NewDecoder(r.Body).Decode(&secretData)
	// fmt.Println(r.Body)
	// fmt.Println(secretData)
}
