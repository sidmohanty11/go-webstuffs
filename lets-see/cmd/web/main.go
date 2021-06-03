package main

import (
	"fmt"
	"github.com/sidmohanty11/go-webstuffs/pkgs/handlers"
	"net/http"
)

//The port you want to serve to
const PORT = ":8000"

func main() {
	//url -> uniform resource locator!
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Sprintf("Listening at PORT%s", PORT))
	http.ListenAndServe(PORT, nil)
}
