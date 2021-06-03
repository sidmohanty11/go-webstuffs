package main

import (
	"fmt"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello world")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)
	_, err := fmt.Fprintf(w, fmt.Sprintf("The About Page and 2 + 2 is %d", sum))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func AddValues(x, y int) int {
	return x + y
}

func main() {
	//url -> uniform resource locator!
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.ListenAndServe(":8000", nil)
}
