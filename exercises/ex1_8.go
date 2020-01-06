package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)
/* Usage: go run ex1_8.go <arg1> <arg2> ... <argN>
 * ... -> variable arguments are accepted
 * arguments should be a list of URLs space separated.
 */
func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		fmt.Printf("%s -> %v\n", url, b)
	}
}
