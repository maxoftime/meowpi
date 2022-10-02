package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/api/v1/all", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "data/catdata.json")
	})

	
	http.HandleFunc("/api/v1/img", func(w http.ResponseWriter, r *http.Request) {
		imgId := r.URL.Query().Get("id")
		if imgId != "" {
			http.ServeFile(w, r, "data/images/" + imgId + ".jpg")
		} else {
			http.Error(w, "img not found", http.StatusNotFound)
		}
	})


	fmt.Printf("Purring @ 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}