package main

import (
	"fmt"
	"net/http"
)

// 웹 서버 열림
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	http.ListenAndServe(":3000", nil)
}
