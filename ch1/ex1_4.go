// Modified "dup2" program to print the names
// of each file where duplicated line occurs

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]int) // https://go.dev/blog/maps
	files := os.Args[1:]
	if (len(files) == 0) {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
		for file_name, line_map :=range counts {
			printed_file_name := false
			for line, n := range line_map {
				if n > 1 {
					if !printed_file_name {
						printed_file_name = true
						fmt.Printf("In file:\t%s\n", file_name)
					}
					fmt.Printf("%d\t%s\n", n, line)
				}
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// https://go.dev/blog/maps
		line_map, exists := counts[f.Name()]
		if !exists {
			line_map = make(map[string]int)
			counts[f.Name()] = line_map
		}

		counts[f.Name()][input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}