package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// any request will be handled by this function
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello World!")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "Oop", http.StatusBadRequest)
			return
		}
		// write the response
		fmt.Fprintf(rw, "Hello %s", d)
	})
	// reqeusts to the path /goodbye will be handled by this function
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})
	// Listen for connections on all ip addresses (0.0.0.0)
	// port 9090
	http.ListenAndServe(":9090", nil)
}
