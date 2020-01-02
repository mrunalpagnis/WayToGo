package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:], " "))
	time1 := time.Since(start).Seconds()
	fmt.Printf("time taken by string join : %.2fs\n", time1)

	start = time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	time2 := time.Since(start).Seconds()
	fmt.Printf("time taken by loop : %.2fs\n", time2)
}
