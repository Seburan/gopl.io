// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
)

// echo1 implements same behavior as unix echo command
func echo1(args []string) {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

// echo2 implements same behavior as unix echo command
func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// echo3 implements same behavior as unix echo command
func echo3(args []string) {
	fmt.Println(strings.Join(args, " "))
}

// printArgsOnePerLine prints the index and value of each of args slice,
// one per line
func printArgsOnePerLine(args []string) {
	for idx, arg := range args {
		fmt.Println(idx, arg)
	}
}

func main() {
	echo1(os.Args[1:])
	echo2(os.Args[1:])
	echo3(os.Args[1:])
	printArgsOnePerLine(os.Args[1:])
}
