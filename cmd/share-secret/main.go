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

	r.HandleFunc("/secret", sharesecretpost).Methods(http.MethodPost)
	r.HandleFunc("/secret/{key}", sharesecretlink).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", r))
	return nil
}

func sharesecretpost(w http.ResponseWriter, r *http.Request) {

	var secretData Secret
	json.NewDecoder(r.Body).Decode(&secretData)

	fmt.Println(r.Body)
	fmt.Println(secretData)
	secretData.SecretTime = 0

	json.NewEncoder(w).Encode(secretData)
}

func sharesecretlink(w http.ResponseWriter, r *http.Request) {

	fmt.Println("xxxxxxx")

	// var secretData Secret
	// json.NewDecoder(r.Body).Decode(&secretData)
	// fmt.Println(r.Body)
	// fmt.Println(secretData)
}
