package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	// http.HandleFunc takes two parameters, the path abd
	// the function, we can pass variable that has function
	// assigned to it
	http.HandleFunc
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8000", nil))

}

// func maina() {
// 	http.HandleFunc("/hello", func(rw http.ResponseWriter, req *http.Request) {
// 		name := req.URL.Query().Get("name")
// 		rw.Write([]byte(fmt.Sprintf("Hello, %s", name)))
// 	}
// )
// 	http.ListenAndServe(":8080", nil)
// }
