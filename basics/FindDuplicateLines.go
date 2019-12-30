package main

import (
	"bufio"
	"fmt"
	"os"
)
/*
Usage: go run FindDuplicateLines.go
* once the program is running enter lines, each on new line.
* end input by typing exit (press enter)
*/
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		if input.Text() == "exit" {
			break
		}
		counts[input.Text()]++
	}
	for line, count := range counts {
		if count > 1 {
			fmt.Printf("%s :\t%d\n",line,count)
		}
	}
}
