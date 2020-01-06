package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)
/* Usage: go run ex1_7.go <arg1> <arg2> ... <argN>
 * ... -> variable arguments are accepted
 * arguments should be a list of URLs space separated.
 */
func main() {
	for _,url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
	}
}
