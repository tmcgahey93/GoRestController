package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type person struct {
	name string
	age  int
}

func (p person) Greet() string {
	return "Hello, my name is " + p.name + " and I am " + fmt.Sprint(p.age) + " years old."
}

func getEnvironmentHandler(w http.ResponseWriter, r *http.Request) {
	returnMessage := "Troy's Go RestController Running in " + os.Getenv("ENV") + " mode"
	fmt.Fprintln(w, returnMessage)
}

func getTroyHandler(w http.ResponseWriter, r *http.Request) {
	p := person{name: "Troy", age: 50}
	returnMessage := p.Greet()
	fmt.Fprintln(w, returnMessage)
}

func main() {

	env := os.Getenv("ENV")
	if env == "development" || env == "" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	http.HandleFunc("/getTroy", getTroyHandler)

	http.HandleFunc("/getEnvironment", getEnvironmentHandler)

	fmt.Println("Server Running on port 8080....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
