package main

import (
	"io"
	"net/http"
)


func DefaultPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"<h1>hello,golang!</h1>")
}

// GOOS=linux GOARCH=amd64 go build -o web-server main.go

func main() {
	http.HandleFunc("/", DefaultPage)
	http.ListenAndServe(":9000", nil)
}
