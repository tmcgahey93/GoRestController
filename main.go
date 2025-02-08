package main

import (
	"fmt"
	"log"
	"net/http"
)

func GoRestControllerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Troy's Go RestController Running")

}

func GoRestControllerHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Troy's Go RestController Second Entry Point Runnin")
}

// Commenting my code for comments sake
func main() {
	http.HandleFunc("/", GoRestControllerHandler)

	http.HandleFunc("/anotherentrypoint", GoRestControllerHandler2)

	fmt.Println("Server Running on portg 8080....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
