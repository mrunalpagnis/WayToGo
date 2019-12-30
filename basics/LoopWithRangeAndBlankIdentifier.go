package main

import (
	"fmt"
	"os"
)
/*
Usage: go run LoopWithRangeAndBlankIdentifier.go <arg1> <arg2> <arg3> ...
* ... -> accepts give multiple arguments
*/
func main() {
	s, sep:= "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
