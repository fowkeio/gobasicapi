package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rest API with GO!")
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		log.Println("'/status' was called")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
