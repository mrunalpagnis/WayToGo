package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

/* Usage: go run ex1_10.go <arg1> <arg2> ... <argN>
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
	var x io.Writer
	f, err := os.Create("output" + strconv.Itoa(rand.Int()) + ".txt")
	if err != nil {
		fmt.Println(err)
		x = ioutil.Discard
	} else {
		x = f
	}

	bytes, err := io.Copy(x, resp.Body)
	_ = resp.Body.Close() //close resource
	_ = f.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s : %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, bytes, url)
}
