package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
Usage go run WebServer.go
the program does not take any command line args.
To see the output launch in browser "localhost:8000" with any URL path.
eg: http://localhost:8000/this/is/a/test
Output will be -> URL.path = "/this/is/a/test"
*/
func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "URL.path = %q\n", r.URL.Path)
}
