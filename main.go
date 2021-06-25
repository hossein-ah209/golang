package main

import (
	"fmt"
	"net/http"
)

func helloWeb(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "helloWeb")
}
func main() {
	fmt.Println("Start a web server....")
	http.HandleFunc("/h", helloWeb)
	http.ListenAndServe(":8000", nil)
}
