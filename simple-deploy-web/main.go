package main

import (
	"io"
	"net/http"
)


func DefaultPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"<h1> this is web-app</h1>")
}

/*
GOOS=linux GOARCH=amd64 go build -o deploy-web main.go
 */
func main() {
	http.HandleFunc("/", DefaultPage)
	http.ListenAndServe(":9000", nil)
}
