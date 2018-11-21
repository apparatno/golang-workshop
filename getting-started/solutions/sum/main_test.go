package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	input := []int{2, 4, 6}
	result := sum(input)

	if result != 12 {
		t.Errorf("wrong result, got %d want 12", result)
	}
}

func TestSumEmptyInput(t *testing.T) {
	input := make([]int, 0, 0)
	result := sum(input)

	if result != 0 {
		t.Errorf("wrong result, got %d want 0", result)
	}
}

func TestConvertArgs(t *testing.T) {
	args := []string{"2", "9", "5"}
	result, err := convertArgs(args)
	if err != nil {
		t.Errorf("error %v", err)
	}

	expected := []int{2, 9, 5}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("wrong result, got %v want %v", result, expected)
	}
}

func TestConvertBadArgs(t *testing.T) {
	args := []string{"2", "9", "a"}
	_, err := convertArgs(args)
	if err == nil {
		t.Error("expected an error")
	}
}
