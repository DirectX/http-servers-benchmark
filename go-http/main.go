package main

import (
	"fmt"
	"net/http"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	port := 4000

	http.HandleFunc("/", getRoot)

	fmt.Printf("Listening on %d\nCtrl-C to shutdown server\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
