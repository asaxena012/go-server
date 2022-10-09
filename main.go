package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 path not found 1"))
		return
	}

	w.Write([]byte("Hey there from Aditya's go lang server!"))
}

func main() {
	fmt.Println("System alive...")

	fileServer := http.FileServer(http.Dir("/static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.Handle("/form", fileServer)

	fmt.Println("listening on port 8000")
	http.ListenAndServe("localhost:8000", nil)
}
