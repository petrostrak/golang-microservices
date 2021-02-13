package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Hello world!")
	w.Write([]byte("hello 世界!"))
}

func main() {
	http.HandleFunc("/", hello)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Sprintf("error serving: %v", err)
	}
}
