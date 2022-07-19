package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}

	http.HandleFunc("/", h1)

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}
	log.Println("START server port:" + httpPort)

	log.Fatal(http.ListenAndServe(":"+httpPort, nil))
}
