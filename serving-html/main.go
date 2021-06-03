package main

import (
	"fmt"
	"net/http"
)

//The port you want to serve to
const PORT = ":8000"

func main() {
	//url -> uniform resource locator!
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Sprintf("Listening at PORT%s", PORT))
	http.ListenAndServe(PORT, nil)
}
