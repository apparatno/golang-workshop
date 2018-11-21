package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 3 {
		flag.Usage()
		os.Exit(1)
	}

	cmd := args[0]
	msg := args[1]
	key, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("illegal key %s (%v)\n", args[2], err)
		os.Exit(1)
	}

	if cmd == "encrypt" {
		res := encrypt(msg, key)
		fmt.Printf("%s is encrypted as %s\n", msg, res)
		os.Exit(0)
	}

	if cmd == "decrypt" {
		res := decrypt(msg, key)
		fmt.Printf("%s is decrypted to %s\n", msg, res)
		os.Exit(0)
	}

	fmt.Printf("unknown command '%s'\n", cmd)
	os.Exit(1)
}

var a = 97
var z = 122

func encrypt(message string, key int) string {
	res := make([]rune, len(message))
	rx, _ := regexp.Compile("[a-z]")

	for i, c := range message {
		if !rx.MatchString(string(c)) {
			res[i] = c
			continue
		}

		b := int(c)
		r := b + key
		if r > z {
			r = a + r - z - 1
		}
		res[i] = rune(r)
	}
	return string(res)
}

func decrypt(message string, key int) string {
	res := make([]rune, len(message))
	rx, _ := regexp.Compile("[a-z]")

	for i, c := range message {
		if !rx.MatchString(string(c)) {
			res[i] = c
			continue
		}

		b := int(c)
		r := b - key
		if r < a {
			r = z - (a - r) + 1
		}
		res[i] = rune(r)
	}
	return string(res)
}
