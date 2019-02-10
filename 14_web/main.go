package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "<h1>About</h1>")
}
func main()  {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server Running on: 3000")
	http.ListenAndServe(":3000", nil)
}
