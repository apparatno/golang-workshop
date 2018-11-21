package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	args, err := convertArgs(flag.Args())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	result := sum(args)
	fmt.Printf("The sum is %d\n", result)
}

func convertArgs(a []string) ([]int, error) {
	res := make([]int, len(a))
	for i, n := range a {
		v, err := strconv.ParseInt(n, 10, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to parse number '%s': %v", n, err)
		}
		res[i] = int(v)
	}
	return res, nil
}

func sum(numbers []int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}
