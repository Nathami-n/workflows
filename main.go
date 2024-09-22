package main 

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server is running on port 8000")
	http.ListenAndServe(":8000", nil)
}