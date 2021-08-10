package main

import (
	"fmt"
	"net/http"

	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

type User struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fmt.Println("Hello, world")
	mux := muxtrace.NewRouter()
	mux.HandleFunc("/", HomeHandler)

}
