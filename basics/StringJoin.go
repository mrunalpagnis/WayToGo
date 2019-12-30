package main

import (
	"fmt"
	"os"
	"strings"
)

/*
Usage: go run StringJoin.go <arg1> <arg2> <arg3> ...
* ... -> accepts give multiple arguments
*/
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
