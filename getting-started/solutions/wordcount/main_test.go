package main

import (
	"testing"
)

func TestTopFive(t *testing.T) {
	data := map[string]int{
		"foo":     3,
		"bar":     6,
		"baz":     1,
		"lorem":   99,
		"ipsum":   87,
		"doloret": 7,
		"sit":     6,
		"amet":    11,
		"and":     2,
		"maybe":   3,
		"some":    51,
		"more":    8,
		"words":   12,
	}
	res := topfive(data)
	if len(res) != 5 {
		t.Errorf("expected result to be 5 items, got %d", len(res))
	}

	helper(t, res[0], word{"lorem", 99})
	helper(t, res[1], word{"ipsum", 87})
	helper(t, res[2], word{"some", 51})
	helper(t, res[3], word{"words", 12})
	helper(t, res[4], word{"amet", 11})
}

func helper(t *testing.T, w, e word) {
	if w.value != e.value {
		t.Errorf("expected %s got %s", e.value, w.value)
	}
	if w.occurrences != e.occurrences {
		t.Errorf("expected %d got %d", e.occurrences, w.occurrences)
	}
}

func TestCount(t *testing.T) {
	text := "a text that makes no sense to a person that reads it"
	res := count(text)
	if len(res) != 10 {
		t.Errorf("expected 10 items got %d", len(res))
	}
	countHelper(t, res, "a", 2)
	countHelper(t, res, "that", 2)
	countHelper(t, res, "makes", 1)
	countHelper(t, res, "no", 1)
	countHelper(t, res, "person", 1)
}

func countHelper(t *testing.T, res map[string]int, word string, occ int) {
	if res[word] != occ {
		t.Errorf("expected '%s' to be counted %d times got %d", word, occ, res[word])
	}
}
