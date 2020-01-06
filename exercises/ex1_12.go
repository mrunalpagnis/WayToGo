package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

/*
Usage go run ex1_12.go
the program does not take any command line args.
To see the output launch in browser "localhost:8000" with any URL path and parameter cycles=<value>.
eg: http://localhost:8000/this/is/a/test?cycles=20
Output will be -> cycles = "20"
*/
func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	c := 5
	for k, v := range r.Form {
		if k == "cycles" {
			c, _ = strconv.Atoi(v[0])
		}
	}

	_, _ = fmt.Fprintf(w, "cycles = %d\n", c)
}
