package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup4: %v\n", err)
				continue
			}

			defer f.Close()

			countLines(f, counts)
		}
	}
	printDuplicates(counts)
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	fName := f.Name();

	for input.Scan() {
		counts[fName + ":" + input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

func printDuplicates(counts map[string]int) {

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d occurences of\t%s\n", n, line)
		}
	}
}
