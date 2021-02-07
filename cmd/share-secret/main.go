package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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
	r.HandleFunc("/secret/{key}", GetSecret).Methods(http.MethodGet)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func PostSecret() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		input := struct {
			Secret string
			Unit   string
			Time   int64
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

func GetSecret(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Not Implemented")

	// var secretData Secret
	// json.NewDecoder(r.Body).Decode(&secretData)
	// fmt.Println(r.Body)
	// fmt.Println(secretData)
}
