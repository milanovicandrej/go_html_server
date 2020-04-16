package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	indexPage, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fmt.Fprintf(w, string(indexPage))
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Server Starting (port = 3000)")
	http.ListenAndServe(":3000", nil)
}
