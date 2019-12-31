package main

import (
	"fmt"
	"os"
)

/*
Usage: go run ImplicitStringJoinUsingFmt.go <arg1> <arg2> <arg3> ...
* ... -> accepts give multiple arguments
*/
func main() {
	fmt.Println(os.Args[1:])
}
