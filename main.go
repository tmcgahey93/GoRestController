package main

import (
	"fmt"
	"log"
	"net/http"
)

func GoRestControllerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Troy's Go RestController Running")
}

func main() {
	http.HandleFunc("/", GoRestControllerHandler)

	fmt.Println("Server Running on portg 8080....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
