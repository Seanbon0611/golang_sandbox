package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type Dog struct {
	ID    string `json:"id"`
	Breed string `json:"breed"`
	Name  string `json:"name"`
	Age   string `json:"age"`
	Notes string `json:"notes"`
}

func getDogs(w http.ResponseWriter, r *http.Request) {
	span := tracer.StartSpan("web.request", tracer.ResourceName("/dogs"))

	defer span.Finish()
	// Set tag
	span.SetTag("http.url", r.URL.Path)
	span.SetTag("testkey", "testvalue")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dogs)
}

func addDog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var dog Dog
	_ = json.NewDecoder(r.Body).Decode(&dog)
	dog.ID = strconv.Itoa(rand.Intn(10000000))
	dogs = append(dogs, dog)
	json.NewEncoder(w).Encode(dog)
}

var dogs []Dog

func main() {
	tracer.Start(
		tracer.WithService("test1"),
		tracer.WithEnv("prod"),
	)
	defer tracer.Stop()
	mux := muxtrace.NewRouter()

	dogs = append(dogs, Dog{ID: "1", Breed: "Shiba Inu", Name: "Hachiko", Age: "10", Notes: "A bit of a brat"})
	dogs = append(dogs, Dog{ID: "2", Breed: "Golden Retriever", Name: "Lassy", Age: "8", Notes: "Likes to run"})
	mux.HandleFunc("/dogs", getDogs).Methods("GET")
	mux.HandleFunc("/dogs/new", addDog).Methods("POST")
	// mux.HandleFunc("/dogs/{id}", getSingleDog).Methods(("GET"))
	port := ":8080"
	fmt.Println("\nListening on port " + port)
	log.Fatal(http.ListenAndServe(port, mux))
}
