package main

import (
	"errors"
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

//Divide is the div page handler
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.00, 0.00)

	if err != nil {
		fmt.Fprintf(w, "cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.00, 10.00, f))
}

//addValues add two ints and returns the sum
func addValues(x, y int) int {
	return x + y
}

//divideValues divides two nums
func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	res := x / y
	return res, nil
}

func main() {
	//url -> uniform resource locator!
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/div", Divide)
	fmt.Println(fmt.Sprintf("Listening at PORT%s", PORT))
	http.ListenAndServe(PORT, nil)
}
