package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//url -> uniform resource locator!
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello world")
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Println(fmt.Sprintf("num of bytes written: %d", n))
	})

	http.ListenAndServe(":8000", nil)
}
