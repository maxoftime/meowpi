package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/api/v1/all", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/data/catdata.json")
	})

	
	http.HandleFunc("/api/v1/img", func(w http.ResponseWriter, r *http.Request) {
		imgId := r.URL.Query().Get("id")
		if imgId != "" {
			http.ServeFile(w, r, "api/data/images/" + imgId + ".jpg")
		} else {
			http.Error(w, "img not found", http.StatusNotFound)
		}
	})

	fmt.Printf("Purring @ 3333\n")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatal(err)
	}
}