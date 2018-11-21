package main

import "testing"

func TestReversing(t *testing.T) {
	input := "test string here"
	expected := "ereh gnirts tset"

	res := reverse(input)
	if res != expected {
		t.Errorf("got %s want %s", res, expected)
	}
}
