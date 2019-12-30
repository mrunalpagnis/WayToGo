package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Usage: go run FindDuplicateLinesInMultipleFiles.go <arg1> <arg2> <arg3> ...
* ... -> accepts give multiple arguments
* arguments should be a file name (including absolute path)
*/
func main() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			_ = f.Close()
		}
	}
	for line, fileMap := range counts {
		if len(fileMap) > 1 {
			fmt.Printf("%s\t%v\n", line, fileMap)
		} else {
			for _, count := range fileMap {
				if count > 1 {
					fmt.Printf("%s\t%v\n", line, fileMap)
				}
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		fileMap := counts[input.Text()]
		if fileMap == nil {
			fileMap = make(map[string]int)
		}
		fileMap[f.Name()]++
		counts[input.Text()] = fileMap
	}
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
