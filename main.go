package main

import (
	"net/http"
	"io"

	log "github.com/Sirupsen/logrus"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
