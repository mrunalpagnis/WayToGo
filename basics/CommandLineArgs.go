package main

import (
	"fmt"
	"os"
)

/*
Usage: go run CommandLineArgs.go <arg1> <arg2> <arg3> ...
* ... -> accepts give multiple arguments
*/
func main() {
	var s, sep string
	for i:= 1; i< len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
