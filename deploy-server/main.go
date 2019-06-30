package main

import (
	"io"
	"log"
	"net/http"
	"os/exec"
)


func Run(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"<h1>Deploy running </h1>")

	cmd:=exec.Command("sh","./deploy.sh")
	err := cmd.Start()
	if err!=nil{
		log.Fatalln("cmd  start error")
	}


	io.WriteString(w,"<h1>Deploy ending </h1>")
	cmd.Wait()

}

func Ping(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"<h1>Deploy pong</h1>")
}



// GOOS=linux GOARCH=amd64 go build -o deploy-server main.go
func main() {
	http.HandleFunc("/", Ping)
	http.HandleFunc("/run", Run)
	http.ListenAndServe(":5000", nil)
}
