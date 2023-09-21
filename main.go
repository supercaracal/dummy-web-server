package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	contentTypeKey  = "Content-Type"
	contentTypeJSON = "application/json;charset=UTF-8"
	defaultPort     = 3000
	rootDir         = "/"
)

var (
	responseBody = []byte("{}\n")
)

func main() {
	port := defaultPort
	if v := os.Getenv("PORT"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 1023 {
			port = n
		}
	}

	flag.IntVar(&port, "port", port, "listen port number")
	flag.Parse()

	http.HandleFunc(rootDir, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentTypeKey, contentTypeJSON)
		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
		log.Print(r.URL.Path)
	})

	log.Printf("listen: 0.0.0.0:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))
}
