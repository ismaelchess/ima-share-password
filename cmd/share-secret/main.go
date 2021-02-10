package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

var mData sync.Map

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
		mData.Store(keyData, input.Secret)

		time.AfterFunc(input.expirationDate(), func() {
			fmt.Println("Borrado")
			mData.Delete(keyData)
		})

		data, err := json.Marshal(&struct {
			URI string `json:"uri"`
		}{
			URI: "http://localhost:8080/secret/" + keyData,
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

		result, ok := mData.Load(key)
		if !ok {
			http.Error(w, "No data", http.StatusInternalServerError)
			return
		}

		wData, err := json.Marshal(&struct {
			Data string `json:"data"`
		}{
			Data: result.(string),
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, _ = w.Write(wData)

	}
}
