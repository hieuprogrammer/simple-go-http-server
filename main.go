package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/sign-up", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting Go HTTP server at port: 8080.\n")
	if error := http.ListenAndServe(":8080", nil); error != nil {
		log.Fatal(error)
	}
}

func helloHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(response, "Not found.", http.StatusNotFound)
		return
	}

	if request.Method != "GET" {
		http.Error(response, "HTTP GET is not supported.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(response, "Hello! :D")
}

func formHandler(response http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/sign-up" {
		http.Error(response, "Not found.", http.StatusNotFound)
		return
	}

	if request.Method != "POST" {
		http.Error(response, "Not supported.", http.StatusNotFound)
		return
	}

	if error := request.ParseForm(); error != nil {
		fmt.Fprintf(response, "ParseForm() error: %v", error)
		return
	}

	fmt.Fprintf(response, "POST request success.")
	name := request.FormValue("name")
	address := request.FormValue("address")
	fmt.Fprintf(response, "Name: %s.\n", name)
	fmt.Fprintf(response, "Address: %s.\n", address)
}
