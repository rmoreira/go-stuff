package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
	cmd := exec.Command("/usr/local/opt/go/src/github.com/rmoreira/go-stuff/healthcheck")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, out.String())
}

func main() {
	http.HandleFunc("/healthcheck", handler)
	http.ListenAndServe(":8080", nil)
}
