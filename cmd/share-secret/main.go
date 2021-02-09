package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var mData map[string]StoreData

func main() {
	mData = make(map[string]StoreData)

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
	r.Handle("/secret/{key}", Apply(GetSecret(), CORSWithDefaults())).Methods(http.MethodGet)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func PostSecret() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var input StoreData

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		keyData := input.GetKey()
		mData[keyData] = input

		data, err := json.Marshal(&struct {
			URI string `json:"uri"`
		}{
			URI: "https://localhost:8080/secret/" + keyData,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(data)
	}
}

func GetSecret() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		key := mux.Vars(r)["key"]
		if key == "" {
			http.Error(w, "The path is not complete", http.StatusInternalServerError)
			return
		}

		data := mData[key]
		if data.Secret == "" {
			http.Error(w, "No data", http.StatusInternalServerError)
			return
		}

		wData, err := json.Marshal(&struct {
			Data string `json:"data"`
		}{data.Secret})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(wData)

	}
}
