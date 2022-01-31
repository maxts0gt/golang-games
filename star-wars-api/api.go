package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const BaseURL = "https://swapi.dev/api/"

// creating Planet struct
type Planet struct {
	Name       string `json:"name"`
	Population string `json:"population"`
	Terrain    string `json:"terrain"`
}

// creating Person struct
type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld    Planet
}

// Creating a people struct which is collection of Person type
type AllPeople struct {
	// it has "People" attributes will be assigned to collection of person types
	People []Person `json:"results"`
}

// creating method for Person
func (p *Person) getHomeworld() {
	// this gets the homeworld url
	res, err := http.Get(p.HomeworldURL)
	if err != nil {
		log.Print("Error fetching homeworld", err)
	}

	// create var to access data inside if
	var bytes []byte
	if bytes, err = ioutil.ReadAll(res.Body); err != nil {
		log.Print("Error reading response body", err)
	}
	// put the bytes to Homeworld
	json.Unmarshal(bytes, &p.Homeworld)
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	// getting data here to res
	res, err := http.Get(BaseURL + "people")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request star wars people")
	}

	fmt.Println(res)
	// res is read here and assinged to bytes
	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to parse request body")
	}

	var people AllPeople
	fmt.Println(string(bytes))

	// create variable people with struct Allpeople
	// this basically parses our bytes data to a memory representation called "people"
	if err := json.Unmarshal(bytes, &people); err != nil {
		fmt.Println("Error parsing json", err)
	}

	fmt.Println(people)

	for _, pers := range people.People {
		// calling getHomeworld method from person struct
		pers.getHomeworld()
		fmt.Println(pers)
	}

}

func main() {
	http.HandleFunc("/people", getPeople)
	fmt.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
