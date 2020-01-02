package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)
/* Usage: go run ex1_9.go <arg1> <arg2> ... <argN>
 * ... -> variable arguments are accepted
 * arguments should be a list of URLs space separated.
 */
func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\t%v\n", resp.Status, err)
			continue
		}
		fmt.Printf("%s -> %s\t%v\n", resp.Status, url, b)
	}
}
