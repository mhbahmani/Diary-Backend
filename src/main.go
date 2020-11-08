package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePageHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Welcome! This is home page.</h1>")
}

func handleRequests(){
	http.HandleFunc("/", homePageHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {	
	handleRequests()
}
