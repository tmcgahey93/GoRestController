package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (p *person) HaveBirthday() {
	p.age++
}

type dog struct {
	name string
	age  int
}

func (d dog) Greet() string {
	return "Hello, my name is " + d.name + " and I am " + fmt.Sprint(d.age) + " years old."
}

func (d *dog) HaveBirthday() {
	d.age++
}

type mammal interface {
	Greet() string
	HaveBirthday()
}

type TimeResponse struct {
	Datetime string `json:"datetime"`
}

func getTime() (string, error) {
	url := "https://worldtimeapi.org/api/timezone/Etc/UTC"
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("error: received status code %d", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	var timeResponse TimeResponse
	err = json.Unmarshal(bodyBytes, &timeResponse)
	if err != nil {
		return "", fmt.Errorf("error unmarshalling response: %v", err)
	}

	return timeResponse.Datetime, nil

}

func getEnvironmentHandler(w http.ResponseWriter, r *http.Request) {
	returnMessage := "Troy's Go RestController Running in " + os.Getenv("ENV") + " mode"
	fmt.Fprintln(w, returnMessage)
}

func getTimeHandler(w http.ResponseWriter, r *http.Request) {
	returnString, err := getTime()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	returnMessage := "Current UTC Time: " + returnString

	fmt.Fprintln(w, returnMessage)
}

func getMammalHandler(w http.ResponseWriter, r *http.Request) {

	var p *person = &person{name: "Troy", age: 50}
	var d *dog = &dog{name: "Fido", age: 8}

	var value mammal

	if r.URL.Query().Get("type") == "person" {
		value = p
	} else {
		value = d
	}

	value.HaveBirthday()
	returnMessage := value.Greet()
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

	http.HandleFunc("/getMammal", getMammalHandler)

	http.HandleFunc("/getEnvironment", getEnvironmentHandler)

	http.HandleFunc("/getTime", getTimeHandler)

	fmt.Println("Server Running on port 8080....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
