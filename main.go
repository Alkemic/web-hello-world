package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/env", dumpEnvVariables)
	if err := http.ListenAndServe(":8080", nil); err != http.ErrServerClosed {
		log.Fatalln("server failed: %v", err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world!"))
}

func dumpEnvVariables(w http.ResponseWriter, r *http.Request) {
	for _, pair := range os.Environ() {
		w.Write([]byte(pair))
		w.Write([]byte("\n"))
	}
}
