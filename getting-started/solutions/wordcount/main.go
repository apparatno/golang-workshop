package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	text, err := readfile()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	res := count(text)
	top := topfive(res)
	for i, w := range top {
		fmt.Printf("%d. %s (%d)\n", i+1, w.value, w.occurrences)
	}
}

func count(text string) map[string]int {
	r := regexp.MustCompile("[^a-z]")
	words := strings.Split(text, " ")
	res := make(map[string]int)
	for _, w := range words {
		w = strings.ToLower(w)
		w = r.ReplaceAllString(w, "")
		w = strings.TrimSpace(w)
		if w == "" {
			continue
		}
		res[w]++
	}
	return res
}

func readfile() (string, error) {
	bytes, err := ioutil.ReadFile("text.txt")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

type word struct {
	value       string
	occurrences int
}

type byOccurrences []word

func (w byOccurrences) Len() int           { return len(w) }
func (w byOccurrences) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w byOccurrences) Less(i, j int) bool { return w[i].occurrences > w[j].occurrences }

func topfive(words map[string]int) []word {
	res := make(byOccurrences, len(words))
	c := 0
	for w, n := range words {
		res[c] = word{w, n}
		c++
	}
	sort.Sort(res)
	return res[:5]
}
