package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 path not found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not supported", http.StatusNotFound)
	}

	fmt.Fprintf(w, "Hey there, this is Aditya's simple go server!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Unable to parse form, try again later")
	}

	name := r.Form["name"]
	nickname := r.Form["nickname"]

	fmt.Fprintf(w, "Hey %v, %v could be a really great nickname for you!", name, nickname)

}

func main() {
	fmt.Println("System alive...")

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Println("listening on port 8000")
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatal("I might be away for now, but I will be back alive better than ever...")
	}
}
