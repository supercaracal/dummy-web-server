package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

const (
	contentTypeKey  = "Content-Type"
	contentTypeJSON = "application/json;charset=UTF-8"
	bindAddr        = "0.0.0.0"
	defaultPort     = 3000
	rootDir         = "/"
)

var (
	responseBody = []byte("{}\n")
)

func main() {
	port := getPortNumber(defaultPort)

	if isSubCommand("health") {
		if err := checkHealth(port, rootDir); err != nil {
			log.Fatalf("NG: %s", err)
		}

		log.Printf("OK")
		os.Exit(0)
	}

	setHandler(rootDir, responseBody)

	log.Printf("listen: %s:%d", bindAddr, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", bindAddr, port), nil))
}

func getPortNumber(defaultNumber int) int {
	port := defaultNumber
	if v := os.Getenv("PORT"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 1023 {
			port = n
		}
	}

	flag.IntVar(&port, "port", port, "listen port number")
	flag.Parse()

	return port
}

func isSubCommand(cmd string) bool {
	for _, arg := range os.Args {
		if arg == cmd {
			return true
		}
	}

	return false
}

func checkHealth(port int, dir string) error {
	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%d%s", port, dir))
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if _, err := io.ReadAll(resp.Body); err != nil {
		return err
	}

	return nil
}

func setHandler(dir string, body []byte) {
	http.HandleFunc(dir, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentTypeKey, contentTypeJSON)
		w.WriteHeader(http.StatusOK)
		w.Write(body)
		log.Print(r.URL.Path)
	})
}
