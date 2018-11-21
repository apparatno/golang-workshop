package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("no input")
		os.Exit(1)
	}

	s := args[0]
	result := reverse(s)
	fmt.Println(result)
}

func reverse(s string) string {
	c := len(s)
	res := make([]rune, c)

	for _, r := range s {
		c--
		res[c] = r
	}

	return string(res)
}
