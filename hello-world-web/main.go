package main

import (
	"fmt"
	"log"
	"net/http"
)

//The port you want to serve to
const PORT = ":8000"

//Home is the home page Handler
func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello world")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

//About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, err := fmt.Fprintf(w, fmt.Sprintf("The About Page and 2 + 2 is %d", sum))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

//addValues add two ints and returns the sum
func addValues(x, y int) int {
	return x + y
}

func main() {
	//url -> uniform resource locator!
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Sprintf("Listening at PORT%s", PORT))
	http.ListenAndServe(PORT, nil)
}
