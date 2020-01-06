package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

/* Usage: go run FetchUrlsConcurrently.go <arg1> <arg2> ... <argN>
 * ... -> variable arguments are accepted
 * arguments should be a list of URLs space separated.
 */
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //receive from channel ch
	}
	fmt.Printf("%.2fs elapsed \n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- err.Error() //send error to channel
		return
	}
	bytes, err := io.Copy(ioutil.Discard, resp.Body)
	_ = resp.Body.Close() //close resource
	if err != nil {
		ch <- fmt.Sprintf("while reading %s : %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, bytes, url)
}
